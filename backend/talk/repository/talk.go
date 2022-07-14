package repository

import (
	"chat/domain"
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type talkRepository struct {
	DB *mongo.Database
}

func NewTalkRepository(db *mongo.Database) domain.TalkRepository {
	return &talkRepository{db}
}

// if sigle talk not exists to create talk
// else handle talkusecase

// func (t *talkRepository) CreateTalk() {
// 	// create talk with people
// }

// id user1 && id user2
func (t *talkRepository) AddTalk(id string, id1 string) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Millisecond)

	talkc := t.DB.Collection("talk")

	var message []domain.Message

	res, err := talkc.InsertOne(ctx, bson.D{
		{Key: "user1", Value: id},
		{Key: "user2", Value: id1},
		{Key: "data", Value: message},
	})

	if err != nil {
		return errors.New("Cannot add talk")
	} else {
		fmt.Println(res.InsertedID)
		return nil
	}
}

func (t *talkRepository) GetTalk(id string, id1 string) (domain.Talk, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Millisecond)
	var talk domain.Talk
	talkc := t.DB.Collection("talk")

	err := talkc.FindOne(ctx, bson.D{
		{"user1", id},
		{"user2", id1},
	}).Decode(&talk)

	fmt.Println("repo get talk", talk.Messages)
	// fmt.Println(id,id1)
	if err == nil {
		return talk, errors.New("data exists")
	} else {
		return talk, nil
	}
}

func (t *talkRepository) AddMessage(talkId string, m domain.Message) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Millisecond)

	talkc := t.DB.Collection("talk")

	new_id, _ := primitive.ObjectIDFromHex(talkId)

	messages := t.GetMessage(new_id)
	messages = append(messages, m)

	res, err := talkc.UpdateOne(ctx, bson.M{
		"_id": new_id},
		bson.D{
			{"$set", bson.D{{"data", messages}}},
		},
	)
	if err != nil {
		return errors.New("Cannot update data")
	}
	fmt.Println(res.ModifiedCount)
	return nil
}

func (t *talkRepository) GetMessage(talkId primitive.ObjectID) []domain.Message {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Millisecond)
	var talk domain.Talk
	talkc := t.DB.Collection("talk")

	err := talkc.FindOne(ctx, bson.M{"_id": talkId}).Decode(&talk)

	fmt.Println("repo get message", talk.Messages)
	// fmt.Println(id,id1)
	if err == nil {
		return nil
	} else {
		return talk.Messages
	}
}
