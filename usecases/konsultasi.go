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
	EditKonsultasi(konsultasiId uint, payload entities.Konsultasi) (entities.Konsultasi, error)
	DeleteKonsultasi(konsultasiId uint) (entities.Konsultasi, error)
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

func (k *konsultasi) EditKonsultasi(konsultasiId uint, konsultasiData entities.Konsultasi) (entities.Konsultasi, error) {
	data := entities.Konsultasi{
		KonsultasiTime: konsultasiData.KonsultasiTime,
	}
	fmt.Println(data)
	return k.KonsultasiStorer.EditKonsultasi(konsultasiId, data)
}

func (k *konsultasi) DeleteKonsultasi(konsultasiId uint) (entities.Konsultasi, error) {
	return k.KonsultasiStorer.DeleteKonsultasi(konsultasiId)
}