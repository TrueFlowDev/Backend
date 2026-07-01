package value_object

import "strings"

type UserID struct {
	value string
}

func NewUserID(id string) UserID {
	id = strings.TrimSpace(id)
	return UserID{value: id}
}

func (u UserID) Value() string { return u.value }
