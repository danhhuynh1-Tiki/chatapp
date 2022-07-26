package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"

	"chat/services/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	FindByQuery(primitive.M) (*models.DBResponse, error)
	GetAll(primitive.M) ([]models.DBResponse, error)
	UpdateStatus(primitive.M, int) error
	FilterUser(primitive.ObjectID, int64) ([]models.DBResponse, error)
	GetUser(id primitive.ObjectID) (models.DBResponse, error)
}

type userRepository struct {
	context    context.Context
	collection *mongo.Collection
}

func NewUserRepository(context context.Context, collection *mongo.Collection) UserRepository {
	return &userRepository{context, collection}
}

func (repository *userRepository) FindByQuery(query primitive.M) (*models.DBResponse, error) {
	var user *models.DBResponse
	err := repository.collection.FindOne(repository.context, query).Decode(&user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repository *userRepository) GetAll(query primitive.M) ([]models.DBResponse, error) {
	var alluser []models.DBResponse
	cursor, err := repository.collection.Find(repository.context, query)
	for cursor.Next(repository.context) {
		var user models.DBResponse
		cursor.Decode(&user)
		alluser = append(alluser, user)
	}
	if err != nil {
		return nil, err
	} else {
		return alluser, nil
	}
}

func (repository *userRepository) UpdateStatus(query primitive.M, status int) error {
	t := time.Now()

	res, err := repository.collection.UpdateOne(repository.context, query, bson.D{
		{"$set", bson.D{
			{"status", status},
			{"request_at", t},
		}},
	})
	if res.ModifiedCount == 0 || err != nil {
		return err
	}
	return nil
}

// pagination users
func (repository *userRepository) FilterUser(id primitive.ObjectID, c int64) ([]models.DBResponse, error) {
	option := options.Find()
	option.SetLimit(c)

	cursor, err := repository.collection.Find(repository.context, bson.M{"_id": bson.M{
		"$ne": id,
	}}, option)
	if err != nil {
		return nil, err
	}
	var filter []models.DBResponse
	for cursor.Next(repository.context) {
		var user models.DBResponse
		cursor.Decode(&user)

		filter = append(filter, user)
	}
	return filter, nil
}

func (u *userRepository) GetUser(id primitive.ObjectID) (models.DBResponse, error) {
	var user models.DBResponse
	err := u.collection.FindOne(u.context, bson.M{
		"_id": id,
	}).Decode(&user)
	if err != nil {
		return models.DBResponse{}, nil
	}
	return user, nil
}
