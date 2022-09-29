package user

type User struct {
	Id           uint32 `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"_"`
}

func NewUser(username, passwordHash string) *User {
	return &User{
		Username:     username,
		PasswordHash: passwordHash,
	}
}

type UsersRepo interface {
	//Authorize(username, password string) (*User, error)
	CreateUser(user *User) (*User, error)
	GetUser(username string) (*User, error)
}
