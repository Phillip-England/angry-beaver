package validation

import (
	"fmt"
	"net/mail"
	"strconv"
)

type VString struct {
	value string
}

func NewVString(value string) VString {
	return VString {
		value: value,
	}
}

func (v VString) AssertEmail() error {
	_, err := mail.ParseAddress(v.value)
	if err != nil {
		return fmt.Errorf("not a valid email")
	}
	return nil
}

func (v VString) AssertMax(max int) error {
	if len(v.value) > max {
		return fmt.Errorf("value too long")
	}
	return nil
}

func (v VString) AssertMin(min int) error {
	if len(v.value) < min {
		return fmt.Errorf("value too short")
	}
	return nil
}

func (v VString) AssertRequired() error {
	if len(v.value) == 0 {
		return fmt.Errorf("value is required")
	}
	return nil
}

func (v VString) AssertNumber() error {
	_, err := strconv.ParseInt(v.value, 10, 64)
	if err != nil {
		return fmt.Errorf("value must be a number")
	}
	return nil
}