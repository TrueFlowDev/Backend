package value_object

import (
	"github.com/TrueFlowDev/Backend/internal/pkg/phonenumber"
)

type Phone struct {
	value string
}

func NewPhone(phone string) (Phone, error) {
	normalized, err := phonenumber.NormalizePhone(phone)
	if err != nil {
		return Phone{}, err
	}

	return Phone{
		value: normalized,
	}, nil
}

func (p Phone) Value() string { return p.value }
