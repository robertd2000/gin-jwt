package domain

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" bson:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name     string    `json:"name" bson:"name" gorm:"UNIQUE_INDEX:compositeindex;index"`
	Email    string    `json:"email" bson:"email" gorm:"UNIQUE_INDEX:compositeindex;index"`
	Password string    `json:"-" bson:"password" gorm:"UNIQUE_INDEX:compositeindex;index"`
	Updated  int64     `json:"updated_at" bson:"-" gorm:"autoUpdateTime:milli"`
	Created  int64     `json:"created_at" bson:"-" gorm:"autoCreateTime:milli"`
	Objects  []Object  `json:"objects" bson:"objects" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
