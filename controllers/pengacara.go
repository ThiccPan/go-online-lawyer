package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/thiccpan/go-online-lawyer/usecases"
)

type Pengacara interface {
	GetAll() error
}

type pengacara struct {
	useCase usecases.Pengacara
}

func NewController(pengacaraUseCase usecases.Pengacara) *pengacara {
	return &pengacara{
		pengacaraUseCase,
	}
}

func (p *pengacara) GetAll(c echo.Context) error {
	data, err := p.useCase.GetAll()
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"data": data,
	})
}