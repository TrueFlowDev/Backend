package port

import "github.com/TrueFlowDev/Backend/internal/module/auth/domain"

type AccessTokenProvider interface {
	Generate(claims domain.AccessTokenClaims) (string, error)
	Verify(token string) (domain.AccessTokenClaims, error)
}
