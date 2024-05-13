package repository

import (
	"go-jwt/internal/domain"
	object_repository "go-jwt/internal/repository/object"
	user_repository "go-jwt/internal/repository/user"

	"gorm.io/gorm"
)

type Repositories struct {
	User   user_repository.User
	Object object_repository.Object
}

func NewRepositories(db *gorm.DB) *Repositories {
	db.AutoMigrate(&domain.User{}, &domain.Object{})

	return &Repositories{
		User:   user_repository.NewUsersRepo(db),
		Object: object_repository.NewObjectRepo(db),
	}
}
