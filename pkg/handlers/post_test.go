package handlers

import (
	"bytes"
	"fmt"
	"github.com/Julia1505/RedditCloneBack/pkg/post"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostsHandlerList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := post.NewMockPostsRepo(ctrl)
	postHandler := &PostHandler{
		PostStorage: repo,
	}

	t.Run("positive test", func(t *testing.T) {
		resultPosts := []*post.Post{
			{
				Author:           post.Author{Id: "12", Username: "sfd"},
				Category:         "music",
				Comments:         nil,
				Id:               bson.ObjectId("fsd"),
				Score:            23,
				Text:             "fsdsd",
				Title:            "sdas",
				UpvotePercentage: 23,
				Views:            4,
				Type:             "text",
				Votes:            []*post.Vote{{UserId: "34", Vote: 1}},
			},
			{
				Author:   post.Author{Id: "23", Username: "fsd"},
				Category: "dsa",
				Type:     "text",
			},
		}

		repo.EXPECT().GetAll().Return(resultPosts, nil)

		req := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()

		postHandler.List(w, req)
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		if !bytes.Contains(body, []byte(resultPosts[0].Id)) ||
			!bytes.Contains(body, []byte(resultPosts[0].Username)) ||
			!bytes.Contains(body, []byte(resultPosts[0].Category)) ||
			!bytes.Contains(body, []byte(resultPosts[0].Title)) {
			t.Errorf("no data found")
			return
		}
	})

	t.Run("negative test", func(t *testing.T) {
		repo.EXPECT().GetAll().Return(nil, fmt.Errorf("no results"))
		req := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		postHandler.List(w, req)

		resp := w.Result()
		if resp.StatusCode != http.StatusInternalServerError {
			t.Errorf("expected 500 got %v", resp.StatusCode)
			return
		}
	})

}

func TestPostHandlerCategoryList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := post.NewMockPostsRepo(ctrl)
	postHandler := &PostHandler{
		PostStorage: repo,
	}

	t.Run("positive test", func(t *testing.T) {
		resultPosts := []*post.Post{
			{
				Author:           post.Author{Id: "12", Username: "sfd"},
				Category:         "music",
				Comments:         nil,
				Id:               bson.ObjectId("fsd"),
				Score:            23,
				Text:             "fsdsd",
				Title:            "sdas",
				UpvotePercentage: 23,
				Views:            4,
				Type:             "text",
				Votes:            []*post.Vote{{UserId: "34", Vote: 1}},
			},
		}

		category := "music"

		repo.EXPECT().GetByCategory(category).Return(resultPosts, nil)

		req := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()

		vars := map[string]string{
			"category_name": category,
		}
		req = mux.SetURLVars(req, vars)

		postHandler.CategoryList(w, req)
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		if !bytes.Contains(body, []byte(resultPosts[0].Id)) ||
			!bytes.Contains(body, []byte(resultPosts[0].Username)) ||
			!bytes.Contains(body, []byte(resultPosts[0].Category)) ||
			!bytes.Contains(body, []byte(resultPosts[0].Title)) {
			t.Errorf("no data found")
			return
		}
	})

	t.Run("no category", func(t *testing.T) {

		category := "music"

		repo.EXPECT().GetByCategory(category).Return(nil, fmt.Errorf("no data"))

		req := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()

		vars := map[string]string{
			"category_name": category,
		}
		req = mux.SetURLVars(req, vars)

		postHandler.CategoryList(w, req)
		resp := w.Result()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200 got %v", resp.StatusCode)
			return
		}

		//if resp.Body != nil {
		//	t.Errorf("expected empty")
		//	return
		//}

	})
}
