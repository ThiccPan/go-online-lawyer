package entities

import (
	"gorm.io/gorm"
)

type RatingPengacara struct {
	gorm.Model
	UserId uint `json:"userId"`
	PengacaraId uint `json:"pengacaraId"`
	Rating uint8 `json:"rating"` //make sure to limit to 5
}

type RatingPengacaraCreateDTO struct {
	UserId uint `json:"userId" validate:"required"`
	PengacaraId uint `json:"pengacaraId" validate:"required"`
	Rating uint8 `json:"rating" validate:"required"` //make sure to limit to 5
}