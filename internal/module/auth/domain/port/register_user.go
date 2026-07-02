package port

import "context"

type UserRegistererInput struct {
	Phone          string
	HashedPassword string
}

type UserRegistererOutput struct {
	ID string
}

type UserRegisterer interface {
	Register(ctx context.Context, input UserRegistererInput) (UserRegistererOutput, error)
}
