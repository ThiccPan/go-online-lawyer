package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserDTO struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}