package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// create single chat
type Room struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
}

type RoomMembers struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	RoomId  primitive.ObjectID `json:"room_id" bson:"room_id"`
	UserId1 primitive.ObjectID `json:"user_id1" bson:"user_id1"`
	UserId2 primitive.ObjectID `json:"user_id2" bson:"user_id2"`
}

//create group chat

type GroupMemebrs struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	RoomID primitive.ObjectID `json:"room_id" bson:"room_id"`
	Email  primitive.ObjectID `json:"room_id" bson`
}
