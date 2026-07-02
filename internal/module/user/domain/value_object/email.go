package value_object

import (
	"errors"
	"net/mail"
	"strings"
)

var (
	ErrEmailRequired      = errors.New("email is required")
	ErrEmailInvalidFormat = errors.New("email format is invalid")
)

type Email struct {
	value string
}

func NewEmail(email string) (Email, error) {
	email = strings.TrimSpace(email)

	if email == "" {
		return Email{}, ErrEmailRequired
	}

	email = strings.ToLower(email)

	addr, err := mail.ParseAddress(email)
	if err != nil || addr.Address != email {
		return Email{}, ErrEmailInvalidFormat
	}

	return Email{value: email}, nil
}

func (e Email) Value() string { return e.value }
