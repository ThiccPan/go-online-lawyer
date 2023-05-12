package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/thiccpan/go-online-lawyer/entities"
	"github.com/thiccpan/go-online-lawyer/storage"
	"gorm.io/gorm"
)

func TestCreateRating(t *testing.T) {
	testCases := []struct {
		desc         string
		ratingDataIn entities.RatingPengacara
		expectedErr  error
		expectedRet  entities.RatingPengacara
	}{
		{
			desc: "success",
			ratingDataIn: entities.RatingPengacara{
				UserId:      1,
				PengacaraId: 1,
				Rating:      5,
			},
			expectedErr: nil,
			expectedRet: entities.RatingPengacara{
				UserId:      1,
				PengacaraId: 1,
				Rating:      5,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mockStorer := storage.NewMockRatingPengacara()
			mockStorer.On("CreateRating", mock.Anything, mock.Anything).Return(tC.expectedRet, tC.expectedErr)
			usecase := NewRatingPengacaraUsecase(mockStorer)
			data, err := usecase.CreateRating(tC.ratingDataIn.UserId, tC.ratingDataIn)
			assert.Equal(t, tC.expectedErr, err)
			assert.Equal(t, tC.expectedRet.Rating, data.Rating)
		})
	}
}

func TestGetAllRatingByUser(t *testing.T) {
	testCases := []struct {
		desc        string
		userIdIn    uint
		expectedErr error
		expectedRet []entities.RatingPengacara
	}{
		{
			desc:        "success",
			userIdIn:    1,
			expectedErr: nil,
			expectedRet: []entities.RatingPengacara{
				{
					UserId:      1,
					PengacaraId: 1,
					Rating:      5,
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mockStorer := storage.NewMockRatingPengacara()
			mockStorer.On("GetAllRatingByUser", mock.Anything).Return(tC.expectedRet, tC.expectedErr)
			usecase := NewRatingPengacaraUsecase(mockStorer)
			data, err := usecase.GetAllRatingByUser(tC.userIdIn)
			assert.Equal(t, tC.expectedErr, err)
			assert.Equal(t, tC.expectedRet, data)
		})
	}
}

func TestGetAllRatingByPengacara(t *testing.T) {
	testCases := []struct {
		desc          string
		penfacaraIdIn uint
		expectedErr   error
		expectedRet   []entities.RatingPengacara
	}{
		{
			desc:          "success",
			penfacaraIdIn: 1,
			expectedErr:   nil,
			expectedRet: []entities.RatingPengacara{
				{
					UserId:      1,
					PengacaraId: 1,
					Rating:      5,
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mockStorer := storage.NewMockRatingPengacara()
			mockStorer.On("GetAllRatingByPengacara", mock.Anything).Return(tC.expectedRet, tC.expectedErr)
			usecase := NewRatingPengacaraUsecase(mockStorer)
			data, err := usecase.GetAllRatingByPengacara(tC.penfacaraIdIn)
			assert.Equal(t, tC.expectedErr, err)
			assert.Equal(t, tC.expectedRet, data)
		})
	}
}

func TestDeleteRating(t *testing.T) {
	testCases := []struct {
		desc	string
		idPayloadIn uint
		ratingIdIn uint
		expectedErr error
		expectedRet entities.RatingPengacara
	}{
		{
			desc: "success",
			idPayloadIn: 1,
			ratingIdIn: 1,
			expectedErr: nil,
			expectedRet: entities.RatingPengacara{
				Model: gorm.Model{
					ID: 1,
				},
				UserId:      1,
				PengacaraId: 1,
				Rating:      5,
			},
			
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mockStorer := storage.NewMockRatingPengacara()
			mockStorer.On("DeleteRating", mock.Anything, mock.Anything).Return(tC.expectedRet, tC.expectedErr)
			usecase := NewRatingPengacaraUsecase(mockStorer)
			data, err := usecase.DeleteRating(tC.idPayloadIn, tC.ratingIdIn)
			assert.Equal(t, tC.expectedErr, err)
			assert.Equal(t, tC.expectedRet.ID, data.ID)
		})
	}
}