package storage

import (
	"github.com/thiccpan/go-online-lawyer/entities"
	"gorm.io/gorm"
)

type RatingPengacaraStorer interface {
	CreateRating(ratingData entities.RatingPengacara) (entities.RatingPengacara, error)
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

