package user_service

import (
	"context"
	"go-jwt/internal/domain"
	"go-jwt/internal/pkg/hash"
	user_repository "go-jwt/internal/repository/user"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserService struct {
	repo   user_repository.User
	hasher hash.PasswordHasher
}

func NewUserService(repo user_repository.User, hasher hash.PasswordHasher) *UserService {
	return &UserService{
		repo,
		hasher,
	}
}

// SignUp is a function that handles user sign up.
// It takes in a context.Context object, UserSignUpInput, and returns an error.
// This function hashes the password, creates a new user domain object, and saves it to the repository.
func (s *UserService) SignUp(ctx context.Context, input UserSignUpInput) (string, error) {
	// Hash the password
	passwordHash, err := s.hasher.Hash(input.Password)
	if err != nil {
		return "", err
	}

	// Create a new user domain object
	user := domain.User{
		Name:     input.Name,
		Password: passwordHash,
		Email:    input.Email,
		// ID:       uuid.New(), // Generate a new UUID
	}

	// Save the user to the repository

	id, err := s.repo.Create(ctx, &user)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

// SignIn authenticates a user and generates a JSON Web Token (JWT).
// The JWT contains the user's email as the subject and expires in 72 hours.
//
// Parameters:
// - ctx: The context.Context object for handling the request.
// - input: The UserSignInInput object containing the user's email and password.
//
// Returns:
// - string: The generated JWT if the sign in is successful.
// - error: An error if the sign in is not successful.
func (s *UserService) SignIn(ctx context.Context, input UserSignInInput) (string, error) {
	hashedPassword, err := s.hasher.Hash(input.Password)
	if err != nil {
		return "", err
	}

	user, err := s.repo.FindByEmail(input.Email)
	if err != nil {
		return "", err
	}

	if user.Password != hashedPassword {
		return "", err
	}

	expirationTime := time.Now().Add(time.Hour * 72)
	payload := jwt.MapClaims{
		"sub": input.Email,
		"exp": expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	secretKey := os.Getenv("SECRET_KEY")
	jwtSecret := []byte(secretKey)

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *UserService) Update(ctx context.Context, input UserUpdateInput) error {
	user, err := s.repo.FindByID(input.ID.String())
	if err != nil {
		return err
	}

	err = s.repo.Update(ctx, &domain.User{
		// ID:       user.ID,
		Name:     input.Name,
		Email:    input.Email,
		Password: user.Password,
	})

	return err
}

func (s *UserService) FindByEmail(email string) (*domain.User, error) {
	user, err := s.repo.FindByEmail(email)

	return user, err
}

func (s *UserService) FindById(id string) (*domain.User, error) {
	user, err := s.repo.FindByID(id)

	return user, err
}

func (s *UserService) FindAll() ([]domain.User, error) {
	users, err := s.repo.FindAll()

	return users, err
}

func (s *UserService) Delete(userId string) error {
	return s.repo.Delete(userId)
}
