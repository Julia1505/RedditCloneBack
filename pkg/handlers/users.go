package handlers

import (
	"encoding/json"
	"fmt"
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

type Token struct {
	Token string `json:"token"`
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

	newUser, err := h.UserStorage.CreateUser(form.Username, form.Password)
	if err != nil {
		http.Error(w, "Can't sign up", http.StatusUnauthorized)
		return
	}

	token, err := jwt.GenerateJWT(newUser.Id, newUser.Username)
	if err != nil {
		http.Error(w, "Can't create token", http.StatusInternalServerError)
		return
	}

	utils.JSON(w, &Token{Token: token}, http.StatusCreated)
	fmt.Println(h.UserStorage)
}

func (h *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println(h.UserStorage)
	var form SignForm
	err := json.NewDecoder(r.Body).Decode(&form)

	if err != nil {
		http.Error(w, "error in body", http.StatusInternalServerError)
		return
	}

	userIn, err := h.UserStorage.GetUser(form.Username)
	if err != nil {
		utils.ErrorJSON(w, "user not found", 401)
		return
	}

	if !utils.CheckPasswordHash(form.Password, userIn.PasswordHash) {
		http.Error(w, "username or password error", http.StatusUnauthorized)
		return
	}

	token, err := jwt.GenerateJWT(userIn.Id, userIn.Username)
	if err != nil {
		http.Error(w, "Can't create token", http.StatusInternalServerError)
		return
	}

	utils.JSON(w, &Token{Token: token}, http.StatusCreated)

}
