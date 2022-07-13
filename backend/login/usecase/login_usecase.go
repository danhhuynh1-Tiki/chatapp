package usecase

import (
	"chat/domain"
)

type loginUsecase struct {
	userRepo domain.UserRepository
}

func NewLoginUsecase(userRepo domain.UserRepository) domain.LoginUsecase {
	return &loginUsecase{userRepo}
}

func (login *loginUsecase) GetUser(user domain.User) (*domain.User, error) {
	users, err := login.userRepo.GetUser(user)
	if err != nil {
		return nil, err
	} else {
		return users, nil
	}
}
