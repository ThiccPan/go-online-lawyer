package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/thiccpan/go-online-lawyer/entities"
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

func (p *pengacara) Add(c echo.Context) error {
	newPengacaraReq := entities.PengacaraDTO{}

	err := c.Bind(&newPengacaraReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "wrong field value",
		})
	}

	newPengacaraData := entities.Pengacara{
		Username: newPengacaraReq.Username,
		Email: newPengacaraReq.Email,
		Password: newPengacaraReq.Password,
		Category: newPengacaraReq.Category,
		Price: newPengacaraReq.Price,
	}

	newP, err := p.useCase.Add(newPengacaraData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "something's wrong",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "new lawyer successfully added",
		"data":    newP,
	})
}
