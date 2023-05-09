package controllers

import (
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/thiccpan/go-online-lawyer/entities"
	"github.com/thiccpan/go-online-lawyer/exceptions"
	"github.com/thiccpan/go-online-lawyer/usecases"
)

type RatingPengacara interface {}

type ratingPengacara struct {
	usecase usecases.RatingPengacara
}

func NewRatingPengacaraController(usecase usecases.RatingPengacara) *ratingPengacara {
	return &ratingPengacara{
		usecase: usecase,
	}
}

func (rp *ratingPengacara) CreateRating(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	idPayload, ok := claims["id"].(float64) //if error again try convert to string first
	if !ok {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"err": exceptions.ErrInvalidCredentials,
		}) 
	}
	
	ratingDTO := entities.RatingPengacaraCreateDTO{}
	err := c.Bind(&ratingDTO)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"err": err,
		})
	}
	
	ratingData := entities.RatingPengacara{
		PengacaraId: ratingDTO.PengacaraId,
		Rating: ratingDTO.Rating,
	}

	data, err := rp.usecase.CreateRating(uint(idPayload), ratingData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"err": err,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": data,
		"idpayload": idPayload,
	})
}

func (rp *ratingPengacara) GetAllRatingByUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	idPayload, ok := claims["id"].(float64) //if error again try convert to string first
	if !ok {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"err": exceptions.ErrInvalidCredentials,
		}) 
	}
	
	data, err := rp.usecase.GetAllRatingByUser(uint(idPayload))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"err": exceptions.ErrInvalidCredentials,
		}) 
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": data,
	})
}

func (rp *ratingPengacara) GetAllRatingByPengacara(c echo.Context) error {
	pengacaraId, err := strconv.Atoi(c.Param("pengacaraId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"err": err.Error(),
		}) 
	}

	data, err := rp.usecase.GetAllRatingByPengacara(uint(pengacaraId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"err": exceptions.ErrInvalidCredentials,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": data,
	})
}

func (rp *ratingPengacara) DeleteRating(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	idPayload, ok := claims["id"].(float64) //if error again try convert to string first
	if !ok {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"err": exceptions.ErrInvalidCredentials,
		})
	}

	ratingId, err := strconv.Atoi(c.Param("ratingId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"err": err.Error(),
		})
	}

	data, err := rp.usecase.DeleteRating(uint(idPayload), uint(ratingId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"err": exceptions.ErrInvalidCredentials,
		})	
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": data,
	})	
}