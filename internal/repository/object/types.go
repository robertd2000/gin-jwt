package object_repository

import (
	"context"
	"go-jwt/internal/domain"

	"gorm.io/gorm"
)

type ObjectRepo struct {
	db *gorm.DB
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
