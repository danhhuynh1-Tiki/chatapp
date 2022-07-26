package repository

import (
	"chat/services/models"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageRepository interface {
	AddMessage(primitive.ObjectID, primitive.ObjectID, string, string) (*models.RoomMessage, error)
	FindRoomMessage(primitive.ObjectID) (*models.RoomMessage, error)
}

type messageRepository struct {
	context    context.Context
	collection *mongo.Collection
}

func NewMessageRepository(context context.Context, collection *mongo.Collection) MessageRepository {
	return &messageRepository{context, collection}
}

func (m *messageRepository) AddMessage(room_id, user_id primitive.ObjectID, email string, content string) (*models.RoomMessage, error) {
	roomM, err := m.FindRoomMessage(room_id)
	fmt.Println("Room message repository", roomM)
	fmt.Println("err mesage repository", err)
	if err != nil {

		message := models.Message{user_id, email, content}

		roomM = &models.RoomMessage{RoomID: room_id, Messages: []models.Message{message}}
		//fmt.Println(*roomM)
		_, err := m.collection.InsertOne(m.context, *roomM)

		if err != nil {
			return nil, err
		}
		fmt.Println(*roomM)
		return roomM, nil
	} else {
		message := models.Message{user_id, email, content}

		roomM.Messages = append(roomM.Messages, message)

		res, err := m.collection.UpdateOne(m.context, bson.D{{"room_id", room_id}}, bson.D{
			{"$set", roomM},
		})
		if err != nil || res.ModifiedCount == 0 {
			return nil, err
		} else {
			return roomM, nil
		}
	}
	return nil, nil
}

func (m *messageRepository) FindRoomMessage(room_id primitive.ObjectID) (*models.RoomMessage, error) {
	var roomM models.RoomMessage
	err := m.collection.FindOne(m.context, bson.D{
		{"room_id", room_id},
	}).Decode(&roomM)
	if err != nil {
		return nil, errors.New("room_message is not exists")
	}
	return &roomM, nil
}
