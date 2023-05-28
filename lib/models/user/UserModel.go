package user

import (
	"fmt"
	"strings"

	security "github.com/phillip-england/cfa-app/lib/security"
	"github.com/phillip-england/cfa-app/lib/validation"
)

type UserModel struct {
	Email string `bson:"email"`
	Password string `bson:"password"`
}

func NewUserModel(email string, password string) UserModel {
	model := UserModel {
		Email: strings.ToLower(email),
		Password: password,
	}
	return model
}

func (v UserModel) Validate() error {
	vEmail := validation.NewVString(v.Email)
	vPassword := validation.NewVString(v.Password)
	if vEmail.AssertEmail() != nil {
		return fmt.Errorf("invalid email")
	}
	if vPassword.AssertRequired() != nil {
		return fmt.Errorf("password is required")
	}
	if vPassword.AssertMax(64) != nil {
		return fmt.Errorf("password too long")
	}
	if vPassword.AssertMin(5) != nil {
		return fmt.Errorf("password too short")
	}
	return nil
}

func (v *UserModel) HashPassword() {
	v.Password = security.Hash(v.Password)
}
