package usecases

import (
	"fmt"

	"github.com/thiccpan/go-online-lawyer/entities"
	"github.com/thiccpan/go-online-lawyer/storage"
	"github.com/thiccpan/go-online-lawyer/validations"
)

type Konsultasi interface {
	GetKonsultasiByUserId(userId uint) ([]entities.Konsultasi, error)
	CreateKonsultasi(konsultasiData entities.Konsultasi) (entities.Konsultasi, error)
	EditKonsultasi(payload entities.KonsultasiDTO) (entities.Konsultasi, error)
}

type konsultasi struct {
	KonsultasiStorer storage.KonsultasiStorer
	validator validations.CustomValidator
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

func (k *konsultasi) CreateKonsultasi(konsultasiData entities.Konsultasi) (entities.Konsultasi, error) {
	return k.KonsultasiStorer.CreateKonsultasi(konsultasiData)
}

func (k *konsultasi) EditKonsultasi(payload entities.KonsultasiDTO) (entities.Konsultasi, error) {
	data := entities.Konsultasi{
		Status: payload.Status,
	}
	fmt.Println(data)
	return k.KonsultasiStorer.EditKonsultasi()
}