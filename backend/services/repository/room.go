package repository

import (
	"chat/services/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomRepository interface {
	CreateRoom(primitive.ObjectID, primitive.ObjectID) (string, error)
	GetRoom(primitive.ObjectID, primitive.ObjectID) (string, error)
}
type roomRepository struct {
	context    context.Context
	collection *mongo.Collection
}

func NewRoomRepository(context context.Context, collection *mongo.Collection) RoomRepository {
	return &roomRepository{
		context, collection,
	}
}

func (r *roomRepository) CreateRoom(id primitive.ObjectID, id1 primitive.ObjectID) (string, error) {
	fmt.Println(id, id1)
	room_id, err := r.GetRoom(id, id1)
	//fmt.Println("room_id", room_id)
	if err == nil {
		return room_id, nil
	}
	room_id1, err := r.GetRoom(id1, id)
	//fmt.Println("room_id1", room_id1)
	if err == nil {
		return room_id1, nil
	}

	res, err := r.collection.InsertOne(r.context, bson.D{
		{"name", ""},
	})

	if err != nil {
		return "", err
	}
	roomid := res.InsertedID

	result, err := r.collection.InsertOne(r.context, bson.D{
		{"room_id", roomid},
		{"user_id1", id},
		{"user_id2", id1},
	})

	if err != nil {
		return "", err
	}
	fmt.Println("", result.InsertedID)
	//fmt.Println(res.InsertedID.(primitive.ObjectID).Hex())
	return roomid.(primitive.ObjectID).Hex(), nil

}
func (r *roomRepository) CreateGroup() error {
	return nil
}
func (r *roomRepository) GetRoom(id, id1 primitive.ObjectID) (string, error) {
	var roomM models.RoomMembers
	err := r.collection.FindOne(r.context, bson.D{
		{"user_id1", id},
		{"user_id2", id1},
	}).Decode(&roomM)

	if err != nil {
		return "", err
	}
	return roomM.RoomId.Hex(), err
}
