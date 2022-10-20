package user

import (
	"errors"
	"github.com/Julia1505/RedditCloneBack/pkg/jwt"
	"sync"
)

var (
	ErrUserNotExist = errors.New("User is not exist")
	ErrUnauthorized = errors.New("Unauthorized")
)

type UsersStorage struct {
	data map[string]*User
	mu   sync.RWMutex
}

func NewUsersStorage() *UsersStorage {
	return &UsersStorage{
		data: make(map[string]*User, 5),
		mu:   sync.RWMutex{},
	}
}

func (st *UsersStorage) GetByToken(tokenString string) (*User, error) {
	user, err := jwt.ParseToken(tokenString)
	if err != nil {
		return nil, err
	}

	curUser, err := st.GetUser(user.Username)
	if err != nil {
		return nil, ErrUserNotExist
	}
	return curUser, nil
}

func (st *UsersStorage) CreateUser(newUser *User) (*User, error) {
	st.mu.Lock()
	defer st.mu.Unlock()

	st.data[newUser.Username] = newUser
	return newUser, nil
}

func (st *UsersStorage) GetUser(username string) (*User, error) {
	st.mu.RLock()
	defer st.mu.RUnlock()

	user, ok := st.data[username]
	if ok {
		return user, nil
	} else {
		return nil, ErrUserNotExist
	}
}
