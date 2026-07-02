package value_object

import (
	"errors"
	"strings"

	"github.com/nyaruka/phonenumbers"
)

var (
	ErrPhoneRequired      = errors.New("phone is required")
	ErrPhoneInvalidFormat = errors.New("phone format is invalid")
)

type Phone struct {
	value string
}

func NewPhone(phone string) (Phone, error) {
	phone = strings.TrimSpace(phone)

	if phone == "" {
		return Phone{}, ErrPhoneRequired
	}

	number, err := phonenumbers.Parse(phone, "IR")
	if err != nil {
		return Phone{}, ErrPhoneInvalidFormat
	}

	if !phonenumbers.IsValidNumber(number) {
		return Phone{}, ErrPhoneInvalidFormat
	}

	if phonenumbers.GetNumberType(number) != phonenumbers.MOBILE {
		return Phone{}, ErrPhoneInvalidFormat
	}

	return Phone{
		value: phonenumbers.Format(number, phonenumbers.E164),
	}, nil
}

func (p Phone) Value() string { return p.value }
