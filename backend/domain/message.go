package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id" binding:"required"`
	Content string             `json:"content,omitempty" binding:"required"`
}
