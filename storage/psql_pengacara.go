package storage

import (
	"github.com/thiccpan/go-online-lawyer/entities"
	"gorm.io/gorm"
)

type Pengacara struct {
	DB *gorm.DB
}

func NewPengacaraStorer(DB *gorm.DB) *Pengacara {
	return &Pengacara{
		DB: DB,
	}
}

func (p *Pengacara) GetAll() ([]entities.Pengacara, error) {
	var pengacaras []entities.Pengacara
	res := p.DB.Find(&pengacaras)
	if res.Error != nil {
		return nil, res.Error
	}
	return pengacaras, nil
}

func (p *Pengacara) GetByID() {}
func (p *Pengacara) Insert() {}
func (p *Pengacara) Update() {}
func (p *Pengacara) Delete() {}