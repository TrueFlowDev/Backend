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
		Phone:          input.Phone,
		HashedPassword: newUserHashedPassword,
	})
	if err != nil {
		return VerifyOTPAndRegisterOutput{}, err
	}

	// TODO: this value must come from app configs
	duration := time.Hour
	now := time.Now().UTC()
	expiresAt := now.Add(duration)
	tokenClaims := value_object.NewAccessTokenClaims(output.ID, now, expiresAt)

	accessToken, err := u.accessTokenProvider.Generate(tokenClaims)
	if err != nil {
		return VerifyOTPAndRegisterOutput{}, err
	}

	return VerifyOTPAndRegisterOutput{AccessToken: accessToken}, nil
}
