package usecase

import "chat/domain"

type userUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(userRepo domain.UserRepository) domain.UserUsecase {
	return &userUseCase{userRepo}
}

func (u *userUseCase) GetUser(user domain.User) {
	
}
