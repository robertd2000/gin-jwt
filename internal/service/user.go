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

func (s *StudentsService) FindByEmail(email string) (*domain.User, error) {
	user, err := s.repo.FindByEmail(email)

	return user, err
}

func (s *StudentsService) FindById(id string) (*domain.User, error) {
	user, err := s.repo.FindById(id)

	return user, err
}

func (s *StudentsService) FindAll() ([]domain.User, error) {
	users, err := s.repo.FindAll()

	return users, err
}

func NewUserService(repo repository.User, hasher hash.PasswordHasher) *StudentsService {
	return &StudentsService{
		repo,
		hasher,
	}
}
