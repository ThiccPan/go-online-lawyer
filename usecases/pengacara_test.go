package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/thiccpan/go-online-lawyer/entities"
	"github.com/thiccpan/go-online-lawyer/storage"
	"gorm.io/gorm"
)

func TestGetAll(t *testing.T) {
	testCases := []struct {
		desc           string
		expectedErr    error
		expectedData []entities.Pengacara
	}{
		{
			desc: "success",
			expectedErr: nil,
			expectedData: []entities.Pengacara{
				{
					Username: "test",
					Email: "test@gmail.com",
					Password: "test",
					Category: "test",
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mockPengacaraStorer := storage.NewMockPengacaraStorer()
			mockPengacaraStorer.On("GetAll", mock.Anything).Return([]entities.Pengacara{
				{
					Username: tC.expectedData[0].Username,
					Email: tC.expectedData[0].Email,
					Password: tC.expectedData[0].Password,
					Category: tC.expectedData[0].Category,
				},
			}, tC.expectedErr)
			pengacaraUsecase := NewPengacaraUsecase(mockPengacaraStorer)
			data, err := pengacaraUsecase.PengacaraStorer.GetAll()
			assert.Equal(t, tC.expectedErr, err)
			assert.Equal(t, tC.expectedData[0].Email, data[0].Email)
		})
	}
}

func TestGetById(t *testing.T) {
	testCases := []struct {
		desc	string
		idDTO 	uint
		expectedErr error
		expectedData entities.Pengacara
	}{
		{
			desc: "success",
			idDTO: 1,
			expectedErr: nil,
			expectedData: entities.Pengacara{
				Model: gorm.Model{
					ID: 1,
				},
				Username: "test",
				Email: "test@gmail.com",
				Password: "test",
				Category: "test",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mockPengacaraStorer := storage.NewMockPengacaraStorer()
			mockPengacaraStorer.On("GetById", mock.Anything).Return(entities.Pengacara{
				Model: gorm.Model{
					ID: tC.expectedData.ID,
				},
				Username: tC.expectedData.Username,
				Email: tC.expectedData.Email,
				Password: tC.expectedData.Password,
				Category: tC.expectedData.Category,
			}, tC.expectedErr)
			pengacaraUsecase := NewPengacaraUsecase(mockPengacaraStorer)
			data, err := pengacaraUsecase.PengacaraStorer.GetById(int(tC.idDTO))
			assert.Equal(t, tC.expectedErr, err)
			assert.Equal(t, tC.expectedData.Email, data.Email)
		})
	}
}

func TestGetByEmail(t *testing.T) {
	testCases := []struct {
		desc	string
		emailDTO 	string
		expectedErr error
		expectedData entities.Pengacara
	}{
		{
			desc: "success",
			emailDTO: "test@gmail.com",
			expectedErr: nil,
			expectedData: entities.Pengacara{
				Username: "test",
				Email: "test@gmail.com",
				Password: "test",
				Category: "test",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mockPengacaraStorer := storage.NewMockPengacaraStorer()
			mockPengacaraStorer.On("GetByEmail", mock.Anything).Return(entities.Pengacara{
				Model: gorm.Model{
					ID: tC.expectedData.ID,
				},
				Username: tC.expectedData.Username,
				Email: tC.expectedData.Email,
				Password: tC.expectedData.Password,
				Category: tC.expectedData.Category,
			}, tC.expectedErr)
			pengacaraUsecase := NewPengacaraUsecase(mockPengacaraStorer)
			data, err := pengacaraUsecase.PengacaraStorer.GetByEmail(tC.emailDTO)
			assert.Equal(t, tC.expectedErr, err)
			assert.Equal(t, tC.expectedData.Email, data.Email)
		})
	}
}

func TestGetByCategory(t *testing.T) {
	testCases := []struct {
		desc           string
		categoryDTO string
		expectedErr    error
		expectedData []entities.Pengacara
	}{
		{
			desc: "success",
			categoryDTO: "test",
			expectedErr: nil,
			expectedData: []entities.Pengacara{
				{
					Username: "test",
					Email: "test@gmail.com",
					Password: "test",
					Category: "test",
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mockPengacaraStorer := storage.NewMockPengacaraStorer()
			mockPengacaraStorer.On("GetByCategory", mock.Anything).Return([]entities.Pengacara{
				{
					Username: tC.expectedData[0].Username,
					Email: tC.expectedData[0].Email,
					Password: tC.expectedData[0].Password,
					Category: tC.expectedData[0].Category,
				},
			}, tC.expectedErr)
			pengacaraUsecase := NewPengacaraUsecase(mockPengacaraStorer)
			data, err := pengacaraUsecase.PengacaraStorer.GetByCategory(tC.categoryDTO)
			assert.Equal(t, tC.expectedErr, err)
			assert.Equal(t, tC.expectedData[0].Email, data[0].Email)
		})
	}
}