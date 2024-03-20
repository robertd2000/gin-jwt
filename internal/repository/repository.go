package repository

import (
	"context"
	"go-jwt/internal/domain"

	"gorm.io/gorm"
)

type User interface {
	Create(c context.Context, user *domain.User) error
	FindAll() ([]domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindById(id string) (*domain.User, error)
}

type Object interface {
	Create(c context.Context, object *domain.Object) error
	FindAll() ([]domain.Object, error)
	FindById(id string) (*domain.Object, error)
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
