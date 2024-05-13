package object_service

import (
	"context"
	"go-jwt/internal/domain"
	object_repository "go-jwt/internal/repository/object"
	user_repository "go-jwt/internal/repository/user"

	"github.com/google/uuid"
)

type ObjectService struct {
	objectRepo object_repository.Object
	userRepo   user_repository.User
}

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
