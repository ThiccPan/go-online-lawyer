package storage

import (
	"github.com/stretchr/testify/mock"
	"github.com/thiccpan/go-online-lawyer/entities"
)

type mockUserStorer struct {
	mock.Mock
}

func NewMockUserStorer() *mockUserStorer {
	return &mockUserStorer{}
}

func (m *mockUserStorer) Add(data entities.User) (entities.User, error) {
	testObj := m.Called(data)
	return testObj.Get(0).(entities.User), testObj.Error(1)
}

func (m *mockUserStorer) GetByEmail(email string) (entities.User, error) {
	testObj := m.Called(email)	
	return testObj.Get(0).(entities.User), testObj.Error(1)
}

func (m *mockUserStorer) GetAllKonsultasi(user *entities.User) ([]entities.Konsultasi, error) {
	return []entities.Konsultasi{}, nil
}