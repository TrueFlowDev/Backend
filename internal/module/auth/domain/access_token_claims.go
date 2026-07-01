package domain

import (
	"time"

	user "github.com/TrueFlowDev/Backend/internal/module/user/domain/value_object"
)

type AccessTokenClaims struct {
	UserID    user.UserID
	IssuedAt  time.Time
	ExpiresAt time.Time
}
