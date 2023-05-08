package entities

import (
	"time"

	"github.com/thiccpan/go-online-lawyer/constants"
	"gorm.io/gorm"
)

type Konsultasi struct {
	gorm.Model
	UserId uint `json:"userId"`
	PengacaraId uint `json:"pengacaraId"`
	Status constants.KonsultasiStatus `gorm:"konsultasiStatus"`
	KonsultasiTime time.Time `json:"konsultasiTime"`
	Link string `json:"link" gorm:"type:text"`
}

type KonsultasiDTO struct {
	UserId uint `json:"userId" validate:"required"`
	PengacaraId uint `json:"pengacaraId" validate:"required"`
	Status constants.KonsultasiStatus `validate:"required" gorm:"konsultasiStatus"`
	KonsultasiTime string `json:"konsultasiTime" validate:"required"`
	Link string `json:"link" gorm:"type:text"`
}

type EditKonsultasiDTO struct {
	KonsultasiTime string `json:"konsultasiTime" validate:"required"`
}