package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Room struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
}

type RoomMembers struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	RoomId primitive.ObjectID `json:"room_id" bson:"room_id"`
	UserId primitive.ObjectID `json:"user_id" bson:"user_id"`
}
