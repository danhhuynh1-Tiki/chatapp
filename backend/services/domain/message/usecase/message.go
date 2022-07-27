package usecase

import (
	"chat/services/models"
	"chat/services/repository"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageUseCase interface {
	AddMessage(primitive.ObjectID, models.Message) (*models.RoomMessage, error)
	GetMessage(primitive.ObjectID) (*models.RoomMessage, error)
}

type messageUseCase struct {
	messageRepo repository.MessageRepository
	userRepo    repository.UserRepository
}

func NewMessageUseCase(messageRepo repository.MessageRepository, userRepo repository.UserRepository) MessageUseCase {
	return &messageUseCase{messageRepo, userRepo}
}
func (m *messageUseCase) GetMessage(room_id primitive.ObjectID) (*models.RoomMessage, error) {
	return m.messageRepo.GetMessage(room_id)

}
func (m *messageUseCase) AddMessage(room_id primitive.ObjectID, message models.Message) (*models.RoomMessage, error) {
	fmt.Println(message.UserID)
	query := bson.M{
		"_id": message.UserID,
	}
	//fmt.Println(query)
	user, _ := m.userRepo.FindByQuery(query)
	message.Email = user.Email
	fmt.Println("usecase message : ", message)
	return m.messageRepo.AddMessage(room_id, message.UserID, message.Email, message.Content)
	//return nil, nil
}
