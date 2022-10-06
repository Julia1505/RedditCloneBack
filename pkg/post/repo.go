package post

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	ErrNotFoundPost    = errors.New("Not found post")
	ErrInvalidId       = errors.New("Invalid Id")
	ErrInvalidCategory = errors.New("Invalid category")
	ErrInvalidUsername = errors.New("Invalid username")
)

type PostsStorage struct {
	data []*Post
	mu   sync.RWMutex
}

func NewPostsStorage() *PostsStorage {
	return &PostsStorage{
		data: make([]*Post, 0, 10),
		mu:   sync.RWMutex{},
	}
}

func (st *PostsStorage) GetAll() ([]*Post, error) {
	st.mu.RLock()
	defer st.mu.RUnlock()
	return st.data, nil
}

func (st *PostsStorage) GetByCategory(category string) ([]*Post, error) {
	posts := make([]*Post, 0, 10)
	st.mu.RLock()
	defer st.mu.RUnlock()

	for _, post := range st.data {
		if post.Category == category {
			posts = append(posts, post)
		}
	}

	if len(posts) == 0 {
		return posts, ErrInvalidCategory
	}

	return posts, nil
}

func (st *PostsStorage) GetByUser(username string) ([]*Post, error) {
	posts := make([]*Post, 0, 10)
	st.mu.RLock()
	defer st.mu.RUnlock()

	for _, post := range st.data {
		if post.Author.Username == username {
			posts = append(posts, post)
		}
	}

	if len(posts) == 0 {
		return posts, ErrInvalidUsername
	}

	return posts, nil
}

func (st *PostsStorage) GetById(id string) (*Post, error) {
	st.mu.RLock()
	defer st.mu.RUnlock()

	for _, post := range st.data {
		if post.Id == id {
			atomic.AddUint32(&post.Views, 1)
			return post, nil
		}
	}

	return nil, ErrInvalidId
}

func (st *PostsStorage) AddPost(post *Post) (string, error) {
	st.mu.Lock()
	defer st.mu.Unlock()
	st.data = append(st.data, post)
	return post.Id, nil
}

func (st *PostsStorage) UpdatePost(newPost *Post) (*Post, error) {
	st.mu.Lock()
	defer st.mu.Unlock()

	for i, post := range st.data {
		if post.Id == newPost.Id {
			st.data[i] = newPost
			return st.data[i], nil
		}
	}
	return nil, ErrNotFoundPost
}

func (st *PostsStorage) Delete(id string) (bool, error) {
	deleteInd := -1
	st.mu.Lock()
	defer st.mu.Unlock()

	for i, post := range st.data {
		if post.Id == id {
			deleteInd = i
			break
		}
	}

	if deleteInd == -1 {
		return false, ErrNotFoundPost
	}

	if deleteInd < len(st.data)-1 {
		copy(st.data[deleteInd:], st.data[deleteInd+1:])
	}
	st.data[len(st.data)-1] = nil
	st.data = st.data[:len(st.data)-1]
	return true, nil
}
