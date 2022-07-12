package domain

import (
	"github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte("my_secret_key")

type Claims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}
