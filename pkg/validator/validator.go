package validator

import (
	"errors"
	"regexp"
)

type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) ValidateEmail(email string) error {
	if email == "" {
		return errors.New("email cannot be empty")
	}

	// Simple email validation
	pattern := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(pattern, email)
	
	if !match {
		return errors.New("invalid email format")
	}

	return nil
}