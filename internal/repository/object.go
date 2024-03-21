package repository

import (
	"context"
	"errors"
	"go-jwt/internal/domain"

	"gorm.io/gorm"
)

type ObjectRepo struct {
	db *gorm.DB
}

func NewObjectRepo(db *gorm.DB) *ObjectRepo {
	return &ObjectRepo{
		db: db,
	}
}

func (repo *ObjectRepo) Create(_ context.Context, object *domain.Object) error {
	res := repo.db.Create(&object)
	if res.Error != nil {
		return errors.New(res.Error.Error())
	}

	return nil

}

func (repo *ObjectRepo) FindAll() ([]domain.Object, error) {
	return []domain.Object{}, nil
}

func (repo *ObjectRepo) FindById(id string) (*domain.Object, error) {
	return &domain.Object{}, nil
}
