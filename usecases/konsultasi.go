package usecases

import (
	"fmt"

	"github.com/thiccpan/go-online-lawyer/entities"
	"github.com/thiccpan/go-online-lawyer/storage"
	"github.com/thiccpan/go-online-lawyer/validations"
)

type Konsultasi interface {
	GetAllUserKonsultasi(userId uint) ([]entities.Konsultasi, error)
	GetUserKonsultasi(userId uint, konsultasiId uint) (entities.Konsultasi, error)
	CreateKonsultasi(konsultasiData entities.Konsultasi) (entities.Konsultasi, error)
	EditKonsultasi(konsultasiId uint, userId uint, konsultasiData entities.Konsultasi) (entities.Konsultasi, error)
	DeleteKonsultasi(konsultasiId uint, userId uint) (entities.Konsultasi, error)
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

func (k *konsultasi) GetAllUserKonsultasi(userId uint) ([]entities.Konsultasi, error) { //TODO: rename
	return k.KonsultasiStorer.GetAllUserKonsultasi(userId)
}

func (k *konsultasi) GetUserKonsultasi(userId uint, konsultasiId uint) (entities.Konsultasi, error) {
	return k.KonsultasiStorer.GetUserKonsultasi(userId, konsultasiId)
}

func (k *konsultasi) CreateKonsultasi(konsultasiData entities.Konsultasi) (entities.Konsultasi, error) {
	return k.KonsultasiStorer.CreateKonsultasi(konsultasiData)
}

func (k *konsultasi) EditKonsultasi(konsultasiId uint, userId uint, konsultasiData entities.Konsultasi) (entities.Konsultasi, error) {
	data := entities.Konsultasi{
		KonsultasiTime: konsultasiData.KonsultasiTime,
	}
	fmt.Println(data)
	return k.KonsultasiStorer.EditKonsultasi(konsultasiId, userId, data)
}

func (k *konsultasi) DeleteKonsultasi(konsultasiId uint, userId uint,) (entities.Konsultasi, error) {
	return k.KonsultasiStorer.DeleteKonsultasi(konsultasiId, userId)
}