package service

import (
	"go-jwt/internal/pkg/hash"
	"go-jwt/internal/repository"
	object_service "go-jwt/internal/service/object"
	user_service "go-jwt/internal/service/user"
)

type Services struct {
	Users   user_service.Users
	Objects object_service.Objects
}

type Deps struct {
	Repos  *repository.Repositories
	Hasher hash.PasswordHasher
}

func NewServices(deps Deps) *Services {
	userService := user_service.NewUserService(deps.Repos.User, deps.Hasher)
	objectService := object_service.NewObjectService(deps.Repos.Object, deps.Repos.User)

	return &Services{
		Users:   userService,
		Objects: objectService,
	}
}
