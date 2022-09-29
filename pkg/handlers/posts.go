package handlers

import (
	"encoding/json"
	"github.com/Julia1505/RedditCloneBack/pkg/post"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

type PostHandler struct {
	PostStorage *post.PostsStorage
}

func (h *PostHandler) List(w http.ResponseWriter, r *http.Request) {
	elems, err := h.PostStorage.GetAll()

	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(elems)
}

func (h *PostHandler) CategoryList(w http.ResponseWriter, r *http.Request) {
	category, ok := mux.Vars(r)["category_name"]
	if !ok {
		http.Error(w, "no category", http.StatusInternalServerError)
		return
	}
	//fmt.Println(category)
	elems, _ := h.PostStorage.GetByCategory(category)

	//if err != nil {
	//	http.Error(w, "DB error", http.StatusInternalServerError)
	//	return
	//}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(elems)
}

func (h *PostHandler) Post(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["post_id"])
	if err != nil {
		http.Error(w, "uncorrect id", http.StatusInternalServerError)
		//fmt.Println(err)
		return
	}

	elems, err := h.PostStorage.GetById(uint32(id))

	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(elems)
}

func (h *PostHandler) AddPost(w http.ResponseWriter, r *http.Request) {
	elem := &post.Post{}
	elem.Type = r.FormValue("type")
	switch elem.Type {
	case "text":
		elem.Text = r.FormValue("text")
	case "url":
		elem.Url = r.FormValue("url")
	default:
		http.Error(w, "unknown type of post", http.StatusInternalServerError)
		return
	}
	elem.Title = r.FormValue("title")
	elem.Category = r.FormValue("category")
	elem.Created = time.Now()
	_, err := h.PostStorage.AddPost(elem)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(elem)
}
