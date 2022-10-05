package jwt

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

var (
	secretToken = []byte("qwertyuiop")
)

type UserInfo struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type UserClaims struct {
	User *UserInfo `json:"user"`
	jwt.StandardClaims
}

func GenerateJWT(id, username string) (string, error) {
	now := time.Now()
	claims := &UserClaims{
		User: &UserInfo{
			Username: username,
			ID:       id,
		},
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: now.AddDate(0, 1, 0).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretToken)

	if err != nil {
		fmt.Errorf("Jwt error: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*UserInfo, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretToken, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims.User, nil
	}
	return nil, err
}
