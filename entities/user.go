package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserDTO struct {
	Username string `json:"username" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserLoginDTO struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}