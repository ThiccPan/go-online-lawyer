package storage

import "github.com/thiccpan/go-online-lawyer/entities"

type PengacaraStorer interface {
	GetAll() ([]entities.Pengacara, error)
	GetById(id int) (entities.Pengacara, error)
	GetByEmail(email string) (entities.Pengacara, error)
	GetByCategory(category string) (entities.Pengacara, error)
	Insert()
	Update()
	Delete()
}