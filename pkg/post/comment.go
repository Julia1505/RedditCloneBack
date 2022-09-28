package post

import (
	"github.com/Julia1505/RedditCloneBack/pkg/user"
	"time"
)

type Comment struct {
	Author  *user.User `json:"author"`
	Body    string     `json:"body,omitempty"`
	Created time.Time  `json:"created"` // найти норм тип времени
	Id      uint32     `json:"id"`
}
