package service

import (
	"context"
	"go-jwt/internal/domain"
	"go-jwt/internal/pkg/hash"
	"go-jwt/internal/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/google/uuid"
)

type UserService struct {
	repo   repository.User
	hasher hash.PasswordHasher
}

func NewUserService(repo repository.User, hasher hash.PasswordHasher) *UserService {
	return &UserService{
		repo,
		hasher,
	}
}

func (s *UserService) SignUp(ctx context.Context, input UserSignUpInput) error {
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

func (s *UserService) SignIn(_ context.Context, input UserSignInInput) (string, error) {
	passwordHash, err := s.hasher.Hash(input.Password)
	if err != nil {
		return "", err
	}

	user, err := s.repo.FindByEmail(input.Email)
	if err != nil {
		return "", err
	}

	if user.Password != passwordHash {
		return "", err
	}

	payload := jwt.MapClaims{
		"sub": input.Email,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	secretKey := os.Getenv("SECRET_KEY") //
	var jwtSecretKey = []byte(secretKey)

	t, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *UserService) FindByEmail(email string) (*domain.User, error) {
	user, err := s.repo.FindByEmail(email)

	return user, err
}

func (s *UserService) FindById(id string) (*domain.User, error) {
	user, err := s.repo.FindById(id)

	return user, err
}

func (s *UserService) FindAll() ([]domain.User, error) {
	users, err := s.repo.FindAll()

	return users, err
}
