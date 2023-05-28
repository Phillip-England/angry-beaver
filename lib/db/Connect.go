package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (client *mongo.Client, err error) {
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetMaxPoolSize(10))
	if err != nil {
		return nil, err
	}
	return client, nil
}