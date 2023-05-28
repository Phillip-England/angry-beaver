package session

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SessionDoc struct {
	ID primitive.ObjectID `bson:"_id"`
	User primitive.ObjectID `bson:"user"`
	Expires *time.Time `bson:"expires"`
}

func NewSessionDoc(sessionID primitive.ObjectID, model SessionModel) SessionDoc {
	return SessionDoc{
		ID: sessionID,
		User: model.UserID,
		Expires: nil,
	}
}