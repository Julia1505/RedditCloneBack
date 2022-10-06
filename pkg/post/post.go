package post

import (
	"github.com/Julia1505/RedditCloneBack/pkg/utils"
	"time"
)

type Post struct {
	Author           `json:"author"`
	Category         string     `json:"category"`
	Comments         []*Comment `json:"comments"`
	Created          time.Time  `json:"created"`
	Id               string     `json:"id"`
	Score            int32      `json:"score"`
	Text             string     `json:"text"`
	Title            string     `json:"title"`
	Type             string     `json:"type"`
	UpvotePersentage int32      `json:"upvotePersentage"`
	Url              string     `json:"url,omitempty"`
	Views            uint32     `json:"views"`
	Votes            []*Vote    `json:"votes"`
}

type Author struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

func (p *Post) UpdateVotes() {
	if len(p.Votes) == 0 {
		return
	}
	p.UpvotePersentage = p.Score * 100 / int32(len(p.Votes))
}

func NewPost() *Post {
	return &Post{Id: utils.GenarateId(24), Created: time.Now()}
}

type PostsRepo interface {
	GetAll() ([]*Post, error)
	GetByCategory(category string) ([]*Post, error)
	GetByUser(username string) ([]*Post, error)
	GetById(id string) (*Post, error)
	AddPost(post *Post) (string, error)
	UpdatePost(post *Post) (*Post, error)
	Delete(id string) (bool, error)
}
