package post

import (
	"time"
)

//import "RedditCloneBack/pkg/user"

type Post struct {
	Id               string `json:"id"`
	Author           `json:"author"`
	Category         string     `json:"category"`
	Comments         []*Comment `json:"comments"`
	Title            string     `json:"title"`
	Created          time.Time  `json:"created"` // мб исправить
	Score            int32      `json:"score"`
	Type             string     `json:"type"`
	Text             string     `json:"text"`
	UpvotePersentage uint32     `json:"upvotePersentage"`
	Url              string     `json:"url,omitempty"`
	Views            uint32     `json:"views"`
	Votes            []*Vote    `json:"votes"`
}

type Author struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

//func NewPost(au string, cat string, tit string) *Post {
//	return &Post{Author: user.User{Id: 1, Username: au}, Category: cat, Title: tit}
//}

type PostsRepo interface {
	GetAll() ([]*Post, error)
	GetByCategory(category string) ([]*Post, error)
	GetByUser(username string) ([]*Post, error)
	GetById(id string) (*Post, error)
	AddPost(post *Post) (string, error)
	UpdatePost(post *Post) (*Post, error)
	Delete(id string) (bool, error)
}
