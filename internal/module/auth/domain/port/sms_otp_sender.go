package port

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/value_object"
)

type SmsOtpSender interface {
	Send(ctx context.Context, phone value_object.Phone, otp entity.OTP) error
}
