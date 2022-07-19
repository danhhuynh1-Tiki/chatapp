package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Login struct {
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}

type LoginUsecase interface {
	GetUser(user User) (*User, error)
	UpdateStatusUser(id primitive.ObjectID, t time.Time, status int) error
}
