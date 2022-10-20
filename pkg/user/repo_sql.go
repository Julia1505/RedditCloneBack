package user

import (
	"database/sql"
	"fmt"
	"github.com/Julia1505/RedditCloneBack/pkg/jwt"
	_ "github.com/go-sql-driver/mysql"
)

type UsersSQL struct {
	DB *sql.DB
}

func NewUsersSQL() *UsersSQL {
	dsn := "root:password@tcp(localhost:3306)/golang?"
	dsn += "&charset=utf8"
	dsn += "&interpolateParams=true"

	db, err := sql.Open("mysql", dsn)
	db.SetMaxOpenConns(10)
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &UsersSQL{DB: db}
}

func (bd *UsersSQL) CreateUser(newUser *User) (*User, error) {
	result, err := bd.DB.Exec(
		"INSERT INTO users (`id`, `username`, `passwordhash`) VALUES (?, ?, ?)",
		newUser.Id,
		newUser.Username,
		newUser.PasswordHash,
	)

	if err != nil {
		return nil, err
	}
	fmt.Println(result.RowsAffected())
	return newUser, nil
}

func (bd *UsersSQL) GetUser(username string) (*User, error) {
	row := bd.DB.QueryRow("SELECT id, username, passwordhash FROM users WHERE username = ?", username)
	curUser := &User{}
	err := row.Scan(&curUser.Id, &curUser.Username, &curUser.PasswordHash)

	if err != nil {
		return nil, err
	}
	return curUser, nil
}

func (bd *UsersSQL) GetByToken(tokenString string) (*User, error) {
	user, err := jwt.ParseToken(tokenString)
	if err != nil {
		return nil, err
	}

	curUser, err := bd.GetUser(user.Username)
	if err != nil {
		return nil, ErrUserNotExist
	}
	return curUser, nil
}
