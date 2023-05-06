package storage

import (
	"github.com/thiccpan/go-online-lawyer/constants"
	"github.com/thiccpan/go-online-lawyer/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type KonsultasiStorer interface {
	GetAllKonsultasi() ([]entities.Konsultasi, error)
	GetKonsultasiByUserId(userId uint) ([]entities.Konsultasi, error)
	CreateKonsultasi(konsultasiData entities.Konsultasi) (entities.Konsultasi, error)
	EditKonsultasi(id uint, data entities.Konsultasi) (entities.Konsultasi, error)
	DeleteKonsultasi(id uint) (entities.Konsultasi, error)
}

type konsultasiStorer struct {
	DB *gorm.DB
}

func NewKonsultasiStorer(db *gorm.DB) *konsultasiStorer {
	return &konsultasiStorer{
		DB: db,
	}
}

func (k *konsultasiStorer) CreateKonsultasi(konsultasiData entities.Konsultasi) (entities.Konsultasi, error) {
	konsultasiRequest := entities.Konsultasi{
		UserId: konsultasiData.UserId,
		PengacaraId: konsultasiData.PengacaraId,
		Status: constants.DIPROSES,
		KonsultasiTime: konsultasiData.KonsultasiTime,
		Link: "",
	}
	res := k.DB.Create(&konsultasiRequest)
	if res.Error != nil {
		return entities.Konsultasi{}, res.Error
	}
	return konsultasiRequest, nil
}

func (k *konsultasiStorer) GetAllKonsultasi() ([]entities.Konsultasi, error) {
	var daftarKonsultasi []entities.Konsultasi
	res := k.DB.Find(&daftarKonsultasi)
	if res != nil {
		return nil, res.Error
	}
	return daftarKonsultasi, nil
}

func (k *konsultasiStorer) GetKonsultasiByUserId(userId uint) ([]entities.Konsultasi, error) {
	var daftarKonsultasi []entities.Konsultasi
	res := k.DB.Debug().Where("user_id = ?", userId).Find(&daftarKonsultasi)
	if res.Error != nil {
		return nil, res.Error
	}
	return daftarKonsultasi, nil
}

func (k *konsultasiStorer) EditKonsultasi(id uint, data entities.Konsultasi) (entities.Konsultasi, error) {
	var res entities.Konsultasi
	err := k.DB.Where("id = ?", id).First(&res).Error
	if err != nil {
		return entities.Konsultasi{}, err
	}
	res.KonsultasiTime = data.KonsultasiTime
	err = k.DB.Save(&res).Error
	if err != nil {
		return entities.Konsultasi{}, err
	}
	return res, nil
}

func (k *konsultasiStorer) DeleteKonsultasi(id uint) (entities.Konsultasi, error) {
	deletedData := entities.Konsultasi{}
	err := k.DB.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&deletedData).Error
	return deletedData, err
}