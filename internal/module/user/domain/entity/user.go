package entity

import (
	"time"

	user "github.com/TrueFlowDev/Backend/internal/module/user/domain/value_object"
	shared "github.com/TrueFlowDev/Backend/internal/shared/domain/value_object"
)

type User struct {
	id       user.UserID
	phone    shared.Phone
	password user.HashedPassword

	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

func NewUser(
	id user.UserID,
	phone shared.Phone,
	password user.HashedPassword,
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

func (u *User) ID() user.UserID               { return u.id }
func (u *User) Phone() shared.Phone           { return u.phone }
func (u *User) Password() user.HashedPassword { return u.password }
func (u *User) CreatedAt() time.Time          { return u.createdAt }
func (u *User) UpdatedAt() time.Time          { return u.updatedAt }
func (u *User) DeletedAt() *time.Time         { return u.deletedAt }

// <-- Setters -->

func (u *User) ChangePassword(newPassword user.HashedPassword) {
	u.password = newPassword
	u.updatedAt = time.Now().UTC()
}
