package handlers

import (
	"encoding/json"
	"github.com/Julia1505/RedditCloneBack/pkg/jwt"
	"github.com/Julia1505/RedditCloneBack/pkg/user"
	"github.com/Julia1505/RedditCloneBack/pkg/utils"
	"net/http"
)

type UserHandler struct {
	UserStorage user.UsersRepo
}

type SignForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var form SignForm
	err := json.NewDecoder(r.Body).Decode(&form)
	defer r.Body.Close() // нужно ли??

	if err != nil {
		http.Error(w, "error in body", http.StatusInternalServerError)
		return
	}

	_, err = h.UserStorage.GetUser(form.Username)
	if err == nil {
		http.Error(w, "username is busy", http.StatusUnauthorized)
		return
	}

	form.Password, err = utils.GenarateHashPassword(form.Password)
	if err != nil {
		http.Error(w, "generate hash error", http.StatusInternalServerError)
		return
	}
	newUser := user.NewUser(form.Username, form.Password)

	newUser, err = h.UserStorage.CreateUser(newUser)
	if err != nil {
		http.Error(w, "Can't sign up", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json") // TODO вынести в отдельную функцию с проверкой на ошибку
	json.NewEncoder(w).Encode(newUser)
}

func (h *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var form SignForm
	err := json.NewDecoder(r.Body).Decode(&form)

	if err != nil {
		http.Error(w, "error in body", http.StatusInternalServerError)
		return
	}

	userIn, err := h.UserStorage.GetUser(form.Username)
	if err != nil {
		http.Error(w, "The user doesn't exist", http.StatusUnauthorized)
		return
	}

	if !utils.CheckPasswordHash(form.Password, userIn.PasswordHash) {
		http.Error(w, "username or password error", http.StatusUnauthorized)
		return
	}

	validToken, err := jwt.GenerateJWT(userIn.Username)
	if err != nil {
		http.Error(w, "err token", http.StatusUnauthorized)
		return
	}

	var token jwt.Token

	token.Username = userIn.Username
	token.TokenString = validToken

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}
