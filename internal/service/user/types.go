package user_service

import (
	"context"
	"go-jwt/internal/domain"
	"go-jwt/internal/pkg/hash"
	user_repository "go-jwt/internal/repository/user"

	"github.com/google/uuid"
)

type UserService struct {
	repo   user_repository.User
	hasher hash.PasswordHasher
}

type UserSignUpInput struct {
	Name     string
	Email    string
	Password string
}

type UserUpdateInput struct {
	ID    uuid.UUID
	Email string
	Name  string
}

type UserSignInInput struct {
	Email    string
	Password string
}

type UserSignInResponse struct {
	Token string
}

type Users interface {
	SignUp(ctx context.Context, input UserSignUpInput) (string, error)
	SignIn(ctx context.Context, input UserSignInInput) (string, error)
	FindAll() ([]domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindById(id string) (*domain.User, error)
	Update(ctx context.Context, input UserUpdateInput) error
	Delete(userId string) error
}
