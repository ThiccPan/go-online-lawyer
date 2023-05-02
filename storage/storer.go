package storage

import "github.com/thiccpan/go-online-lawyer/entities"

type PengacaraStorer interface {
	GetAll() ([]entities.Pengacara, error)
	GetByID()
	Insert()
	Update()
	Delete()
}