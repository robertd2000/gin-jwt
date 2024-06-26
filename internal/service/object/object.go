package object_service

import (
	"context"
	"go-jwt/internal/domain"
	object_repository "go-jwt/internal/repository/object"
	user_repository "go-jwt/internal/repository/user"

	"github.com/google/uuid"
)

func NewObjectService(objectRepo object_repository.Object, userEepo user_repository.User) *ObjectService {
	return &ObjectService{
		objectRepo,
		userEepo,
	}
}

// Create creates a new object in the database.
//
// The function takes a context and an ObjectCreateInput struct as parameters.
// It returns an error if any of the underlying repository methods fail.
func (s *ObjectService) Create(ctx context.Context, input ObjectCreateInput) (uuid.UUID, error) {
	// Find the user by their ID
	user, err := s.userRepo.FindByID(input.UserID.String())
	if err != nil {
		return uuid.Nil, err
	}

	// Create a new object with the provided input data
	object := domain.Object{
		Name:        input.Name,
		Type:        input.Type,
		Coords:      input.Coords,
		Radius:      input.Radius,
		Description: input.Description,
		Color:       input.Color,
		UserID:      input.UserID,
		User:        *user, // Assign the found user to the object
	}

	// Add the object to the user in the user repository
	if err := s.userRepo.AddObject(*user, object); err != nil {
		return uuid.Nil, err
	}

	// Create the object in the object repository
	id, err := s.objectRepo.Create(ctx, &object)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (s *ObjectService) Update(ctx context.Context, objectInput ObjectUpdateInput) error {
	user, err := s.userRepo.FindByID(objectInput.UserID.String())

	if err != nil {
		return err
	}

	object := domain.Object{
		ID:          objectInput.ID,
		Name:        objectInput.Name,
		Type:        objectInput.Type,
		Coords:      objectInput.Coords,
		Radius:      objectInput.Radius,
		Description: objectInput.Description,
		Color:       objectInput.Color,
		UserID:      objectInput.UserID,
		User:        *user,
	}

	// err = s.userRepo.AddObject(*user, object)
	// if err != nil {
	// 	return err
	// }

	if err := s.objectRepo.Update(ctx, &object); err != nil {
		return err
	}

	return nil
}

func (s *ObjectService) Delete(objectId string) error {
	return s.objectRepo.Delete(objectId)
}

func (s *ObjectService) FindAll() ([]domain.Object, error) {
	return s.objectRepo.FindAll()
}

func (s *ObjectService) FindById(id string) (*domain.Object, error) {
	return s.objectRepo.FindById(id)
}

func (s *ObjectService) FindByUserId(userId string) ([]domain.Object, error) {
	return s.objectRepo.FindByUserId(userId)
}
