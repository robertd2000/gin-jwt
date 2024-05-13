package user_repository

import (
	"context"
	"errors"
	"go-jwt/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// NewUsersRepo creates a new instance of UsersRepo.
//
// It takes a *gorm.DB as a parameter and returns a pointer to a UsersRepo struct.
// If the provided *gorm.DB is nil, it will panic.
//
// Parameters:
//   - db: A pointer to a *gorm.DB. It cannot be nil.
//
// Returns:
//   - A pointer to a UsersRepo struct.
func NewUsersRepo(db *gorm.DB) *UsersRepo {
	if db == nil {
		panic("db cannot be nil")
	}

	return &UsersRepo{
		db: db,
	}
}

// Create adds a new user to the UsersRepo and returns the ID of the created user.
func (repo *UsersRepo) Create(ctx context.Context, user *domain.User) (uuid.UUID, error) {
	if repo == nil || user == nil {
		return uuid.Nil, errors.New("invalid arguments")
	}

	var existingUser domain.User
	if err := repo.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return uuid.Nil, errors.New("User with this credentials already exists")
	}

	if err := repo.db.Create(user).Error; err != nil {
		return uuid.Nil, err
	}

	return user.ID, nil
}

// Update updates an existing user in the UsersRepo.
func (repo *UsersRepo) Update(ctx context.Context, user *domain.User) error {
	if repo == nil {
		return errors.New("repository is nil")
	}
	if user == nil {
		return errors.New("user is nil")
	}

	if err := repo.db.Save(user).Error; err != nil {
		return errors.New(err.Error())
	}

	return nil
}

// FindByEmail retrieves a user by email from the UsersRepo.
//
// Takes an email string as a parameter.
// Returns a pointer to a domain.User and an error.
func (repo *UsersRepo) FindByEmail(email string) (*domain.User, error) {
	user := new(domain.User)
	err := repo.db.Model(&domain.User{}).Preload("Objects").Where("email = ?", email).First(user).Error
	return user, err
}

// FindByID retrieves a user by ID from the database.
//
// Parameters:
//
//	id - the ID of the user to find.
//
// Returns:
//
//	*domain.User - the user found.
//	error - an error, if any.
func (repo *UsersRepo) FindByID(id string) (*domain.User, error) {
	user := &domain.User{ID: uuid.MustParse(id)}
	err := repo.db.Model(&domain.User{}).Preload("Objects").First(user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

// FindAll retrieves all users from the UsersRepo.
//
// None.
// ([]domain.User, error)
func (repo *UsersRepo) FindAll() ([]domain.User, error) {
	var users []domain.User
	err := repo.db.Model(&domain.User{}).Preload("Objects").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// AddObject adds an object to a user in the UsersRepo.
//
// Parameters:
//
//	user domain.User - the user to add the object to.
//	object domain.Object - the object to be added.
//
// Return type: error
func (repo *UsersRepo) AddObject(user domain.User, object domain.Object) error {
	repo.db.Model(&user).Association("Languages").Append(object)

	return nil
}

// Delete removes a user and their associated objects from the repository.
//
// Parameters:
//
//	id - the ID of the user to be deleted.
//
// Returns:
//
//	error - if any occurred during the deletion process.
func (repo *UsersRepo) Delete(id string) error {
	err := repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&domain.User{}, "id = ?", id).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}
