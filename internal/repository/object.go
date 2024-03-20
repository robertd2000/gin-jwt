package repository

import "gorm.io/gorm"

type ObjectRepo struct {
	db *gorm.DB
}

func NewObjectRepo(db *gorm.DB) *ObjectRepo {
	return &ObjectRepo{
		db: db,
	}
}
