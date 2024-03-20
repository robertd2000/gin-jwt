package service

import (
	"context"
	"fmt"
	"go-jwt/internal/domain"
	"go-jwt/internal/repository"

	"github.com/google/uuid"
)

type ObjectService struct {
	repo repository.Object
}

func NewObjectService(repo repository.Object) *ObjectService {
	return &ObjectService{
		repo,
	}
}

func (s *ObjectService) Create(ctx context.Context, objectInput ObjectCreateInput) error {
	fmt.Println("objectInput", objectInput)
	object := domain.Object{
		ID:          uuid.New(),
		Name:        objectInput.Name,
		Type:        objectInput.Type,
		Coords:      objectInput.Coords,
		Radius:      objectInput.Radius,
		Description: objectInput.Description,
		Color:       objectInput.Color,
		UserID:      objectInput.UserID,
	}

	if err := s.repo.Create(ctx, &object); err != nil {
		return err
	}

	return nil
}
