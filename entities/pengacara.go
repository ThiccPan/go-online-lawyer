package entities

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Pengacara struct {
	gorm.Model
	Username string `json:"Username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Category string `json:"category"`
	Price decimal.Decimal `json:"price" gorm:"type:decimal(20,2);"`
	KonsultasiList []Konsultasi
}

type PengacaraDTO struct {
	Username string `json:"Username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Category string `json:"category"`
	Price decimal.Decimal `json:"price" gorm:"type:decimal(20,2);"`
}