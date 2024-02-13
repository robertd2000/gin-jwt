package domain

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name     string    `gorm:"UNIQUE_INDEX:compositeindex;index"`
	Email    string    `gorm:"UNIQUE_INDEX:compositeindex;index"`
	Password string    `gorm:"UNIQUE_INDEX:compositeindex;index"`
}
