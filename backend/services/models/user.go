package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DBResponse struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	Name            string             `json:"name" bson:"name"`
	Email           string             `json:"email" bson:"email"`
	Password        string             `json:"password" bson:"password"`
	PasswordConfirm string             `json:"password_confirm,omitempty" bson:"password_confirm,omitempty"`
	Phone           string             `json:"phone" bson:"phone"`
	Address         string             `json:"address" bson:"address"`
	CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at" bson:"updated_at"`
	RequestAt       time.Time          `json:"request_at" bson:"request_at"`
	Status          int                `json:"status" bson:"status"`
}

type UserResponse struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Phone     string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Address   string             `json:"address,omitempty" bson:"address,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	RequestAt time.Time          `json:"request_at" bson:"request_at"`
	Status    int                `json:"status" bson:"status"`
}

type SignUpInput struct {
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required,min=8"`
	//PasswordConfirm string    `json:"password_confirm" bson:"password_confirm,omitempty"`
	Phone     string    `json:"phone" bson:"phone"`
	Address   string    `json:"address" bson:"address"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	RequestAt time.Time `json:"update_at" bson:"request_at"`
	Status    int       `json:"status" bson:"status"`
}

type SignInInput struct {
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}
