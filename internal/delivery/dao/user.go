package dao

import "github.com/google/uuid"

type UserSignUpInput struct {
	Name     string `json:"name" binding:"required,min=2,max=64"`
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

type UserUpdateInput struct {
	ID    uuid.UUID `json:"id" binding:"required"`
	Name  string    `json:"name" binding:"required,min=2,max=64"`
	Email string    `json:"email" binding:"required,email,max=64"`
}
