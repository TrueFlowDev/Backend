package port

import (
	"github.com/TrueFlowDev/Backend/internal/module/auth/domain"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/value_object"
)

type SmsOtpSender interface {
	Send(phone value_object.Phone, otp domain.OTP) error
}
