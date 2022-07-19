package repository

import (
	"chat/domain"
	"errors"
	"fmt"

	// "fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	// fmt.Println(user)
	// fmt.Println(users)

	if err != nil {
		return nil, err
	} else {
		return &users, nil
	}

}
func (u *userRepository) Create(user domain.User) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	user.Created_At = time.Now()
	user.Updated_At = time.Now()
	user.Request_At = time.Now()
	user.Status = 1

	defer cancel()
	userc := u.DB.Collection("users")
	_, err := userc.InsertOne(ctx, bson.D{
		{Key: "email", Value: user.Email},
		{Key: "password", Value: user.Password},
		{Key: "name", Value: user.Name},
		{Key: "address", Value: user.Address},
		{Key: "phone", Value: user.Phone},
		{Key: "created_at", Value: user.Created_At},
		{Key: "updated_at", Value: user.Updated_At},
		{Key: "request_at", Value: user.Request_At},
		{Key: "status", Value: user.Status},
	})
	if err != nil {
		return nil, err
	}

	return &user.Email, nil
}

func (u *userRepository) UpdateStatusUser(id primitive.ObjectID, t time.Time, status int) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	userc := u.DB.Collection("users")
	res, err := userc.UpdateOne(ctx,
		bson.M{"_id": id}, bson.D{
			{"$set", bson.D{
				{"status", status},
				{"request_at", t},
			}},
		})
	fmt.Println(res.ModifiedCount)
	if err != nil || res.ModifiedCount == 0 {
		return errors.New("Cannot update status user")
	}
	return nil
}
