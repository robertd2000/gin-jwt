package domain

import "github.com/google/uuid"

type User struct {
	ID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name             string               `json:"name" bson:"name"`
	Email            string               `json:"email" bson:"email"`
	Password         string               `json:"password" bson:"password"`
} 