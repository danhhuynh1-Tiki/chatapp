package usecase

import "chat/services/repository"

type RoomUsecase interface {
	CreateRoom() error
}
type roomUsecase struct {
	roomRepo repository.RoomRepository
}

func NewRoomUsecase(roomRepo repository.RoomRepository) RoomUsecase {
	return &roomUsecase{roomRepo}
}

func (r *roomUsecase) CreateRoom() error {
	return r.roomRepo.CreateRoom()
}
