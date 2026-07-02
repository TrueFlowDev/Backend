package value_object

import (
	"time"
)

type AccessTokenClaims struct {
	UserID    string
	IssuedAt  time.Time
	ExpiresAt time.Time
}
