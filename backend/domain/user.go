package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id" `
	Password string             `json:"password" bson:"password"`
	// Name     string             `json:"password" bson:"password"`
	Email string `json:"email" bson:"email"`
}

type UserRepository interface {
	GetUser(user Login) (*User, error)
}

type UserUsecase interface {
	GetUser(user User)
}
