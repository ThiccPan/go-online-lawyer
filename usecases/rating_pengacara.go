package usecases

import (
	"github.com/thiccpan/go-online-lawyer/entities"
	"github.com/thiccpan/go-online-lawyer/exceptions"
	"github.com/thiccpan/go-online-lawyer/storage"
)

type RatingPengacara interface {
	CreateRating(userId uint, ratingData entities.RatingPengacara) (entities.RatingPengacara, error)
	GetAllRatingByUser(userId uint) ([]entities.RatingPengacara, error)
	GetAllRatingByPengacara(pengacaraId uint) ([]entities.RatingPengacara, error)
	DeleteRating(idPayload uint, ratingId uint) (entities.RatingPengacara, error)
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

func (rp *ratingPengacara) GetAllRatingByPengacara(pengacaraId uint) ([]entities.RatingPengacara, error) {
	return rp.Storer.GetAllRatingByPengacara(pengacaraId)
}

func (rp *ratingPengacara) DeleteRating(idPayload uint, ratingId uint) (entities.RatingPengacara, error) {
	data, err := rp.Storer.DeleteRating(idPayload, ratingId)
	if data.ID == 0 {
		return entities.RatingPengacara{}, exceptions.ErrRatingNotFound
	}
	return data, err
}