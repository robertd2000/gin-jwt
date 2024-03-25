package service

import (
	"context"
	"go-jwt/internal/domain"
	"go-jwt/internal/pkg/hash"
	"go-jwt/internal/repository"

	"github.com/google/uuid"
)

type UserSignUpInput struct {
	Name     string
	Email    string
	Password string
}

type UserSignInInput struct {
	Email    string
	Password string
}

type UserSignInResponse struct {
	Token string
}

type ObjectCreateInput struct {
	Name        string
	Type        int
	Coords      string
	Radius      int
	Description string
	Color       string
	UserID      uuid.UUID
}

type ObjectUpdateInput struct {
	ID          uuid.UUID
	Name        string
	Type        int
	Coords      string
	Radius      int
	Description string
	Color       string
	UserID      uuid.UUID
}

type Users interface {
	SignUp(ctx context.Context, input UserSignUpInput) error
	SignIn(ctx context.Context, input UserSignInInput) (string, error)
	FindAll() ([]domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindById(id string) (*domain.User, error)
	Delete(userId string) error
}

type Objects interface {
	Create(ctx context.Context, objectInput ObjectCreateInput) error
	Update(ctx context.Context, objectInput ObjectUpdateInput) error
	FindAll() ([]domain.Object, error)
	FindById(id string) (*domain.Object, error)
	FindByUserId(userId string) ([]domain.Object, error)
	Delete(objectId string) error
}

type Services struct {
	Users   Users
	Objects Objects
}

type Deps struct {
	Repos  *repository.Repositories
	Hasher hash.PasswordHasher
}

func NewServices(deps Deps) *Services {
	userService := NewUserService(deps.Repos.User, deps.Hasher)
	objectService := NewObjectService(deps.Repos.Object, deps.Repos.User)
	return &Services{
		Users:   userService,
		Objects: objectService,
	}
}
