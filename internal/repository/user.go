package repository

import (
	"context"
	"errors"
	"fmt"
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

func (repo *UsersRepo) Create(_ context.Context, user *domain.User) error {
	var exists bool
	userDb := repo.db.Model(user).Where("email = ?", user.Email).Find(&exists)
	fmt.Print(userDb)
	if userDb.Error != nil {
		return errors.New("user already exists")
	}
	res := repo.db.Create(&user)
	if res.Error != nil {
		return errors.New(res.Error.Error())
	}
	return nil
}

func NewUsersRepo(db *gorm.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}
