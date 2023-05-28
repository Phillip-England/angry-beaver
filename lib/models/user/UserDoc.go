package user

import "go.mongodb.org/mongo-driver/bson/primitive"


type UserDoc struct {
	ID primitive.ObjectID `bson:"_id"`
	Email string `bson:"email"`
	Password string `bson:"password"`
}

func NewUserDoc(id primitive.ObjectID, userModel UserModel) UserDoc {
	return UserDoc{
		ID: id,
		Email: userModel.Email,
		Password: userModel.Password,
	}
}