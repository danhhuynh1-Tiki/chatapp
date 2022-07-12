package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	model "jwt-demo/models"
)

var collection *mongo.Collection

type UserMongo struct {
	collection *mongo.Collection
	database *mongo.Client
}

var NewUserMongo = func(db *mongo.Client) *UserMongo {
	collection = GetCollection(db, model.DB_NAME)
	return &UserMongo{
		collection: GetCollection(db, model.DB_NAME),
		database: db,
	}
}

/**
 * Used to insert a new user into the database
 */
func (userMongo *UserMongo) Create(user model.User) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	
	defer cancel()

	_, err := collection.InsertOne(ctx, &user)
	if err != nil {
		return nil, err
	}

	return &user.Email, nil
}

func (userMongo *UserMongo) IsExisted(email string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	
	defer cancel()

	var user model.User
	filter := bson.D{{ "email", email }}
	opts := options.FindOne()
	err := collection.FindOne(ctx, filter, opts).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}