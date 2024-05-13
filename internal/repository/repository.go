package repository

import (
	"context"
	"go-jwt/internal/domain"
	user_repository "go-jwt/internal/repository/user"

	"gorm.io/gorm"
)

type Object interface {
	Create(c context.Context, object *domain.Object) error
	Update(_ context.Context, object *domain.Object) error
	FindAll() ([]domain.Object, error)
	FindById(id string) (*domain.Object, error)
	FindByUserId(userId string) ([]domain.Object, error)
	Delete(objectId string) error
	DeleteByUserId(userId string) error
}

type Repositories struct {
	User user_repository.User
	Object
}

func NewRepositories(db *gorm.DB) *Repositories {
	db.AutoMigrate(&domain.User{}, &domain.Object{})

	return &Repositories{
		User:   user_repository.NewUsersRepo(db),
		Object: NewObjectRepo(db),
		//  user.NewUsersRepo(db),
		// 	NewObjectRepo(db),
	}
}
