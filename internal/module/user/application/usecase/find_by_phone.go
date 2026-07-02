package usecase

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/user/domain/entity"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/port"
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/value_object"
)

type FindUserByPhoneInput struct {
	Phone string
}

type FindUserByPhoneUsecase struct {
	userRepository port.UserRepository
}

func NewFindUserByPhoneUsecase(
	userRepository port.UserRepository,
) *FindUserByPhoneUsecase {
	return &FindUserByPhoneUsecase{
		userRepository: userRepository,
	}
}

func (u *FindUserByPhoneUsecase) Execute(ctx context.Context, input FindUserByPhoneInput) (*entity.User, error) {
	userPhone, err := value_object.NewPhone(input.Phone)
	if err != nil {
		return nil, err
	}

	user, err := u.userRepository.FindByPhone(ctx, userPhone)
	if err != nil {
		return nil, err
	}

	return user, nil
}
