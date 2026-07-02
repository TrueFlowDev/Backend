package usecase

import (
	"context"
	"time"

	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/value_object"
)

type VerifyOTPAndRegisterInput struct {
	Phone    string
	Password string
	Code     string
}

type VerifyOTPAndRegisterOutput struct {
	AccessToken string
}

type VerifyOTPAndRegisterUsecase struct {
	otpStore            port.OTPStore
	userRegisterer      port.UserRegisterer
	accessTokenProvider port.AccessTokenProvider
	passwordHasher      port.PasswordHasher
}

func NewVerifyOTPAndRegisterUsecase(
	otpStore port.OTPStore,
	userRegisterer port.UserRegisterer,
	accessTokenProvider port.AccessTokenProvider,
	passwordHasher port.PasswordHasher,
) *VerifyOTPAndRegisterUsecase {
	return &VerifyOTPAndRegisterUsecase{
		otpStore:            otpStore,
		userRegisterer:      userRegisterer,
		accessTokenProvider: accessTokenProvider,
		passwordHasher:      passwordHasher,
	}
}

func (u *VerifyOTPAndRegisterUsecase) Execute(ctx context.Context, input VerifyOTPAndRegisterInput) (VerifyOTPAndRegisterOutput, error) {
	phone, err := value_object.NewPhone(input.Phone)
	if err != nil {
		return VerifyOTPAndRegisterOutput{}, err
	}

	otp, err := u.otpStore.Get(ctx, phone)
	if err != nil {
		return VerifyOTPAndRegisterOutput{}, err
	}

	if err := otp.Verify(input.Code); err != nil {
		return VerifyOTPAndRegisterOutput{}, err
	}

	newUserHashedPassword, err := u.passwordHasher.Hash(input.Password)
	if err != nil {
		return VerifyOTPAndRegisterOutput{}, err
	}

	output, err := u.userRegisterer.Register(ctx, port.UserRegistererInput{
		Phone:        input.Phone,
		HashPassword: newUserHashedPassword,
	})
	if err != nil {
		return VerifyOTPAndRegisterOutput{}, err
	}

	tokenClaims := value_object.AccessTokenClaims{
		UserID:   output.ID,
		IssuedAt: time.Now().UTC(),
		// TODO: must come from configs
		ExpiresAt: time.Now().UTC().Add(time.Hour),
	}

	accessToken, err := u.accessTokenProvider.Generate(tokenClaims)
	if err != nil {
		return VerifyOTPAndRegisterOutput{}, err
	}

	return VerifyOTPAndRegisterOutput{AccessToken: accessToken}, nil
}
