package storage

import (
	"github.com/stretchr/testify/mock"
	"github.com/thiccpan/go-online-lawyer/entities"
)

type mockKonsultasiStorer struct {
	mock.Mock
}

func NewMockKonsultasiStorer() *mockKonsultasiStorer {
	return &mockKonsultasiStorer{}
}

func (mk *mockKonsultasiStorer) GetAllKonsultasi() ([]entities.Konsultasi, error) {
	ret := mk.Called()
	return ret.Get(0).([]entities.Konsultasi), ret.Error(1)
}

func (mk *mockKonsultasiStorer) GetAllUserKonsultasi(userId uint) ([]entities.Konsultasi, error) {
	ret := mk.Called()
	return ret.Get(0).([]entities.Konsultasi), ret.Error(1)
}

func (mk *mockKonsultasiStorer) GetUserKonsultasi(userId uint, konsultasiId uint) (entities.Konsultasi, error) {
	ret := mk.Called(userId, konsultasiId)
	return ret.Get(0).(entities.Konsultasi), ret.Error(1)
}

func (mk *mockKonsultasiStorer) CreateKonsultasi(konsultasiData entities.Konsultasi) (entities.Konsultasi, error) {
	ret := mk.Called(konsultasiData)
	return ret.Get(0).(entities.Konsultasi), ret.Error(1)
}

func (mk *mockKonsultasiStorer) EditKonsultasi(konsultasiId uint, userId uint, data entities.Konsultasi) (entities.Konsultasi, error) {
	ret := mk.Called(konsultasiId, userId, data)
	return ret.Get(0).(entities.Konsultasi), ret.Error(1)
}

func (k *mockKonsultasiStorer) EditKonsultasiPengacara(konsultasiId uint, konsultasiData entities.Konsultasi) (entities.Konsultasi, error) {
	return entities.Konsultasi{}, nil
}

func (mk *mockKonsultasiStorer) DeleteKonsultasi(konsultasiId uint, userId uint) (entities.Konsultasi, error) {
	ret := mk.Called(konsultasiId, userId)
	return ret.Get(0).(entities.Konsultasi), ret.Error(1)
}