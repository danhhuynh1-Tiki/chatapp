package domain

import (
	"github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte("chat-app")

type Claims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}
