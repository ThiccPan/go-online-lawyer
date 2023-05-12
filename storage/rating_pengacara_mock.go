package storage

import (
	"github.com/stretchr/testify/mock"
	"github.com/thiccpan/go-online-lawyer/entities"
)

type mockRatingPengacara struct {
	mock.Mock	
}

func NewMockRatingPengacara() *mockRatingPengacara {
	return &mockRatingPengacara{}
}

func (mr *mockRatingPengacara) CreateRating(ratingData entities.RatingPengacara) (entities.RatingPengacara, error) {
	ret := mr.Mock.Called(ratingData)
	return ret.Get(0).(entities.RatingPengacara), ret.Error(1)
}
func (mr *mockRatingPengacara) GetAllRatingByUser(userId uint) ([]entities.RatingPengacara, error) {
	ret := mr.Mock.Called(userId)
	return ret.Get(0).([]entities.RatingPengacara), ret.Error(1)
}
func (mr *mockRatingPengacara) GetAllRatingByPengacara(pengacaraId uint) ([]entities.RatingPengacara, error) {
	ret := mr.Mock.Called(pengacaraId)
	return ret.Get(0).([]entities.RatingPengacara), ret.Error(1)
}
func (mr *mockRatingPengacara) DeleteRating(userId uint, ratingId uint) (entities.RatingPengacara, error) {
	ret := mr.Mock.Called(userId, ratingId)
	return ret.Get(0).(entities.RatingPengacara), ret.Error(1)
}