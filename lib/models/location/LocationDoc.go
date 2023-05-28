package location

import "go.mongodb.org/mongo-driver/bson/primitive"


type LocationDoc struct {
	ID primitive.ObjectID `bson:"_id"`
	Name string `bson:"name"`
	Number string `bson:"number"`
}

func NewLocationDoc(id primitive.ObjectID, locationModel LocationModel) LocationDoc {
	return LocationDoc{
		ID: id,
		Name: locationModel.Name,
		Number: locationModel.Number,
	}
}