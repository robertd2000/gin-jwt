package service

import (
	"context"
	"go-jwt/internal/domain"
	"go-jwt/internal/repository"

	"github.com/google/uuid"
)

type ObjectService struct {
	objectRepo repository.Object
	userRepo   repository.User
}

func NewObjectService(objectRepo repository.Object, userEepo repository.User) *ObjectService {
	return &ObjectService{
		objectRepo,
		userEepo,
	}
}

func (s *ObjectService) Create(ctx context.Context, objectInput ObjectCreateInput) error {
	user, err := s.userRepo.FindByID(objectInput.UserID.String())

	if err != nil {
		return err
	}

	object := domain.Object{
		ID:          uuid.New(),
		Name:        objectInput.Name,
		Type:        objectInput.Type,
		Coords:      objectInput.Coords,
		Radius:      objectInput.Radius,
		Description: objectInput.Description,
		Color:       objectInput.Color,
		UserID:      objectInput.UserID,
		User:        *user,
	}

	err = s.userRepo.AddObject(*user, object)
	if err != nil {
		return err
	}

	if err := s.objectRepo.Create(ctx, &object); err != nil {
		return err
	}

	return nil
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
