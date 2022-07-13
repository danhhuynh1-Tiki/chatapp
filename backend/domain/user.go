package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// type User struct {
// 	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id" `
// 	Password string             `json:"password" bson:"password"`
// 	// Name     string             `json:"password" bson:"password"`
// 	Email string `json:"email" bson:"email"`
// }
type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Email    string             `json:"email,omitempty" binding:"email,required"`
	Name     string             `json:"name,omitempty" `
	Password string             `json:"password,omitempty" binding:"required"`
	Phone    string             `json:"phone,omitempty"`
	Address  string             `json:"address,omitempty"`
}

type UserRepository interface {
	Create(user User) (*string, error) // used to create a new user
	// IsExisted(email string) (*User, error) // used to check if the user is existed
	GetUser(user User) (*User, error)
}
type UserUsecase interface {
	Create(user User) (*string, error)
	// IsExisted(email string) (*User, error)
	GetUser(user User)
}
