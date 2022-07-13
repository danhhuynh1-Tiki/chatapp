package user

import (
	"chat/domain"
)

// type UserUsecase interface {
// 	Create(domain.User) (*string, error)
// 	IsExisted(string) (*domain.User, error)
// }

type userUsecase struct {
	repo domain.UserRepository
}

func NewSignupUsecase(repo domain.UserRepository) domain.SingupUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (userUC userUsecase) Create(user domain.User) (*string, error) {
	return userUC.repo.Create(user)
}

// func (userUC userUsecase) IsExisted(email string) (*domain.User, error) {
// 	return userUC.repo.IsExisted(email)
// }
