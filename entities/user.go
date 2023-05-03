package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"Username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserDTO struct {
	Username string `json:"Username"`
	Email string `json:"email"`
	Password string `json:"password"`
}