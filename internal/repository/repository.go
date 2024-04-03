package repository

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

type Object interface {
	Create(c context.Context, object *domain.Object) error
	Update(_ context.Context, object *domain.Object) error
	FindAll() ([]domain.Object, error)
	FindById(id string) (*domain.Object, error)
	FindByUserId(userId string) ([]domain.Object, error)
	Delete(objectId string) error
	DeleteByUserId(userId string) error
}

type Repositories struct {
	User
	Object
}

func NewRepositories(db *gorm.DB) *Repositories {
	db.AutoMigrate(&domain.User{}, &domain.Object{})

	return &Repositories{
		NewUsersRepo(db),
		NewObjectRepo(db),
	}
}
