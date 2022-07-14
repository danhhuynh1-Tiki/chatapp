package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Talk struct {
	ID       string    `json:"id,omitempty" bson:"_id"`
	User1    string    `json:"user1" bson:"user1"`
	User2    string    `json:"user2" bson:"user2"`
	Messages []Message `json:"data" bson:"data"`
}
type TalkGroup struct {
}

type TalkRepository interface {
	AddTalk(id string, id1 string) error
	GetTalk(id string, id1 string) (Talk, error)
	AddMessage(id string, m Message) error
	GetMessage(id primitive.ObjectID) []Message
}

type TalkUsecase interface {
	AddTalk(id string, id1 string) error
	GetTalk(id string, id1 string) (Talk, error)
	GetMessage(id string) []Message
	AddMessage(id string, m Message) error
}
