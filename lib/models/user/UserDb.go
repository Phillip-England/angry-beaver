package user

import (
	"context"
	"fmt"

	"github.com/phillip-england/cfa-app/lib/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDb struct {
	coll *mongo.Collection
	model *UserModel
}

func NewUserDb(client *mongo.Client, model *UserModel) UserDb {
	return UserDb {
		coll: db.Collection(client, "users"),
		model: model,
	}
}

func (v UserDb) AssertUnique() error {
	var userExists UserModel
	filter := bson.D{{Key: "email", Value: v.model.Email}}
	err := v.coll.FindOne(context.Background(), filter).Decode(&userExists)
	if userExists.Email == v.model.Email && err != mongo.ErrNoDocuments {
		return fmt.Errorf("user already exists")
	}
	return nil
}

func (v UserDb) Insert() UserDoc {
	result, _ := v.coll.InsertOne(context.Background(), v.model)
	userDoc := NewUserDoc(result.InsertedID.(primitive.ObjectID), *v.model)
	return userDoc
}

func (v UserDb) Pull() (*UserDoc, error) {
	var userDoc UserDoc
	filter := bson.D{{Key: "email", Value: v.model.Email}}
	err := v.coll.FindOne(context.Background(), filter).Decode(&userDoc)
	if err == mongo.ErrNoDocuments {
		return nil, err
	}
	return &userDoc, nil
}

func (v UserDb) FindById(userId primitive.ObjectID) (*UserDoc, error) {
	var userDoc UserDoc
	filter := bson.D{{Key: "_id", Value: userId}}
	err := v.coll.FindOne(context.Background(), filter).Decode(&userDoc)
	if err == mongo.ErrNoDocuments {
		return nil, err
	}
	return &userDoc, nil
}