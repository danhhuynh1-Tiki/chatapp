package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	UserID  primitive.ObjectID `json:"user_id" bson:"user_id"`
	Email   string             `json:"email" bson:"email"`
	Content string             `json:"content" bson:"content" binding:"required"`
}

type RoomMessage struct {
	//ID       primitive.ObjectID `json:"id"`
	RoomID   primitive.ObjectID `json:"room_id" bson:"room_id"`
	Messages []Message
}
