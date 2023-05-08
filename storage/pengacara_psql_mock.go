package storage

import (
	"github.com/stretchr/testify/mock"
	"github.com/thiccpan/go-online-lawyer/entities"
)

type mockPengacaraStorer struct {
	mock.Mock
}

func NewMockPengacaraStorer() *mockPengacaraStorer {
	return &mockPengacaraStorer{}
}

func (mp *mockPengacaraStorer) GetAll() ([]entities.Pengacara, error) {
	ret := mp.Called()
	return ret.Get(0).([]entities.Pengacara), ret.Error(1)
}

func (mp *mockPengacaraStorer) GetById(id int) (entities.Pengacara, error) {
	ret := mp.Called(id)
	return ret.Get(0).(entities.Pengacara), ret.Error(1)
}

func (mp *mockPengacaraStorer) GetByEmail(email string) (entities.Pengacara, error) {
	ret := mp.Called(email)
	return ret.Get(0).(entities.Pengacara), ret.Error(1)
}

func (mp *mockPengacaraStorer) GetByCategory(category string) ([]entities.Pengacara, error) {
	ret := mp.Called()
	return ret.Get(0).([]entities.Pengacara), ret.Error(1)
}

