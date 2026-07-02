package main

import (
	"github.com/TrueFlowDev/Backend/internal/module/auth"
	"github.com/TrueFlowDev/Backend/internal/module/notification"
	"github.com/TrueFlowDev/Backend/internal/module/user"
	"github.com/TrueFlowDev/Backend/internal/shared"
	"go.uber.org/fx"
)

func main() {
	fx.
		New(
			shared.Module,
			user.Module,
			auth.Module,
			notification.Module,
		).
		Run()
}
