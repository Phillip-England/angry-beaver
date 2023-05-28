package session

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SessionModel struct {
	UserID primitive.ObjectID `bson:"user"`
	Expires time.Time
}

func NewSessionModel(userID primitive.ObjectID) SessionModel {
	model := SessionModel {
		UserID: userID,
		Expires: time.Now().Add(24 * time.Hour),
	}
	return model
}