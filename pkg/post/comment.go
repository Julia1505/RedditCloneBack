package post

import (
	"github.com/Julia1505/RedditCloneBack/pkg/utils"
	"time"
)

type Comment struct {
	Author  Author    `json:"author" bson:"author"`
	Body    string    `json:"body" bson:"body"`
	Created time.Time `json:"created" bson:"created"`
	Id      string    `json:"id" bson:"id"`
}

func NewComment() *Comment {
	return &Comment{Id: utils.GenarateId(24), Created: time.Now()}
}
