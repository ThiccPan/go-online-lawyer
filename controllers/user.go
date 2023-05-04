package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/thiccpan/go-online-lawyer/entities"
	"github.com/thiccpan/go-online-lawyer/usecases"
)

type User interface {
	Register() error
	Login() error
}

type user struct {
	useCase usecases.User
}

func NewUserController(UserUseCase usecases.User) *user {
	return &user{
		UserUseCase,
	}
}

func (u *user) UserRegister(c echo.Context) error {
	userDTO := entities.UserDTO{}
	err := c.Bind(&userDTO)
	if err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}

	data, err := u.useCase.Register(userDTO)
	if err != nil {	
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(200, echo.Map{
		"data": data,
	})
}

func (u *user) UserLogin(c echo.Context) error {
	userDTO := entities.UserDTO{}
	err := c.Bind(&userDTO)
	if err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}

	data, err := u.useCase.Login(userDTO)
	if err != nil {	
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	
	return c.JSON(200, echo.Map{
		"data": data,
	})
}