package user

import (
	"github.com/TrueFlowDev/Backend/internal/module/user/application/usecase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"user",
	fx.Provide(
		usecase.NewRegisterUserUsecase,
	),
)
