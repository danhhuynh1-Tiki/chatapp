package repository

import (
	"chat/domain"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type userRepository struct {
	DB *mongo.Database
}

func NewRepository(DB *mongo.Database) domain.UserRepository {
	return &userRepository{DB}
}
func (u *userRepository) GetUser(user domain.Login) (*domain.User, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	userc := u.DB.Collection("users")

	var users domain.User

	err := userc.FindOne(ctx, bson.M{"email": user.Email, "password": user.Password}).Decode(&users)

	fmt.Println(user)
	fmt.Println(users)

	if err != nil {
		return nil, err
	} else {
		return &users, nil
	}

}
