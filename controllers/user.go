package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/thiccpan/go-online-lawyer/entities"
	"github.com/thiccpan/go-online-lawyer/exceptions"
	"github.com/thiccpan/go-online-lawyer/helper"
	"github.com/thiccpan/go-online-lawyer/usecases"
)

type User interface {
	Register() error
	Login() error
}

type user struct {
	useCase usecases.User
	tokenManager helper.AuthJWT
}

func NewUserController(UserUseCase usecases.User, tokenManager helper.AuthJWT) *user {
	return &user{
		UserUseCase,
		tokenManager,
	}
}

func (u *user) UserRegister(c echo.Context) error {
	userDTO := entities.UserDTO{}
	if err := c.Bind(&userDTO); err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	if err := c.Validate(userDTO); err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}

	data, err := u.useCase.Register(userDTO)
	if err == exceptions.ErrUserAlreadyExist {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}
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
	userLoginDTO := entities.UserLoginDTO{}
	err := c.Bind(&userLoginDTO)
	if err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}

	data, err := u.useCase.Login(userLoginDTO)
	if err != nil {	
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	token, err := u.tokenManager.GenerateToken(data.Email)
	if err != nil {	
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	
	return c.JSON(200, echo.Map{
		"data": data,
		"token": token,
	})
}