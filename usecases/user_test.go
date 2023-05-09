package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/thiccpan/go-online-lawyer/entities"
	"github.com/thiccpan/go-online-lawyer/storage"
)

func TestRegister(t *testing.T) {
	testCases := []struct {
		desc         string
		userPayload  entities.UserDTO
		expectedErr  error
		expectedData entities.User
	}{
		{
			desc:        "success",
			userPayload: entities.UserDTO{
				Username: "tes",
				Email: "test@email.com",
				Password: "123",
			},
			expectedErr: nil,
			expectedData: entities.User{
				Username: "tes",
				Email: "test@email.com",
				Password: "123",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mockUserStorer := storage.NewMockUserStorer()
			mockUserStorer.On("GetByEmail", mock.Anything).Return(entities.User{}, tC.expectedErr)
			mockUserStorer.On("Add", mock.Anything).Return(entities.User{
				Username: tC.expectedData.Username,
				Email: tC.expectedData.Email,
				Password: tC.expectedData.Password,
			}, tC.expectedErr)

			userUseCase := NewUserUsecase(mockUserStorer)
			data, err := userUseCase.Register(tC.userPayload)
			assert.Equal(t, tC.expectedErr, err)
			assert.Equal(t, tC.expectedData.Email, data.Email)
		})
	}
}

func TestLogin(t *testing.T) {
	testCases := []struct {
		desc	string
		userPayload  entities.UserLoginDTO
		expectedErr  error
		expectedData entities.User
	}{
		{
			desc: "success",
			userPayload: entities.UserLoginDTO{
				Email: "test@email.com",
				Password: "123",
			},
			expectedErr: nil,
			expectedData: entities.User{
				Email: "test@email.com",
				Password: "123",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mockUserStorer := storage.NewMockUserStorer()
			mockUserStorer.On("GetByEmail", mock.AnythingOfType("string")).Return(entities.User{
				Username: tC.expectedData.Username,
				Email: tC.expectedData.Email,
				Password: tC.expectedData.Password,
			}, tC.expectedErr)
			userUsecase := NewUserUsecase(mockUserStorer)
			data, err := userUsecase.Login(tC.userPayload)
			assert.Equal(t, tC.expectedErr, err)
			assert.Equal(t, tC.expectedData.Email, data.Email)
		})
	}
}
