package usecase

import (
	"chat/services/models"
	"chat/services/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageUseCase interface {
	AddMessage(primitive.ObjectID, primitive.ObjectID, string, string) (*models.RoomMessage, error)
}

type messageUseCase struct {
	messageRepo repository.MessageRepository
}

func NewMessageUseCase(messageRepo repository.MessageRepository) MessageUseCase {
	return &messageUseCase{messageRepo: messageRepo}
}
func (m *messageUseCase) AddMessage(room_id, user_id primitive.ObjectID, email string, content string) (*models.RoomMessage, error) {
	return m.messageRepo.AddMessage(room_id, user_id, email, content)
}
