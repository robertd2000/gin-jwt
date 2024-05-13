package service

import (
	"context"
	"go-jwt/internal/domain"
	"go-jwt/internal/pkg/hash"
	"go-jwt/internal/repository"
	user_service "go-jwt/internal/service/user"

	"github.com/google/uuid"
)

type ObjectCreateInput struct {
	Name        string
	Type        int
	Coords      string
	Radius      int
	Description string
	Color       string
	UserID      uuid.UUID
}

type ObjectUpdateInput struct {
	ID          uuid.UUID
	Name        string
	Type        int
	Coords      string
	Radius      int
	Description string
	Color       string
	UserID      uuid.UUID
}

type Objects interface {
	Create(ctx context.Context, objectInput ObjectCreateInput) error
	Update(ctx context.Context, objectInput ObjectUpdateInput) error
	FindAll() ([]domain.Object, error)
	FindById(id string) (*domain.Object, error)
	FindByUserId(userId string) ([]domain.Object, error)
	Delete(objectId string) error
}

type Services struct {
	Users   user_service.Users
	Objects Objects
}

type Deps struct {
	Repos  *repository.Repositories
	Hasher hash.PasswordHasher
}

func NewServices(deps Deps) *Services {
	userService := user_service.NewUserService(deps.Repos.User, deps.Hasher)
	objectService := NewObjectService(deps.Repos.Object, deps.Repos.User)
	return &Services{
		Users:   userService,
		Objects: objectService,
	}
}
