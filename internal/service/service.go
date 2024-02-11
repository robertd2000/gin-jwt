package service

import (
	"context"
	"go-jwt/internal/pkg/hash"
	"go-jwt/internal/repository"
)

type UserSignUpInput struct {
	Name         string
	Email        string
	Password     string
	
}


type Users interface {
	SignUp(ctx context.Context, input UserSignUpInput) error
//	SignIn(ctx context.Context, input UserSignInInput) (Tokens, error)
}

type Services struct {
	Users          Users
}

type Deps struct {
	Repos                  *repository.Repositories
	Hasher                 hash.PasswordHasher
}

func NewServices(deps Deps) *Services {
	userService := NewUserService(deps.Repos.User, deps.Hasher)
	return &Services{
		Users: userService,
	}
}