package repository

import (
	"context"
	"errors"
	"go-jwt/internal/domain"

	"gorm.io/gorm"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UsersRepo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

func (repo *UsersRepo) Create(_ context.Context, user *domain.User) error {
	var exists bool

	userDb := repo.db.Model(user).Where("email = ?", user.Email).Find(&exists)

	if userDb.Error != nil {
		return errors.New("user already exists")
	}
	res := repo.db.Create(&user)
	if res.Error != nil {
		return errors.New(res.Error.Error())
	}
	return nil
}

func (repo *UsersRepo) Update(_ context.Context, user *domain.User) error {
	res := repo.db.Save(&user)

	if res.Error != nil {
		return errors.New(res.Error.Error())
	}

	return nil
}

func (repo *UsersRepo) FindByEmail(email string) (*domain.User, error) {
	var user domain.User

	err := repo.db.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &user, nil
}

func (repo *UsersRepo) FindById(id string) (*domain.User, error) {
	var user domain.User

	err := repo.db.Model(&domain.User{}).Preload("Objects").Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &user, nil
}

func (repo *UsersRepo) FindAll() ([]domain.User, error) {
	var users []domain.User

	err := repo.db.Model(&domain.User{}).Preload("Objects").Find(&users).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return users, nil
}

func (repo *UsersRepo) AddObject(user domain.User, object domain.Object) error {
	repo.db.Model(&user).Association("Languages").Append(object)

	return nil
}

func (repo *UsersRepo) Delete(userId string) error {
	res := repo.db.Where("id = ?", userId).Delete(&domain.User{})
	if res.Error != nil {
		return errors.New(res.Error.Error())
	}

	res = repo.db.Where("user_id = ?", userId).Delete(&domain.Object{})
	if res.Error != nil {
		return errors.New(res.Error.Error())
	}

	return nil
}
