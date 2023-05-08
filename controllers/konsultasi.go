package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/thiccpan/go-online-lawyer/constants"
	"github.com/thiccpan/go-online-lawyer/entities"
	"github.com/thiccpan/go-online-lawyer/exceptions"
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
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	idPayload, ok := claims["id"].(float64) //if error again try convert to string first
	if !ok {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"err": exceptions.ErrInvalidCredentials,
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

func (k *konsultasi) CreateKonsultasi(c echo.Context) error {
	payload := entities.KonsultasiDTO{}
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	idPayload, ok := claims["id"].(float64) //if error again try convert to string first
	if !ok {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"err": exceptions.ErrInvalidCredentials,
		})
	}
	payload.UserId = uint(idPayload)

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"err": err.Error(),
		})
	}

	timeFormatted, err := time.Parse(time.RFC3339, payload.KonsultasiTime)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"err": err.Error(),
		})
	}
	
	konsultasiData := entities.Konsultasi{
		UserId: payload.UserId,
		PengacaraId: payload.PengacaraId,
		Status: constants.DIPROSES,
		KonsultasiTime: timeFormatted,
	}
	
	err = c.Validate(konsultasiData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"err": err.Error(),
		})
	}

	data, err := k.useCase.CreateKonsultasi(konsultasiData)
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
	payloadKonsultasiId, err := strconv.Atoi(c.Param("konsultasiId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err,
		})
	}
	
	payload := entities.EditKonsultasiDTO{}
	c.Bind(&payload)
	if err := c.Validate(payload); err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{
			"error": err,
		})
	}

	timeFormatted, err := time.Parse(time.RFC3339, payload.KonsultasiTime)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"err": err.Error(),
		})
	}

	konsultasiData := entities.Konsultasi{
		KonsultasiTime: timeFormatted,
	}

	err = c.Validate(konsultasiData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"err": err.Error(),
		})
	}

	data, err := k.useCase.EditKonsultasi(uint(payloadKonsultasiId), konsultasiData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"err": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
			"data": data,
		})
}

func (k *konsultasi) DeleteKonsultasi(c echo.Context) error {
	// userId, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, echo.Map{
	// 		"err": err.Error(),
	// 	})
	// }

	konsultasiId, err := strconv.Atoi(c.Param("konsultasiId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"err": err.Error(),
		})
	}

	data, err := k.useCase.DeleteKonsultasi(uint(konsultasiId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"err": err.Error(),
		})
	}
	
	return c.JSON(http.StatusOK, echo.Map{
		"data": data,
	})
}