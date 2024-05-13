package user_repository

import (
	"context"
	"go-jwt/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User interface {
	Create(ctx context.Context, user *domain.User) (uuid.UUID, error)
	Update(_ context.Context, user *domain.User) error
	FindAll() ([]domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindByID(id string) (*domain.User, error)
	AddObject(user domain.User, object domain.Object) error
	Delete(userId string) error
}

type UsersRepo struct {
	db *gorm.DB
}
