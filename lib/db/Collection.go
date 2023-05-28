package db

import (
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func Collection(client *mongo.Client, name string) (coll *mongo.Collection) {
	coll = client.Database(os.Getenv("MONGO_DB")).Collection(name)
	return coll
}