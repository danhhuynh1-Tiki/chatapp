package usecase

import (
	"chat/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type statusUsecase struct {
	userRepo domain.UserRepository
}

func NewStatusUsecase(userRepo domain.UserRepository) domain.StatusUsecase {
	return &statusUsecase{
		userRepo: userRepo,
	}
}

func (s *statusUsecase) UpdateStatusUser(id primitive.ObjectID, t time.Time, status int) error {
	return s.userRepo.UpdateStatusUser(id, t, status)
}
