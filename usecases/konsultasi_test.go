package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/thiccpan/go-online-lawyer/constants"
	"github.com/thiccpan/go-online-lawyer/entities"
	"github.com/thiccpan/go-online-lawyer/storage"
	"gorm.io/gorm"
)

func TestGetAllKonsultasi(t *testing.T) {
	testCases := []struct {
		desc         string
		expectedErr  error
		expectedData []entities.Konsultasi
	}{
		{
			desc:        "success",
			expectedErr: nil,
			expectedData: []entities.Konsultasi{
				{
					UserId:      1,
					PengacaraId: 1,
					Status:      constants.DIPROSES,
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mockKonsultasiStorer := storage.NewMockKonsultasiStorer()
			mockKonsultasiStorer.On("GetAllKonsultasi", mock.Anything).Return([]entities.Konsultasi{
				{
					UserId:      tC.expectedData[0].UserId,
					PengacaraId: tC.expectedData[0].PengacaraId,
					Status:      tC.expectedData[0].Status,
				},
			}, tC.expectedErr)

			mockKonsultasiUsecase := NewKonsultasiUsecase(mockKonsultasiStorer)
			data, err := mockKonsultasiUsecase.GetAllKonsultasi()
			assert.Equal(t, tC.expectedErr, err)
			assert.Equal(t, tC.expectedData[0].UserId, data[0].UserId)
			assert.Equal(t, tC.expectedData[0].PengacaraId, data[0].PengacaraId)
		})
	}
}

func TestGetAllUserKonsultasi(t *testing.T) {
	testCases := []struct {
		desc         string
		userIdDTO    uint
		expectedErr  error
		expectedData []entities.Konsultasi
	}{
		{
			desc:        "success",
			userIdDTO:   1,
			expectedErr: nil,
			expectedData: []entities.Konsultasi{
				{
					UserId:      1,
					PengacaraId: 1,
					Status:      constants.DIPROSES,
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mockKonsultasiStorer := storage.NewMockKonsultasiStorer()
			mockKonsultasiStorer.On("GetAllUserKonsultasi", mock.Anything).Return([]entities.Konsultasi{
				{
					UserId:      tC.expectedData[0].UserId,
					PengacaraId: tC.expectedData[0].PengacaraId,
					Status:      tC.expectedData[0].Status,
				},
			}, tC.expectedErr)

			mockKonsultasiUsecase := NewKonsultasiUsecase(mockKonsultasiStorer)
			data, err := mockKonsultasiUsecase.GetAllUserKonsultasi(tC.userIdDTO)
			assert.Equal(t, tC.expectedErr, err)
			assert.Equal(t, tC.expectedData[0].UserId, data[0].UserId)
			assert.Equal(t, tC.expectedData[0].PengacaraId, data[0].PengacaraId)
		})
	}
}

func TestGetUserKonsultasi(t *testing.T) {
	testCases := []struct {
		desc            string
		userIdDTO       uint
		konsultasiIdDTO uint
		expectedErr     error
		expectedData    entities.Konsultasi
	}{
		{
			desc:            "success",
			userIdDTO:       1,
			konsultasiIdDTO: 1,
			expectedErr:     nil,
			expectedData: entities.Konsultasi{
				Model: gorm.Model{
					ID: 1,
				},
				UserId:      1,
				PengacaraId: 1,
				Status:      constants.DIPROSES,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mockKonsultasiStorer := storage.NewMockKonsultasiStorer()
			mockKonsultasiStorer.On("GetUserKonsultasi", mock.Anything, mock.Anything).Return(entities.Konsultasi{
				Model: tC.expectedData.Model,
				UserId:      tC.expectedData.UserId,
				PengacaraId: tC.expectedData.PengacaraId,
				Status:      tC.expectedData.Status,
			}, tC.expectedErr)

			mockKonsultasiUsecase := NewKonsultasiUsecase(mockKonsultasiStorer)
			data, err := mockKonsultasiUsecase.GetUserKonsultasi(tC.userIdDTO, tC.konsultasiIdDTO)
			assert.Equal(t, tC.expectedErr, err)
			assert.Equal(t, tC.userIdDTO, data.UserId)
			assert.Equal(t, tC.konsultasiIdDTO, data.ID)
			assert.Equal(t, tC.expectedData.UserId, data.UserId)
			assert.Equal(t, tC.expectedData.PengacaraId, data.PengacaraId)
		})
	}
}

func TestCreateKonsultasi(t *testing.T) {
	testCases := []struct {
		desc            string
		konsultasiDTO entities.Konsultasi
		expectedErr     error
		expectedData    entities.Konsultasi
	}{
		{
			desc:            "success",
			konsultasiDTO: entities.Konsultasi{
				Model: gorm.Model{
					ID: 1,
				},
				UserId:      1,
				PengacaraId: 1,
				Status:      constants.DIPROSES,
			},
			expectedErr:     nil,
			expectedData: entities.Konsultasi{
				Model: gorm.Model{
					ID: 1,
				},
				UserId:      1,
				PengacaraId: 1,
				Status:      constants.DIPROSES,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mockKonsultasiStorer := storage.NewMockKonsultasiStorer()
			mockKonsultasiStorer.On("CreateKonsultasi", mock.Anything).Return(entities.Konsultasi{
				Model: tC.expectedData.Model,
				UserId:      tC.expectedData.UserId,
				PengacaraId: tC.expectedData.PengacaraId,
				Status:      tC.expectedData.Status,
			}, tC.expectedErr)

			mockKonsultasiUsecase := NewKonsultasiUsecase(mockKonsultasiStorer)
			data, err := mockKonsultasiUsecase.CreateKonsultasi(tC.konsultasiDTO)
			assert.Equal(t, tC.expectedErr, err)
			assert.Equal(t, tC.konsultasiDTO.UserId, data.UserId)
			assert.Equal(t, tC.konsultasiDTO.ID, data.ID)
			assert.Equal(t, tC.konsultasiDTO.PengacaraId, data.PengacaraId)
			assert.Equal(t, tC.expectedData.UserId, data.UserId)
			assert.Equal(t, tC.expectedData.PengacaraId, data.PengacaraId)
		})
	}
}

func TestEditKonsultasi(t *testing.T) {
	testCases := []struct {
		desc            string
		userIdDTO uint
		konsultasiIdDTO uint
		konsultasiDTO entities.Konsultasi
		expectedErr     error
		expectedData    entities.Konsultasi
	}{
		{
			desc:            "success",
			userIdDTO: 1,
			konsultasiIdDTO: 1, 
			konsultasiDTO: entities.Konsultasi{
				Model: gorm.Model{
					ID: 1,
				},
				UserId:      1,
				PengacaraId: 1,
				Status:      constants.DIPROSES,
			},
			expectedErr:     nil,
			expectedData: entities.Konsultasi{
				Model: gorm.Model{
					ID: 1,
				},
				UserId:      1,
				PengacaraId: 1,
				Status:      constants.DIPROSES,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mockKonsultasiStorer := storage.NewMockKonsultasiStorer()
			mockKonsultasiStorer.On("EditKonsultasi", mock.Anything, mock.Anything, mock.Anything).Return(entities.Konsultasi{
				Model: tC.expectedData.Model,
				UserId:      tC.expectedData.UserId,
				PengacaraId: tC.expectedData.PengacaraId,
				Status:      tC.expectedData.Status,
			}, tC.expectedErr)

			mockKonsultasiUsecase := NewKonsultasiUsecase(mockKonsultasiStorer)
			data, err := mockKonsultasiUsecase.EditKonsultasi(tC.konsultasiIdDTO, tC.userIdDTO, tC.konsultasiDTO)
			assert.Equal(t, tC.expectedErr, err)
			assert.Equal(t, tC.konsultasiDTO.UserId, data.UserId)
			assert.Equal(t, tC.konsultasiDTO.ID, data.ID)
			assert.Equal(t, tC.konsultasiDTO.PengacaraId, data.PengacaraId)
			assert.Equal(t, tC.expectedData.UserId, data.UserId)
			assert.Equal(t, tC.expectedData.PengacaraId, data.PengacaraId)
		})
	}
}

func TestDeleteKonsultasi(t *testing.T) {
	testCases := []struct {
		desc            string
		userIdDTO       uint
		konsultasiIdDTO uint
		expectedErr     error
		expectedData    entities.Konsultasi
	}{
		{
			desc:            "success",
			userIdDTO:       1,
			konsultasiIdDTO: 1,
			expectedErr:     nil,
			expectedData: entities.Konsultasi{
				Model: gorm.Model{
					ID: 1,
				},
				UserId:      1,
				PengacaraId: 1,
				Status:      constants.DIPROSES,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mockKonsultasiStorer := storage.NewMockKonsultasiStorer()
			mockKonsultasiStorer.On("DeleteKonsultasi", mock.Anything, mock.Anything).Return(entities.Konsultasi{
				Model: tC.expectedData.Model,
				UserId:      tC.expectedData.UserId,
				PengacaraId: tC.expectedData.PengacaraId,
				Status:      tC.expectedData.Status,
			}, tC.expectedErr)

			mockKonsultasiUsecase := NewKonsultasiUsecase(mockKonsultasiStorer)
			data, err := mockKonsultasiUsecase.DeleteKonsultasi(tC.konsultasiIdDTO, tC.userIdDTO)
			assert.Equal(t, tC.expectedErr, err)
			assert.Equal(t, tC.userIdDTO, data.UserId)
			assert.Equal(t, tC.konsultasiIdDTO, data.ID)
			assert.Equal(t, tC.expectedData.UserId, data.UserId)
			assert.Equal(t, tC.expectedData.PengacaraId, data.PengacaraId)
		})
	}
}