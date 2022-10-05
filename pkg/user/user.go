package user

import "context"

type User struct {
	Id           string `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"_"`
}

func NewUser(username, passwordHash string) *User {
	return &User{
		Username:     username,
		PasswordHash: passwordHash,
	}
}

func FromContext(ctx context.Context) (*User, error) {
	user, ok := ctx.Value("user").(*User)
	if ok {
		return user, nil
	}
	return nil, ErrUnauthorized
}

type UsersRepo interface {
	CreateUser(username, password string) (*User, error)
	GetUser(username string) (*User, error)
	GetByToken(tokenString string) (*User, error)
}
