package models

const DB_NAME = "user"

/**
 * Define the user model
 */
type User struct {
	Email    string             `json:"email,omitempty" validate:"email,required"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Password string             `json:"password,omitempty" validate:"required"`
	Phone    string             `json:"phone,omitempty"`
	Address  string             `json:"address,omitempty"`
}

type UserRepository interface {
	Create(User) (*string, error) // used to create a new user
	IsExisted(string) (*User, error) // used to check if the user is existed
}
