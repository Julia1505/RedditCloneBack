package jwt

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

var (
	SecretToken = "qwertyuiop"
)

type Token struct {
	Username    string `json:"username"`
	TokenString string `json:"token"`
}

func GenerateJWT(username string) (string, error) {
	var SigningKey = []byte(SecretToken)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	tokenString, err := token.SignedString(SigningKey)

	if err != nil {
		fmt.Errorf("Jwt error: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
