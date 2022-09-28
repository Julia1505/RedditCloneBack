package handlers

import (
	"encoding/json"
	"github.com/Julia1505/RedditCloneBack/pkg/post"
	"net/http"
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
	json.NewEncoder(w).Encode(elems)
}
