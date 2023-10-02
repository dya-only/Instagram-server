package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"go-template/ent/user"
	"os"
	"time"
)

func GenerateToken() string {
	claims := jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := _token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "error"
	}

	return token
}
