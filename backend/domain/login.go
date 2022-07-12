package domain

type Login struct {
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}

type LoginUsecase interface {
	GetUser(user Login) (*User, error)
}
