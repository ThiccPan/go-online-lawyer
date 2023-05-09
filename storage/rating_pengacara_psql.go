package storage

import (
	"github.com/thiccpan/go-online-lawyer/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RatingPengacaraStorer interface {
	CreateRating(ratingData entities.RatingPengacara) (entities.RatingPengacara, error)
	GetAllRatingByUser(userId uint) ([]entities.RatingPengacara, error)
	GetAllRatingByPengacara(pengacaraId uint) ([]entities.RatingPengacara, error)
	DeleteRating(userId uint, ratingId uint) (entities.RatingPengacara, error)
}

type ratingPengacaraStorer struct {
	DB *gorm.DB
}

func NewRatingPengacaraStorer(DB *gorm.DB) *ratingPengacaraStorer {
	return &ratingPengacaraStorer{
		DB: DB,
	}
}

func (rps *ratingPengacaraStorer) CreateRating(ratingData entities.RatingPengacara) (entities.RatingPengacara, error) {
	err := rps.DB.Create(&ratingData).Error
	if err != nil {
		return entities.RatingPengacara{}, err
	}
	return ratingData, nil
}

func (rps *ratingPengacaraStorer) GetAllRatingByUser(userId uint) ([]entities.RatingPengacara, error) {
	daftarRating := []entities.RatingPengacara{}
	err := rps.DB.Where("user_id = ?", userId).Find(&daftarRating).Error
	if err != nil {
		return []entities.RatingPengacara{}, nil
	}
	return daftarRating, nil
}

func (rps *ratingPengacaraStorer) GetAllRatingByPengacara(pengacaraId uint) ([]entities.RatingPengacara, error) {
	daftarRating := []entities.RatingPengacara{}
	err := rps.DB.Where("pengacara_id = ?", pengacaraId).Find(&daftarRating).Error
	if err != nil {
		return []entities.RatingPengacara{}, nil
	}
	return daftarRating, nil
}

func (rps *ratingPengacaraStorer) DeleteRating(userId uint, ratingId uint) (entities.RatingPengacara, error) {
	deletedData := entities.RatingPengacara{}
	err := rps.DB.Clauses(clause.Returning{}).Where("user_id = ?", userId).Where("id = ?", ratingId).Delete(&deletedData).Error
	if err != nil {
		return entities.RatingPengacara{}, err
	}
	return deletedData, nil
}