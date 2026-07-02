package port

import (
	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/value_object"
)

type AccessTokenProvider interface {
	Generate(claims value_object.AccessTokenClaims) (string, error)
	Verify(token string) (value_object.AccessTokenClaims, error)
}
