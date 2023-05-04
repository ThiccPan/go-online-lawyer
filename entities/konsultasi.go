package entities

import "gorm.io/gorm"

type Konsultasi struct {
	gorm.Model
	UserId uint
	PengacaraId uint
}