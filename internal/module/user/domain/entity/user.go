package entity

import (
	"time"

	"github.com/TrueFlowDev/Backend/internal/module/user/domain/value_object"
)

type User struct {
	id       value_object.UserID
	phone    value_object.Phone
	password value_object.HashedPassword

	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

func NewUser(
	id value_object.UserID,
	phone value_object.Phone,
	password value_object.HashedPassword,
) (*User, error) {
	now := time.Now().UTC()
	return &User{
		id:        id,
		phone:     phone,
		password:  password,
		createdAt: now,
		updatedAt: now,
		deletedAt: nil,
	}, nil
}

// <-- Getters -->

func (u *User) ID() value_object.UserID               { return u.id }
func (u *User) Phone() value_object.Phone             { return u.phone }
func (u *User) Password() value_object.HashedPassword { return u.password }
func (u *User) CreatedAt() time.Time                  { return u.createdAt }
func (u *User) UpdatedAt() time.Time                  { return u.updatedAt }
func (u *User) DeletedAt() *time.Time                 { return u.deletedAt }

// <-- Setters -->

func (u *User) ChangePassword(newPassword value_object.HashedPassword) {
	u.password = newPassword
	u.updatedAt = time.Now().UTC()
}
