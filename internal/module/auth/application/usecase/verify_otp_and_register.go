package usecase

import (
	"context"
	"time"

	"github.com/TrueFlowDev/Backend/internal/module/auth/domain"
	"github.com/TrueFlowDev/Backend/internal/module/auth/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/user/application/usecase"
	user "github.com/TrueFlowDev/Backend/internal/module/user/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/value_object"
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
	registerUserUsecase usecase.RegisterUserUsecase
	accessTokenProvider port.AccessTokenProvider
	passwordHasher      port.PasswordHasher
	userIdGenerator     user.UserIdGenerator
}

func NewVerifyOTPAndRegisterUsecase(
	otpStore port.OTPStore,
	registerUserUsecase usecase.RegisterUserUsecase,
	accessTokenProvider port.AccessTokenProvider,
	passwordHasher port.PasswordHasher,
	userIdGenerator user.UserIdGenerator,
) *VerifyOTPAndRegisterUsecase {
	return &VerifyOTPAndRegisterUsecase{
		otpStore:            otpStore,
		registerUserUsecase: registerUserUsecase,
		accessTokenProvider: accessTokenProvider,
		passwordHasher:      passwordHasher,
		userIdGenerator:     userIdGenerator,
	}
}

func (u *VerifyOTPAndRegisterUsecase) Execute(ctx context.Context, input VerifyOTPAndRegisterInput) (VerifyOTPAndRegisterOutput, error) {
	userPhone, err := value_object.NewPhone(input.Phone)
	if err != nil {
		return VerifyOTPAndRegisterOutput{}, err
	}

	otp, err := u.otpStore.Get(ctx, userPhone)
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

	output, err := u.registerUserUsecase.Execute(ctx, usecase.RegisterUserInput{
		Phone:        input.Phone,
		HashPassword: newUserHashedPassword,
	})
	if err != nil {
		return VerifyOTPAndRegisterOutput{}, err
	}

	tokenClaims := domain.AccessTokenClaims{
		UserID:   value_object.NewUserID(output.ID),
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
