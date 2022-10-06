package post

import (
	"github.com/Julia1505/RedditCloneBack/pkg/utils"
	"time"
)

type Comment struct {
	Author  Author    `json:"author"`
	Body    string    `json:"body"`
	Created time.Time `json:"created"`
	Id      string    `json:"id"`
}

func NewComment() *Comment {
	return &Comment{Id: utils.GenarateId(24), Created: time.Now()}
}
