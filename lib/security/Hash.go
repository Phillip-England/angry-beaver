package security

import "golang.org/x/crypto/bcrypt"

func Hash(value string) string {
	hashedValue, _ := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	return string(hashedValue)
}