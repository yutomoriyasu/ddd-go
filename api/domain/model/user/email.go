package user

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Email string

func NewEmail(email string) (Email, error) {
	if email == "" {
		return "", errors.New("email must not be empty")
	}
	if err := validation.Validate(email, is.Email); err != nil {
		return "", err
	}
	return Email(email), nil
}
