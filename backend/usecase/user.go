package user

import (
	model "jwt-demo/models"
)

type UserUsecase interface {
	Create(model.User) (*string, error)
	IsExisted(string) (*model.User, error)
}

type userUsecase struct {
	repo model.UserRepository
}

func NewUserService(repo model.UserRepository) UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (userUC userUsecase) Create(user model.User) (*string, error) {
	return userUC.repo.Create(user)
}

func (userUC userUsecase) IsExisted(email string) (*model.User, error) {
	return userUC.repo.IsExisted(email)
}