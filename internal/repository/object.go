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

func (repo *ObjectRepo) Update(_ context.Context, object *domain.Object) error {
	res := repo.db.Save(&object)
	if res.Error != nil {
		return errors.New(res.Error.Error())
	}

	return nil

}

func (repo *ObjectRepo) FindAll() ([]domain.Object, error) {
	var objects []domain.Object

	err := repo.db.Model(&domain.Object{}).Preload("User").Find(&objects).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return objects, nil
}

func (repo *ObjectRepo) FindById(id string) (*domain.Object, error) {
	var object domain.Object

	err := repo.db.Where("id = ?", id).First(&object).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &object, nil
}

func (repo *ObjectRepo) FindByUserId(userId string) ([]domain.Object, error) {
	var objects []domain.Object

	err := repo.db.Where("user_id = ?", userId).Find(&objects).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return objects, nil
}

func (repo *ObjectRepo) Delete(objectId string) error {
	res := repo.db.Where("id = ?", objectId).Delete(&domain.Object{})
	if res.Error != nil {
		return errors.New(res.Error.Error())
	}

	return nil
}

func (repo *ObjectRepo) DeleteByUserId(userId string) error {
	res := repo.db.Where("user_id = ?", userId).Delete(&domain.Object{})
	if res.Error != nil {
		return errors.New(res.Error.Error())
	}

	return nil
}
