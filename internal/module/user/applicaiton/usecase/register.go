package usecase

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/user/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/port"
	user "github.com/TrueFlowDev/Backend/internal/module/user/domain/value_object"
	shared "github.com/TrueFlowDev/Backend/internal/shared/domain/value_object"
)

type RegisterUserInput struct {
	Phone        string
	HashPassword string
}

type RegisterUserOutput struct {
	ID string
}

type RegisterUserUsecase struct {
	userIdGenerator port.UserIdGenerator
	userRepository  port.UserRepository
}

func NewRegisterUserUsecase(
	userIdGenerator port.UserIdGenerator,
	userRepository port.UserRepository,
) *RegisterUserUsecase {
	return &RegisterUserUsecase{
		userIdGenerator: userIdGenerator,
		userRepository:  userRepository,
	}
}

func (u *RegisterUserUsecase) Execute(ctx context.Context, input RegisterUserInput) (RegisterUserOutput, error) {
	newUserPhone, err := shared.NewPhone(input.Phone)
	if err != nil {
		return RegisterUserOutput{}, err
	}

	newUserHashedPassword, err := user.NewHashedPassword(input.HashPassword)
	if err != nil {
		return RegisterUserOutput{}, err
	}

	newUserID := u.userIdGenerator.Generate()

	newUser, err := entity.NewUser(newUserID, newUserPhone, newUserHashedPassword)
	if err != nil {
		return RegisterUserOutput{}, err
	}

	if err := u.userRepository.Create(ctx, newUser); err != nil {
		return RegisterUserOutput{}, err
	}

	return RegisterUserOutput{ID: newUserID.Value()}, nil
}
