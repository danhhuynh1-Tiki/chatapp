package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomRepository interface {
	CreateRoom(primitive.ObjectID, primitive.ObjectID) (string, error)
	GetRoom() error
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
	//r1 := models.Room{}
	res, err := r.collection.InsertOne(r.context, bson.D{
		{"name", ""},
	})

	if err != nil {
		return "", err
	}
	room_id := res.InsertedID
	// insert two document into collection room
	docs := []interface{}{
		bson.D{{"room_id", room_id}, {"user_id", id}},
		bson.D{{"room_id", room_id}, {"user_id", id1}},
	}
	//opts := options.InsertMany().SetOrdered(false)
	result, err := r.collection.InsertMany(r.context, docs)

	if err != nil {
		return "", err
	}
	fmt.Println(result.InsertedIDs)
	//fmt.Println(res.InsertedID.(primitive.ObjectID).Hex())
	return "", nil

}
func (r *roomRepository) CreateGroup() error {
	return nil
}
func (r *roomRepository) GetRoom() error {
	return nil
}
