package usecases

import (
	"github.com/thiccpan/go-online-lawyer/entities"
	"github.com/thiccpan/go-online-lawyer/storage"
)

type Pengacara interface {
	GetAll() ([]entities.Pengacara, error)
	// GetByEmail(email string) (entities.Pengacara, error)
	// Create(data entities.PengacaraDTO) error
	// Delete(email string) error
	// Update(email string) error
}

type pengacara struct {
	PengacaraStorer storage.PengacaraStorer
}

func New(pengacaraStorage storage.PengacaraStorer) *pengacara {
	return &pengacara{
		PengacaraStorer: pengacaraStorage,
	}
}

func (p *pengacara) GetAll() ([]entities.Pengacara, error) {
	return p.PengacaraStorer.GetAll()
}

func GetByEmail(email string) (entities.Pengacara, error) {
	panic("unimplemented")
}

func Create(data entities.PengacaraDTO) error {
	panic("unimplemented")
}

func Delete(email string) error {
	panic("unimplemented")
}

func Update(email string) error {
	panic("unimplemented")
}