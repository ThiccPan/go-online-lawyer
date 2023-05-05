package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/thiccpan/go-online-lawyer/entities"
	"github.com/thiccpan/go-online-lawyer/usecases"
)

type Konsultasi interface {
	GetAll() error
	GetKonsultasiByUserId(c echo.Context) error
	TestCreateKonsultasi(c echo.Context) error
}

type konsultasi struct {
	useCase usecases.Konsultasi
}

func NewKonsultasiController(konsultasiUseCase usecases.Konsultasi) *konsultasi {
	return &konsultasi{
		konsultasiUseCase,
	}
}

func (k *konsultasi) TestGetKonsultasiByUserId(c echo.Context) error {
	data, err := k.useCase.GetKonsultasiByUserId(1)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"err": err,
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": data,
	})
}

func (k *konsultasi) GetKonsultasiByUserId(c echo.Context) error {
	idPayload, err := strconv.Atoi(c.Param("id"))
	fmt.Println(idPayload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"err": err,
		})
	}

	data, err := k.useCase.GetKonsultasiByUserId(uint(idPayload))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"err": err,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": data,
	})
}

func (k *konsultasi) TestCreateKonsultasi(c echo.Context) error {
	data, err := k.useCase.CreateKonsultasi(1,1,time.Now())
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"err": err,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": data,
	})
}

func (k *konsultasi) CreateKonsultasi(c echo.Context) error {
	payload := entities.KonsultasiDTO{}
	payloadUserId, err := strconv.Atoi(c.Param("id"))
	payload.UserId = uint(payloadUserId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"err": err,
		})
	}
	
	err = c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"err": err.Error(),
		})
	}

	timeFormatted, err := time.Parse(time.RFC3339, payload.KonsultasiTime)
	fmt.Println(payload.KonsultasiTime)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"err": err.Error(),
		})
	}

	
	data, err := k.useCase.CreateKonsultasi(payload.UserId,payload.PengacaraId,timeFormatted)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"err": err,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": data,
	})
}

func (k *konsultasi) EditKonsultasi(c echo.Context) error {
	return nil
}

func (k *konsultasi) DeleteKonsultasi(c echo.Context) error {
	return nil
}