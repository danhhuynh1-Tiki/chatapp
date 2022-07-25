package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomRepository interface {
	CreateRoom() error
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

func (r *roomRepository) CreateRoom() error {

	return nil
}
func (r *roomRepository) CreateGroup() error {
	return nil
}
func (r *roomRepository) GetRoom() error {
	return nil
}
