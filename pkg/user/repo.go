package user

import (
	"errors"
	"sync"
)

var (
	ErrUserNotExist  = errors.New("User is not exist")
	ErrWrongPassword = errors.New("Wrong password")
	ErrLoginIsBusy   = errors.New("This login is already busy")
	ErrBadPassword   = errors.New("Bad password")
)

type UsersStorage struct {
	lastId uint32
	data   map[string]*User
	mu     sync.RWMutex
}

func NewUsersStorage() *UsersStorage {
	return &UsersStorage{
		lastId: 0,
		data:   make(map[string]*User, 5),
		mu:     sync.RWMutex{},
	}
}

//func (st *UsersStorage) Authorize(username, password string) (*User, error) {
//	st.mu.RLock()
//	defer st.mu.RUnlock()
//
//	user, ok := st.data[username]
//	if !ok {
//		return nil, ErrUserNotExist
//	}
//
//
//	return user, nil
//}

func (st *UsersStorage) CreateUser(username, password string) (*User, error) {
	st.mu.Lock()
	defer st.mu.Unlock()

	//if password == "" { //TODO добавить норм проверку пароля
	//	return nil, ErrBadPassword
	//}

	user := NewUser(username, password)
	st.lastId++
	user.Id = st.lastId
	return user, nil
}

func (st *UsersStorage) GetUser(username string) (*User, error) {
	st.mu.RLock()
	defer st.mu.Unlock()

	user, ok := st.data[username]
	if ok {
		return user, nil
	} else {
		return nil, ErrUserNotExist
	}
}
