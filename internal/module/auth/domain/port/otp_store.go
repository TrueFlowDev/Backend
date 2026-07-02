package port

import (
	"context"
	"time"

	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/value_object"
)

type OTPStore interface {
	Set(ctx context.Context, key value_object.Phone, value entity.OTP, ttl time.Duration) error
	Get(ctx context.Context, key value_object.Phone) (entity.OTP, error)
	Delete(ctx context.Context, key value_object.Phone) error
}
