package handlers

import (
	"encoding/json"
	"github.com/Julia1505/RedditCloneBack/pkg/post"
	"github.com/Julia1505/RedditCloneBack/pkg/user"
	"github.com/Julia1505/RedditCloneBack/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
)

type PostHandler struct {
	PostStorage post.PostsRepo
}

func (h *PostHandler) List(w http.ResponseWriter, r *http.Request) {
	elems, err := h.PostStorage.GetAll()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	utils.JSON(w, elems, http.StatusOK)
}

func (h *PostHandler) CategoryList(w http.ResponseWriter, r *http.Request) {
	category, ok := mux.Vars(r)["category_name"]
	if !ok {
		http.Error(w, "no category", http.StatusInternalServerError)
		return
	}

	elems, err := h.PostStorage.GetByCategory(category)
	if err != nil {
		utils.JSON(w, elems, http.StatusOK)
		return
	}
	utils.JSON(w, elems, http.StatusOK)
}

func (h *PostHandler) UserList(w http.ResponseWriter, r *http.Request) {
	username, ok := mux.Vars(r)["user_login"]
	if !ok {
		http.Error(w, "no category", http.StatusInternalServerError)
		return
	}

	elems, err := h.PostStorage.GetByUser(username)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	utils.JSON(w, elems, http.StatusOK)
}

func (h *PostHandler) Post(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["post_id"]
	if !ok {
		http.Error(w, "uncorrect id", http.StatusInternalServerError)
		return
	}

	elem, err := h.PostStorage.GetById(id)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	utils.JSON(w, elem, http.StatusOK)
}

func (h *PostHandler) AddPost(w http.ResponseWriter, r *http.Request) {
	elem := post.NewPost()
	err := json.NewDecoder(r.Body).Decode(elem)
	if err != nil {
		http.Error(w, "decode err", http.StatusInternalServerError)
		return
	}

	curUser, err := user.FromContext(r.Context())
	if err != nil {
		http.Error(w, "no auth", http.StatusUnauthorized)
		return
	}
	elem.Author = post.Author{Id: curUser.Id, Username: curUser.Username}

	switch elem.Type {
	case "text":
		if elem.Text == "" {
			http.Error(w, "no text", http.StatusInternalServerError)
			return
		}
	case "link":
		if elem.Url == "" {
			http.Error(w, "no url", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "unknown type of post", http.StatusInternalServerError)
		return
	}

	newVote := post.NewVote(curUser.Id, 1)
	elem.Votes = append(elem.Votes, newVote)
	elem.Score = 1
	elem.UpdateVotes()

	_, err = h.PostStorage.AddPost(elem)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	utils.JSON(w, elem, http.StatusCreated)
}

type BodyCom struct {
	Comment string `json:"comment"`
}

func (h *PostHandler) AddComment(w http.ResponseWriter, r *http.Request) {
	postId := mux.Vars(r)["post_id"]
	elem := &BodyCom{}
	err := json.NewDecoder(r.Body).Decode(&elem)
	if err != nil {
		http.Error(w, "decoed err", http.StatusInternalServerError)
		return
	}
	newComment := post.NewComment()
	newComment.Body = elem.Comment

	curUser, err := user.FromContext(r.Context())
	if err != nil {
		http.Error(w, "no auth", http.StatusUnauthorized)
		return
	}
	newComment.Author = post.Author{Id: curUser.Id, Username: curUser.Username}

	curPost, err := h.PostStorage.GetById(postId)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	curPost.Comments = append(curPost.Comments, newComment)
	_, err = h.PostStorage.UpdatePost(curPost)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	utils.JSON(w, curPost, http.StatusCreated)
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	postId := mux.Vars(r)["post_id"]
	ok, err := h.PostStorage.Delete(postId)
	if err != nil || ok == false {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	utils.JSON(w, utils.Message{Message: "success"}, http.StatusOK)
}

func (h *PostHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["post_id"]
	commentId := vars["comment_id"]
	curPost, err := h.PostStorage.GetById(postId)

	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	deleteIndex := -1
	for i, comment := range curPost.Comments {
		if comment.Id == commentId {
			deleteIndex = i
			break
		}
	}

	if deleteIndex == -1 {
		http.Error(w, "no comment", http.StatusInternalServerError)
		return
	}

	if deleteIndex < len(curPost.Comments)-1 {
		copy(curPost.Comments[deleteIndex:], curPost.Comments[deleteIndex+1:])
	}
	curPost.Comments[len(curPost.Comments)-1] = nil
	curPost.Comments = curPost.Comments[:len(curPost.Comments)-1]
	_, err = h.PostStorage.UpdatePost(curPost)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	utils.JSON(w, curPost, http.StatusOK)
}

func (h PostHandler) UpVote(w http.ResponseWriter, r *http.Request) {
	postId := mux.Vars(r)["post_id"]
	curPost, err := h.PostStorage.GetById(postId)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	curUser, err := user.FromContext(r.Context())
	if err != nil {
		http.Error(w, "no auth", http.StatusUnauthorized)
		return
	}

	voteIndex := -1
	for i, vote := range curPost.Votes {
		if vote.UserId == curUser.Id {
			voteIndex = i
			break
		}
	}

	if voteIndex == -1 {
		vote := post.NewVote(curUser.Id, 1)
		curPost.Votes = append(curPost.Votes, vote)
		curPost.Score++
	} else {
		curPost.Votes[voteIndex].Vote = 1
		curPost.Score += 2
	}
	curPost.UpdateVotes()
	_, err = h.PostStorage.UpdatePost(curPost)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	utils.JSON(w, curPost, http.StatusOK)
}

func (h PostHandler) DownVote(w http.ResponseWriter, r *http.Request) {
	postId := mux.Vars(r)["post_id"]
	curPost, err := h.PostStorage.GetById(postId)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	curUser, err := user.FromContext(r.Context())
	if err != nil {
		http.Error(w, "no auth", http.StatusUnauthorized)
		return
	}

	voteIndex := -1
	for i, vote := range curPost.Votes {
		if vote.UserId == curUser.Id {
			voteIndex = i
			break
		}
	}

	if voteIndex == -1 {
		vote := post.NewVote(curUser.Id, -1)
		curPost.Votes = append(curPost.Votes, vote)
		curPost.Score--
	} else {
		curPost.Votes[voteIndex].Vote = -1
		curPost.Score -= 2
	}
	curPost.UpdateVotes()
	_, err = h.PostStorage.UpdatePost(curPost)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	utils.JSON(w, curPost, http.StatusOK)
}

func (h PostHandler) UnVote(w http.ResponseWriter, r *http.Request) {
	postId := mux.Vars(r)["post_id"]
	curPost, err := h.PostStorage.GetById(postId)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	curUser, err := user.FromContext(r.Context())
	if err != nil {
		http.Error(w, "no auth", http.StatusUnauthorized)
		return
	}

	voteIndex := -1
	for i, vote := range curPost.Votes {
		if vote.UserId == curUser.Id {
			voteIndex = i
			break
		}
	}

	if voteIndex == -1 {
		http.Error(w, "no vote", http.StatusInternalServerError)
		return
	} else {
		if curPost.Votes[voteIndex].Vote == 1 {
			curPost.Score--
		} else {
			curPost.Score++
		}

		if voteIndex < len(curPost.Votes)-1 {
			copy(curPost.Votes[voteIndex:], curPost.Votes[voteIndex+1:])
		}
		curPost.Votes[len(curPost.Votes)-1] = nil
		curPost.Votes = curPost.Votes[:len(curPost.Votes)-1]
	}
	curPost.UpdateVotes()
	_, err = h.PostStorage.UpdatePost(curPost)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	utils.JSON(w, curPost, http.StatusOK)
}
