package usecases

import (
	"github.com/go-playground/validator/v10"
	"github.com/thiccpan/go-online-lawyer/entities"
	"github.com/thiccpan/go-online-lawyer/exceptions"
	"github.com/thiccpan/go-online-lawyer/storage"
)

type User interface {
	Register(payload entities.UserDTO) (entities.User, error)
	Login(payload entities.UserLoginDTO) (entities.User, error)
}

type user struct {
	UserStorer storage.UserStorer
}

func NewUserUsecase(userStorage storage.UserStorer) *user {
	return &user{
		UserStorer: userStorage,
	}
}

func (u *user) Register(payload entities.UserDTO) (entities.User, error) {
	validate := validator.New()
	if err := validate.Struct(payload); err != nil {
		err = exceptions.ErrValidationFailed
		return entities.User{}, err
	}
	
	userData := entities.User{}
	userData.Username = payload.Username
	userData.Email = payload.Email
	userData.Password = payload.Password

	return u.UserStorer.Add(userData)
}

func (u *user) Login(payload entities.UserLoginDTO) (entities.User, error) {
	validate := validator.New()
	if err := validate.Struct(payload); err != nil {
		err = exceptions.ErrInvalidCredentials
		return entities.User{}, err
	}

	data, err := u.UserStorer.GetByEmail(payload.Email)
	if err != nil {
		return entities.User{}, err
	}
	
	if data.Password != payload.Password {
		err = exceptions.ErrInvalidCredentials
		return entities.User{}, err
	}
	
	return data, err
}