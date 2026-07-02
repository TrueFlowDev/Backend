package phonenumber

import (
	"errors"
	"strings"

	"github.com/nyaruka/phonenumbers"
)

var (
	ErrRequired      = errors.New("phone is required")
	ErrInvalidFormat = errors.New("phone format is invalid")
)

func NormalizePhone(raw string) (string, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", ErrRequired
	}

	number, err := phonenumbers.Parse(raw, "IR")
	if err != nil {
		return "", ErrInvalidFormat
	}

	if !phonenumbers.IsValidNumber(number) {
		return "", ErrInvalidFormat
	}

	if phonenumbers.GetNumberType(number) != phonenumbers.MOBILE {
		return "", ErrInvalidFormat
	}

	return phonenumbers.Format(number, phonenumbers.E164), nil
}
