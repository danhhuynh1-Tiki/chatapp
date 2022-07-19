package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StatusUsecase interface {
	UpdateStatusUser(id primitive.ObjectID, t time.Time, status int) error
}
