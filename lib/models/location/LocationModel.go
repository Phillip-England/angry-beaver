package location

import (
	"fmt"

	"github.com/phillip-england/cfa-app/lib/validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LocationModel struct {
	User primitive.ObjectID
	Name string
	Number string
}

func NewLocationModel(userId primitive.ObjectID, name string, number string) LocationModel {
	return LocationModel{
		User: userId,
		Name: name,
		Number: number,
	}
}

func (v LocationModel) Validate() error {
	vName := validation.NewVString(v.Name)
	vNumber := validation.NewVString(v.Number)
	if vName.AssertRequired() != nil {
		return fmt.Errorf("name is required")
	}
	if vName.AssertMax(64) != nil {
		return fmt.Errorf("name too long")
	}
	if vName.AssertMin(5) != nil {
		return fmt.Errorf("name too short")
	}
	if vNumber.AssertNumber() != nil {
		return fmt.Errorf("number must be a valid number")
	}
	if vNumber.AssertMax(32) != nil {
		return fmt.Errorf("number too long")
	}
	if vNumber.AssertMin(3) != nil {
		return fmt.Errorf("number too short")
	}
	return nil
}

