package usecases

import (
	"time"

	"github.com/thiccpan/go-online-lawyer/entities"
	"github.com/thiccpan/go-online-lawyer/storage"
)

type Konsultasi interface {
	GetKonsultasiByUserId(userId uint) ([]entities.Konsultasi, error)
	CreateKonsultasi(userId, pengacaraId uint, time time.Time) (entities.Konsultasi, error)
}

type konsultasi struct {
	KonsultasiStorer storage.KonsultasiStorer
}

func NewKonsultasiUsecase(konsultasiStorer storage.KonsultasiStorer) *konsultasi {
	return &konsultasi{
		KonsultasiStorer: konsultasiStorer,
	}
}

func (k *konsultasi) GetAllKonsultasi() ([]entities.Konsultasi, error) {
	return k.KonsultasiStorer.GetAllKonsultasi()
}

func (k *konsultasi) GetKonsultasiByUserId(userId uint) ([]entities.Konsultasi, error) {
	return k.KonsultasiStorer.GetKonsultasiByUserId(userId)
}

func (k *konsultasi) CreateKonsultasi(userId, pengacaraId uint, time time.Time) (entities.Konsultasi, error) {
	return k.KonsultasiStorer.CreateKonsultasi(userId, pengacaraId, time)
}