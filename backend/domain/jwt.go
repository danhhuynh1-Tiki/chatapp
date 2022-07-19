package domain

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var JwtKey = []byte("chat-app")

type Claims struct {
	ID    primitive.ObjectID `json:"id"`
	Email string             `json:"email"`
	jwt.StandardClaims
}
