package usecase

import (
	"chat/services/models"
	"chat/services/repository"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoomUsecase interface {
	CreateRoom(id primitive.ObjectID, objectID primitive.ObjectID) (string, error)
	CreateGroup(string, []models.GroupMembers) (string, error)
	GetGroup(string) ([]models.GroupMembers, error)
}
type roomUsecase struct {
	roomRepo repository.RoomRepository
	userRepo repository.UserRepository
}

func NewRoomUsecase(roomRepo repository.RoomRepository, useRepo repository.UserRepository) RoomUsecase {
	return &roomUsecase{roomRepo, useRepo}
}

func (r *roomUsecase) CreateRoom(id1 primitive.ObjectID, id2 primitive.ObjectID) (string, error) {
	return r.roomRepo.CreateRoom(id1, id2)
}
func (r *roomUsecase) CreateGroup(name string, groupMembers []models.GroupMembers) (string, error) {
	for i := 0; i < len(groupMembers); i++ {
		query := primitive.M{
			"email": groupMembers[i].Email,
		}
		_, err := r.userRepo.FindByQuery(query)
		if err != nil {
			return "", errors.New("Email is not exists")
		}

	}
	return r.roomRepo.CreateGroup(name, groupMembers)
}
func (r *roomUsecase) GetGroup(email string) ([]models.GroupMembers, error) {
	return r.roomRepo.GetGroup(email)
}
