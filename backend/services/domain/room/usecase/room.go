package usecase

import (
	"chat/services/models"
	"chat/services/repository"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoomUsecase interface {
	CreateRoom(id primitive.ObjectID, objectID primitive.ObjectID) (string, error)
	CreateGroup(string, string, []models.GroupMembers) (string, error)
	GetGroup(string) ([]models.GroupMembers, error)
	GetMembers(primitive.ObjectID) ([]string, error)
	RemoveMember(primitive.ObjectID, string) error
	AddMember(primitive.ObjectID, string) error
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
func (r *roomUsecase) CreateGroup(name string, admin string, groupMembers []models.GroupMembers) (string, error) {
	for i := 0; i < len(groupMembers); i++ {
		query := primitive.M{
			"email": groupMembers[i].Email,
		}
		_, err := r.userRepo.FindByQuery(query)
		if err != nil {
			return "", errors.New("Email is not exists")
		}

	}
	return r.roomRepo.CreateGroup(name, admin, groupMembers)
}
func (r *roomUsecase) GetGroup(email string) ([]models.GroupMembers, error) {
	return r.roomRepo.GetGroup(email)
}
func (r *roomUsecase) GetMembers(room_id primitive.ObjectID) ([]string, error) {
	return r.roomRepo.GetMembers(room_id)
}
func (r *roomUsecase) RemoveMember(room_id primitive.ObjectID, email string) error {
	return r.roomRepo.RemoveMember(room_id, email)
}
func (r *roomUsecase) AddMember(room_id primitive.ObjectID, email string) error {
	query := primitive.M{
		"email": email,
	}
	_, err := r.userRepo.FindByQuery(query)

	if err != nil {
		return err
	}
	return r.roomRepo.AddMember(room_id, email)
}
