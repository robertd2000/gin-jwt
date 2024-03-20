package repository

import (
	"context"
	"go-jwt/internal/domain"

	"gorm.io/gorm"
)

type User interface {
	Create(c context.Context, student *domain.User) error
	FindAll() ([]domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindById(id string) (*domain.User, error)
}

type Repositories struct {
	User
}

func NewRepositories(db *gorm.DB) *Repositories {
	db.AutoMigrate(&domain.User{}, &domain.Object{})

	return &Repositories{
		NewUsersRepo(db),
	}
}
