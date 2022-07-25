package repository

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"chat/services/models"
)

type AuthRepository interface {
	AddUser(*models.SignUpInput) (*mongo.InsertOneResult, error)
	IndexesUser() (string, error)
	GetUser(primitive.M) (*models.DBResponse, error)
}

type authRepository struct {
	context    context.Context
	collection *mongo.Collection
}

func NewAuthRepository(context context.Context, collection *mongo.Collection) AuthRepository {
	return &authRepository{context, collection}
}

func (repository *authRepository) AddUser(user *models.SignUpInput) (*mongo.InsertOneResult, error) {
	fmt.Println("auth repo", user)
	res, err := repository.collection.InsertOne(repository.context, user)

	return res, err
}

func (repository *authRepository) IndexesUser() (string, error) {
	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: opt}

	if _, err := repository.collection.Indexes().CreateOne(repository.context, index); err != nil {
		return "", errors.New("could not create index for email")
	}

	return "index created", nil
}

func (repository *authRepository) GetUser(query primitive.M) (*models.DBResponse, error) {
	var user *models.DBResponse
	err := repository.collection.FindOne(repository.context, query).Decode(&user)

	if err != nil {
		return nil, err
	}

	return user, nil

}