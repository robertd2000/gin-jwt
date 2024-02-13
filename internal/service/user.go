package service

import (
	"context"
	"go-jwt/internal/domain"
	"go-jwt/internal/pkg/hash"
	"go-jwt/internal/repository"

	"github.com/google/uuid"
)

type StudentsService struct {
	repo   repository.User
	hasher hash.PasswordHasher
}

func (s *StudentsService) SignUp(ctx context.Context, input UserSignUpInput) error {
	passwordHash, err := s.hasher.Hash(input.Password)
	if err != nil {
		return err
	}

	student := domain.User{
		Name:     input.Name,
		Password: passwordHash,
		Email:    input.Email,
		ID:       uuid.New(),
	}

	if err := s.repo.Create(ctx, &student); err != nil {
		return err
	}

	return nil
}

func NewUserService(repo repository.User, hasher hash.PasswordHasher) *StudentsService {
	return &StudentsService{
		repo,
		hasher,
	}
}
