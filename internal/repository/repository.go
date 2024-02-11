package repository

import (
	"context"
	"go-jwt/internal/domain"
	"gorm.io/gorm"
)

type User interface {
	Create(c context.Context, 	student *domain.User) error

}

type Repositories struct {
	User
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
	NewUsersRepo(db),
	}
}