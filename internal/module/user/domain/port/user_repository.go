package port

import (
	"github.com/TrueFlowDev/Backend/internal/module/user/domain/entity"
	user "github.com/TrueFlowDev/Backend/internal/module/user/domain/value_object"
	shared "github.com/TrueFlowDev/Backend/internal/shared/domain/value_object"
)

type UserRepository interface {
	Create(user *entity.User) error

	FindByID(id user.UserID) (*entity.User, error)
	FindByPhone(phone shared.Phone) (*entity.User, error)
	ExistsByPhone(phone shared.Phone) (bool, error)
}
