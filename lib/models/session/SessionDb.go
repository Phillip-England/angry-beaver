package session

import (
	"context"

	"github.com/phillip-england/cfa-app/lib/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionDb struct {
	coll *mongo.Collection
	model *SessionModel
}

func NewSessionDb(client *mongo.Client, sessionModel *SessionModel) SessionDb {
	db := SessionDb {
		coll: db.Collection(client, "sessions"),
		model: sessionModel,
	}
	return db
}

func (v SessionDb) FindById(token string) (*SessionDoc, error) {
	sessionId, _ := primitive.ObjectIDFromHex(token)
	filter := bson.D{{Key: "_id", Value: sessionId}}
	var sessionDoc SessionDoc
	err := v.coll.FindOne(context.Background(), filter).Decode(&sessionDoc)
	if err == mongo.ErrNoDocuments {
		return nil, err
	}
	return &sessionDoc, nil
	
}

func (v SessionDb) Insert() SessionDoc {
	result, _ := v.coll.InsertOne(context.Background(), v.model)
	sessionDoc := NewSessionDoc(result.InsertedID.(primitive.ObjectID), *v.model)
	return sessionDoc
}

func (v SessionDb) DropUserSessions() {
	v.coll.DeleteMany(context.Background(), bson.D{{Key: "user", Value: v.model.UserID}})
}
