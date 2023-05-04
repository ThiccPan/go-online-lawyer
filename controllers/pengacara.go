package controllers

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/thiccpan/go-online-lawyer/usecases"
)

type Pengacara interface {
	GetAll() error
}

type pengacara struct {
	useCase usecases.Pengacara
}

func NewPengacaraController(pengacaraUseCase usecases.Pengacara) *pengacara {
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

func (p *pengacara) GetById(c echo.Context) error {
	idPayload, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}

	data, err := p.useCase.GetById(idPayload)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(200, echo.Map{
		"data": data,
	})
}

func (p *pengacara) GetWithFilter(c echo.Context) error {
	emailPayload := c.QueryParam("email")
	data, err := p.useCase.GetByEmail(emailPayload)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"data": data,
	})
}

func (p *pengacara) GetByCategory(c echo.Context) error {
	categoryPayload := c.Param("category")
	data, err := p.useCase.GetByCategory(categoryPayload)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"data": data,
	})
}