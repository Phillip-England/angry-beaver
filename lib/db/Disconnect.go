package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func Disconnect(client *mongo.Client) {
	err := client.Disconnect(context.Background())
	if err != nil {
		panic(err)
	}
}