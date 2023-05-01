package usecases

import (
	"github.com/thiccpan/go-online-lawyer/entities"
	"github.com/thiccpan/go-online-lawyer/storage"
)

type Pengacara interface {
	GetAll() ([]entities.Pengacara, error)
	GetByID() (entities.Pengacara, error)
	Create() error
	Delete() error
	Update() error
}

type pengacara struct {
	PengacaraStorage storage.Pengacara
}

func New(pengacaraStorage storage.Pengacara) *pengacara {
	return &pengacara{
		PengacaraStorage: pengacaraStorage,
	}
}
