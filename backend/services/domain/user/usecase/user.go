package usecase

import (
	"chat/services/models"
	"chat/services/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserUseCase interface {
	FindById(string) (*models.DBResponse, error)
	FindByEmail(string) (*models.DBResponse, error)
	GetAll() ([]models.DBResponse, error)
	UpdateStatus(primitive.ObjectID, int) error
	FilterUser(primitive.ObjectID, int64) ([]models.DBResponse, error)
	GetUser(primitive.ObjectID) (models.DBResponse, error)
}

type userUseCase struct {
	repository repository.UserRepository
}

func NewUserUseCase(repository repository.UserRepository) UserUseCase {
	return &userUseCase{
		repository: repository,
	}
}

func (usecase *userUseCase) FindById(id string) (*models.DBResponse, error) {
	oid, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": oid}
	user, err := usecase.repository.FindByQuery(query)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.DBResponse{}, err
		}
		return nil, err
	}

	return user, nil
}

func (usecase *userUseCase) FindByEmail(email string) (*models.DBResponse, error) {
	query := bson.M{"email": email}
	user, err := usecase.repository.FindByQuery(query)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.DBResponse{}, err
		}
		return nil, err
	}

	return user, nil
}

func (usecase *userUseCase) GetAll() ([]models.DBResponse, error) {
	return usecase.repository.GetAll(bson.M{})
}

func (usecase *userUseCase) UpdateStatus(id primitive.ObjectID, status int) error {
	//t := time.Now()

	query := bson.M{"_id": id}
	err := usecase.repository.UpdateStatus(query, status)
	return err
}

func (usecase *userUseCase) FilterUser(id primitive.ObjectID, c int64) ([]models.DBResponse, error) {
	return usecase.repository.FilterUser(id, c)
}

func (usecase *userUseCase) GetUser(id primitive.ObjectID) (models.DBResponse, error) {
	return usecase.repository.GetUser(id)
}
