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
func (u *userRepository) GetAllUser() []domain.User {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// defer DB.Disconnect(ctx)

	// example := DB.Database("example")

	userc := u.DB.Collection("users")

	cursor, err := userc.Find(ctx, bson.M{})

	defer cursor.Close(ctx)

	alluser := []domain.User{}
	for cursor.Next(ctx) {
		var user domain.User
		cursor.Decode(&user)

		alluser = append(alluser, user)
	}
	if err != nil {
		return nil
	} else {
		return alluser
	}
}
func (u *userRepository) GetUser(user domain.User) (*domain.User, error) {

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
func (u *userRepository) Create(user domain.User) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	userc := u.DB.Collection("users")
	_, err := userc.InsertOne(ctx, bson.D{
		{Key: "email", Value: user.Email},
		{Key: "password", Value: user.Password},
		{Key: "name", Value: user.Name},
		{Key: "address", Value: user.Address},
		{Key: "phone", Value: user.Phone},
	})
	if err != nil {
		return nil, err
	}

	return &user.Email, nil
}

// func (u *userRepository) IsExisted(email string) (*domain.User, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

// 	defer cancel()
// 	userc := u.DB.Collection("users")
// 	var user domain.User
// 	filter := bson.D{{"email", email}}
// 	opts := options.FindOne()
// 	err := userc.FindOne(ctx, filter, opts).Decode(&user)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }
