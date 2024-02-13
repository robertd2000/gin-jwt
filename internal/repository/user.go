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

	err := repo.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &user, nil
}

func (repo *UsersRepo) FindAll() ([]domain.User, error) {
	var users []domain.User

	err := repo.db.Find(&users).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return users, nil
}
