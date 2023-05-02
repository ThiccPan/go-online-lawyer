package entities

import (
	"gorm.io/gorm"
)

type Pengacara struct {
	gorm.Model
	Username string `json:"Username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Category string `json:"category"`
}

type PengacaraDTO struct {
	Username string `json:"Username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Category string `json:"category"`
}