package post

import (
	"time"
)

type Comment struct {
	Author  Author    `json:"author"`
	Body    string    `json:"body"`
	Created time.Time `json:"created"` // найти норм тип времени
	Id      string    `json:"id"`
}
