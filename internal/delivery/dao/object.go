package dao

import "github.com/google/uuid"

type ObjectCreateInput struct {
	Name        string    `json:"name" binding:"required,min=2,max=64"`
	Type        int       `json:"type" bson:"type"`
	Coords      string    `json:"coords" bson:"coords"`
	Radius      int       `json:"radius" bson:"radius"`
	Description string    `json:"description" bson:"description"`
	Color       string    `json:"color" bson:"color"`
	UserID      uuid.UUID `json:"userId" bson:"userId"`
}

type ObjectUpdateInput struct {
	ID          uuid.UUID `json:"id" bson:"id"`
	Name        string    `json:"name" binding:"required,min=2,max=64"`
	Type        int       `json:"type" bson:"type"`
	Coords      string    `json:"coords" bson:"coords"`
	Radius      int       `json:"radius" bson:"radius"`
	Description string    `json:"description" bson:"description"`
	Color       string    `json:"color" bson:"color"`
	UserID      uuid.UUID `json:"userId" bson:"userId"`
}