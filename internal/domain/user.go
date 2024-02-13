package domain

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id" bson:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name     string    `json:"name" bson:"name" gorm:"UNIQUE_INDEX:compositeindex;index"`
	Email    string    `json:"email" bson:"email" gorm:"UNIQUE_INDEX:compositeindex;index"`
	Password string    `json:"-" bson:"password" gorm:"UNIQUE_INDEX:compositeindex;index"`
}
