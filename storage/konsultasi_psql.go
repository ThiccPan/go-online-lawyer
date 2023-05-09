package storage

import (
	"github.com/thiccpan/go-online-lawyer/constants"
	"github.com/thiccpan/go-online-lawyer/entities"
	"github.com/thiccpan/go-online-lawyer/exceptions"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type KonsultasiStorer interface {
	GetAllKonsultasi() ([]entities.Konsultasi, error)
	GetAllUserKonsultasi(userId uint) ([]entities.Konsultasi, error)
	GetUserKonsultasi(userId uint, konsultasiId uint) (entities.Konsultasi, error)
	CreateKonsultasi(konsultasiData entities.Konsultasi) (entities.Konsultasi, error)
	EditKonsultasi(konsultasiId uint, userId uint, data entities.Konsultasi) (entities.Konsultasi, error)
	EditKonsultasiPengacara(konsultasiId uint, konsultasiData entities.Konsultasi) (entities.Konsultasi, error)
	DeleteKonsultasi(konsultasiId uint, userId uint) (entities.Konsultasi, error)
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

func (k *konsultasiStorer) GetAllUserKonsultasi(userId uint) ([]entities.Konsultasi, error) {
	var daftarKonsultasi []entities.Konsultasi
	res := k.DB.Debug().Where("user_id = ?", userId).Find(&daftarKonsultasi)
	if res.Error != nil {
		return nil, res.Error
	}
	return daftarKonsultasi, nil
}

func (k *konsultasiStorer) GetUserKonsultasi(userId uint, konsultasiId uint) (entities.Konsultasi, error) {
	var dataKonsultasi entities.Konsultasi
	res := k.DB.Where("user_id = ?", userId).Where("id = ?", konsultasiId).First(&dataKonsultasi)
	if res.Error != nil {
		return entities.Konsultasi{}, res.Error
	}
	return dataKonsultasi, nil
}

func (k *konsultasiStorer) EditKonsultasi(konsultasiId uint, userId uint, data entities.Konsultasi) (entities.Konsultasi, error) {
	var res entities.Konsultasi
	err := k.DB.Where("user_id = ?", userId).Where("id = ?", konsultasiId).First(&res).Error
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
func (k *konsultasiStorer) EditKonsultasiPengacara(konsultasiId uint, konsultasiData entities.Konsultasi) (entities.Konsultasi, error) {
	var res entities.Konsultasi
	err := k.DB.Where("pengacara_id = ?", konsultasiData.PengacaraId).Where("id = ?", konsultasiId).First(&res).Error
	if err != nil {
		return entities.Konsultasi{}, err
	}

	res.Status = konsultasiData.Status
	res.Link = konsultasiData.Link
	if !konsultasiData.KonsultasiTime.IsZero() {
		res.KonsultasiTime = konsultasiData.KonsultasiTime
	}
	err = k.DB.Save(&res).Error
	if err != nil {
		return entities.Konsultasi{}, err
	}

	return res, nil
}

func (k *konsultasiStorer) DeleteKonsultasi(konsultasiId uint, userId uint) (entities.Konsultasi, error) {
	deletedData := entities.Konsultasi{}
	err := k.DB.Clauses(clause.Returning{}).Where("user_id = ?", userId).Where("id = ?", konsultasiId).Delete(&deletedData).Error
	if deletedData.ID == 0 {
		err = exceptions.ErrKonsultasiNotFound
	}
	return deletedData, err
}