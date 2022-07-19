package usecase

import (
	"chat/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (login *loginUsecase) UpdateStatusUser(id primitive.ObjectID, request time.Time, status int) error {
	return login.userRepo.UpdateStatusUser(id, request, status)
}
