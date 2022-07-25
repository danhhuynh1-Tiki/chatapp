package usecase

import (
	"chat/services/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoomUsecase interface {
	CreateRoom(id primitive.ObjectID, objectID primitive.ObjectID) (string, error)
}
type roomUsecase struct {
	roomRepo repository.RoomRepository
}

func NewRoomUsecase(roomRepo repository.RoomRepository) RoomUsecase {
	return &roomUsecase{roomRepo}
}

func (r *roomUsecase) CreateRoom(id1 primitive.ObjectID, id2 primitive.ObjectID) (string, error) {
	return r.roomRepo.CreateRoom(id1, id2)
}
