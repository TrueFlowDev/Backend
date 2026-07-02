package port

import (
	"context"

	"github.com/TrueFlowDev/Backend/internal/module/user/domain/entity"
	user "github.com/TrueFlowDev/Backend/internal/module/user/domain/value_object"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error

	FindByID(ctx context.Context, id user.UserID) (*entity.User, error)
	FindByPhone(ctx context.Context, phone user.Phone) (*entity.User, error)
}
