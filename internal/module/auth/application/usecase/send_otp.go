package usecase

import (
	"context"
	"time"

	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/value_object"
)

type SendOtpInput struct {
	Phone string
}

type SendOtpUsecase struct {
	otpCodeGenerator port.OtpCodeGenerator
	otpStore         port.OTPStore
	otpSender        port.SmsOtpSender
}

func NewSendOtpUsecase(
	otpCodeGenerator port.OtpCodeGenerator,
	otpStore port.OTPStore,
	otpSender port.SmsOtpSender,
) *SendOtpUsecase {
	return &SendOtpUsecase{
		otpCodeGenerator: otpCodeGenerator,
		otpStore:         otpStore,
		otpSender:        otpSender,
	}
}

func (u *SendOtpUsecase) Execute(ctx context.Context, input SendOtpInput) error {
	phone, err := value_object.NewPhone(input.Phone)
	if err != nil {
		return err
	}

	otpCode := u.otpCodeGenerator.Generate()

	// TODO: this value must come from app configs
	duration := 5 * time.Minute
	otpExpiresAt := time.Now().UTC().Add(duration)
	otp, err := entity.NewOTP(otpCode, otpExpiresAt)
	if err != nil {
		return err
	}

	if err := u.otpStore.Set(ctx, phone, otp, duration); err != nil {
		return err
	}

	if err := u.otpSender.Send(ctx, phone, otp); err != nil {
		return err
	}

	return nil
}
