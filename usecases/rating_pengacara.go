package usecases

import (
	"github.com/thiccpan/go-online-lawyer/entities"
	"github.com/thiccpan/go-online-lawyer/storage"
)

type RatingPengacara interface {
	CreateRating(userId uint, ratingData entities.RatingPengacara) (entities.RatingPengacara, error)
	GetAllRatingByUser(userId uint) ([]entities.RatingPengacara, error)
}

type ratingPengacara struct {
	Storer storage.RatingPengacaraStorer
}

func NewRatingPengacaraUsecase(storer storage.RatingPengacaraStorer) *ratingPengacara {
	return &ratingPengacara{
		Storer: storer,
	}
}

func (rp *ratingPengacara) CreateRating(userId uint, ratingData entities.RatingPengacara) (entities.RatingPengacara, error) {
	ratingData.UserId = userId
	return rp.Storer.CreateRating(ratingData)
}

func (rp *ratingPengacara) GetAllRatingByUser(userId uint) ([]entities.RatingPengacara, error) {
	return rp.Storer.GetAllRatingByUser(userId)
}