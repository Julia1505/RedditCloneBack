package post

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Post struct {
	Author           `json:"author" bson:"author"`
	Category         string        `json:"category" bson:"category"`
	Comments         []*Comment    `json:"comments" bson:"comments"`
	Created          time.Time     `json:"created" bson:"created"`
	Id               bson.ObjectId `json:"id" bson:"id"`
	Score            int32         `json:"score" bson:"score"`
	Text             string        `json:"text" bson:"text"`
	Title            string        `json:"title" bson:"title"`
	Type             string        `json:"type" bson:"type"`
	UpvotePercentage int32         `json:"upvotePercentage" bson:"upvotePercentage"`
	Url              string        `json:"url,omitempty" bson:"url"`
	Views            uint32        `json:"views" bson:"views"`
	Votes            []*Vote       `json:"votes" bson:"votes"`
}

type Author struct {
	Id       string `json:"id"`
	Username string `json:"username" bson:"username"`
}

func (p *Post) UpdateVotes() {
	if len(p.Votes) == 0 {
		return
	}
	p.UpvotePercentage = p.Score * 100 / int32(len(p.Votes))
}

func NewPost() *Post {
	return &Post{Id: bson.NewObjectId(), Created: time.Now()}
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
