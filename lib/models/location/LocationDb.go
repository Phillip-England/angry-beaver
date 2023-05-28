package location

import (
	"context"
	"fmt"

	"github.com/phillip-england/cfa-app/lib/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LocationDb struct {
	coll *mongo.Collection
	model *LocationModel
}

func NewLocationDb(client *mongo.Client, model *LocationModel) LocationDb {
	return LocationDb {
		coll: db.Collection(client, "locations"),
		model: model,
	}
}

func (v LocationDb) Insert() LocationDoc {
	result, _ := v.coll.InsertOne(context.Background(), v.model)
	doc := NewLocationDoc(result.InsertedID.(primitive.ObjectID), *v.model)
	return doc	
}

func (v LocationDb) AssertUnique() (error) {
	var locationExists LocationModel
	filter := bson.D{
		{
			Key: "$or",
			Value: bson.A{
				bson.D{{Key: "name", Value: v.model.Name}},
				bson.D{{Key: "number", Value: v.model.Number}},
			},
		},
	}
	err := v.coll.FindOne(context.Background(), filter).Decode(&locationExists)
	
	if err != mongo.ErrNoDocuments {
		return fmt.Errorf("location already exists")
	}
	
	return nil	
}