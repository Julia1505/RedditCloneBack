package post

import (
	"github.com/Julia1505/RedditCloneBack/pkg/user"
	"time"
)

//import "RedditCloneBack/pkg/user"

type Post struct {
	Id               uint32       `json:"id"`
	Author           user.User    `json:"author"`
	Category         string       `json:"category"`
	Comments         []*Comment   `json:"comments"`
	Title            string       `json:"title"`
	Created          time.Time    `json:"created"` // мб исправить
	Score            uint32       `json:"score"`
	Type             string       `json:"type"`
	Text             string       `json:"text,omitempty"`
	UpvotePersentage uint32       `json:"upvotePersentage"`
	Url              string       `json:"url,omitempty"`
	Views            uint32       `json:"views"`
	Votes            []*user.User `json:"votes"`
}

func NewPost(au string, cat string, tit string) *Post {
	return &Post{Author: user.User{Id: "1", Username: au}, Category: cat, Title: tit}
}

type PostsRepo interface {
	GetAll() ([]*Post, error)
	GetByCategory(category string) ([]*Post, error)
	GetByUser(username string) ([]*Post, error)
	GetById(id uint32) (*Post, error)
	AddPost(post *Post) (uint32, error)
	UpdatePost(post *Post) (bool, error)
	Delete(id uint32) (bool, error)
}
