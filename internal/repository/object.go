package repository

import (
	"context"
	"errors"
	"go-jwt/internal/domain"

	"gorm.io/gorm"
)

type ObjectRepo struct {
	db *gorm.DB
}

func NewObjectRepo(db *gorm.DB) *ObjectRepo {
	return &ObjectRepo{
		db: db,
	}
}


// Create saves a new object to the database.
func (r *ObjectRepo) Create(ctx context.Context, obj *domain.Object) error {
	if err := r.db.Create(obj).Error; err != nil {
		return err
	}
	return nil
}


// Update updates an existing object in the database.
func (repo *ObjectRepo) Update(ctx context.Context, obj *domain.Object) error {
	if err := repo.db.Save(obj).Error; err != nil {
		return err
	}
	return nil
}

func (repo *ObjectRepo) FindAll() ([]domain.Object, error) {
	objects := []domain.Object{}

	err := repo.db.Preload("User").Find(&objects).Error
	if err != nil {
		return nil, err
	}

	return objects, nil
}


// FindById retrieves an object from the database by its ID.
//
// Parameters:
//   id - the ID of the object to be retrieved.
//
// Returns:
//   *domain.Object - the retrieved object.
//   error - an error, if any occurred during the retrieval process.
func (repo *ObjectRepo) FindById(id string) (*domain.Object, error) {
	if id == "" {
		return nil, errors.New("id can't be empty")
	}
	
	object := &domain.Object{}
	
	if err := repo.db.Where("id = ?", id).Take(object).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		
		return nil, err
	}
	
	return object, nil
}



// FindByUserId retrieves all objects belonging to a specific user.
//
// Parameters:
//   userId - the ID of the user whose objects should be retrieved.
//
// Returns:
//   []domain.Object - a slice of objects belonging to the user.
//   error - an error, if any occurred during the retrieval process.
func (repo *ObjectRepo) FindByUserId(userId string) ([]domain.Object, error) {
	// Declare an empty slice to store the objects.
	var objects []domain.Object

	// Use the gorm Where function to filter objects by user ID.
	// The Find function retrieves the matching records and stores them in the objects slice.
	// The Error function checks for any errors occurred during the retrieval process.
	err := repo.db.Where("user_id = ?", userId).Find(&objects).Error

	// If an error occurred, wrap it in a new error and return it.
	if err != nil {
		return nil, errors.New(err.Error())
	}

	// Return the retrieved objects and nil error.
	return objects, nil
}


// Delete removes an object from the database.
//
// Parameters:
//   objectId - the ID of the object to be deleted.
//
// Returns:
//   error - an error, if any occurred during the deletion process.
func (repo *ObjectRepo) Delete(objectId string) error {
	// Use gorm's Where function to filter objects by ID, and then Delete to remove the matching records.
	// The Error function checks for any errors occurred during the deletion process.
	// The &domain.Object{} part tells gorm which struct type to delete.
	res := repo.db.Where("id = ?", objectId).Delete(&domain.Object{})

	// If an error occurred, wrap it in a new error and return it.
	if res.Error != nil {
		return errors.New(res.Error.Error())
	}

	// No error occurred, so return nil.
	return nil
}





// DeleteByUserId removes all objects belonging to a specific user from the
// database in a single query, rather than one-by-one.
//
// Parameters:
//   userId - the ID of the user whose objects should be deleted.
//
// Returns:
//   error - an error, if any occurred during the deletion process.
func (repo *ObjectRepo) DeleteByUserId(userId string) error {
	// Use gorm's Unscoped function to bypass the soft delete feature, if
	// it's enabled, and delete the records in a single query.
	res := repo.db.Unscoped().Where("user_id = ?", userId).Delete(&domain.Object{})

	// If there were errors, wrap them in a new error and return it.
	if res.Error != nil {
		return errors.New(res.Error.Error())
	}

	// No errors occurred, so return nil.
	return nil
}
