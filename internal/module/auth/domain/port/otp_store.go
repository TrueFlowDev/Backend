package port

import (
	"context"
	"time"

	"github.com/TrueFlowDev/Backend/internal/module/auth/domain"
	shared "github.com/TrueFlowDev/Backend/internal/shared/domain/value_object"
)

type OTPStore interface {
	Set(ctx context.Context, key shared.Phone, value domain.OTP, ttl time.Duration) error
	Get(ctx context.Context, key shared.Phone) (domain.OTP, error)
	Delete(ctx context.Context, key shared.Phone) error
}
