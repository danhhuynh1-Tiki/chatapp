package usecase

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"chat/pkg/utils"
	"chat/services/models"
	"chat/services/repository"
)

type AuthUseCase interface {
	SignUp(*models.SignUpInput) (*models.DBResponse, error)
}

type authUseCase struct {
	repository repository.AuthRepository
}

func NewAuthUseCase(repository repository.AuthRepository) AuthUseCase {
	return &authUseCase{
		repository: repository,
	}
}

func (usecase *authUseCase) SignUp(userInput *models.SignUpInput) (*models.DBResponse, error) {
	userInput.CreatedAt = time.Now()
	userInput.UpdatedAt = userInput.CreatedAt
	userInput.RequestAt = userInput.CreatedAt
	//userInput.PasswordConfirm = ""
	userInput.Status = 1
	hashedPassword, _ := utils.HashPassword(userInput.Password)
	userInput.Password = hashedPassword

	res, err := usecase.repository.AddUser(userInput)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("user with that email already exist")
		}
		return nil, err
	}

	// create a unique index for the email field
	if _, err := usecase.repository.IndexesUser(); err != nil {
		return nil, err
	}

	query := bson.M{"_id": res.InsertedID}
	newUser, err := usecase.repository.GetUser(query)

	if err != nil {
		return nil, err
	}

	return newUser, nil
}
