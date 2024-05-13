package domain

import (
	"time"

	"github.com/google/uuid"
)

type Object struct {
	ID          uuid.UUID `json:"id" bson:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name        string    `json:"name" bson:"name" gorm:"UNIQUE_INDEX:compositeindex;index"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at,omitempty" gorm:"autoCreateTime:milli"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at,omitempty" gorm:"autoUpdateTime:milli"`
	Type        int       `json:"type" bson:"type"`
	Coords      string    `json:"coords" bson:"coords"`
	Radius      int       `json:"radius" bson:"radius"`
	Description string    `json:"description" bson:"description"`
	Color       string    `json:"color" bson:"color"`
	UserID      uuid.UUID `json:"userId" bson:"userId" gorm:"type:uuid;index"`
	User        User      `json:"user" bson:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
