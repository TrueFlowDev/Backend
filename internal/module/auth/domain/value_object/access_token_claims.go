package value_object

import (
	"time"
)

type AccessTokenClaims struct {
	userID    string
	issuedAt  time.Time
	expiresAt time.Time
}

func NewAccessTokenClaims(userID string, issuedAt time.Time, expiresAt time.Time) AccessTokenClaims {
	return AccessTokenClaims{
		userID:    userID,
		issuedAt:  issuedAt,
		expiresAt: expiresAt,
	}
}

func (a *AccessTokenClaims) UserID() string       { return a.userID }
func (a *AccessTokenClaims) IssuedAt() time.Time  { return a.issuedAt }
func (a *AccessTokenClaims) ExpiresAt() time.Time { return a.expiresAt }
