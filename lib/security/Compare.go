package security

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Compare(hashedValue string, value string) error {
	hashedValueBytes := []byte(hashedValue)
	valueBytes := []byte(value)
	err := bcrypt.CompareHashAndPassword(hashedValueBytes, valueBytes)
	if err != nil {
		return fmt.Errorf("bcrypt compare failure")
	}
	return nil
}