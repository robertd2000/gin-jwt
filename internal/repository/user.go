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

func (repo *UsersRepo ) Create(_ context.Context, student *domain.User) error {
	res := repo.db.Create(&student)
	fmt.Print(student)
	if res.Error != nil {
		return errors.New("error occurs while creating new user")
	}
	return nil
}

func NewUsersRepo(db *gorm.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
		}
}
