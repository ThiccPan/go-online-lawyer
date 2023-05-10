package main

import (
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	gormPsqlConf "github.com/thiccpan/go-online-lawyer/app/config"
	"github.com/thiccpan/go-online-lawyer/constants"
	"github.com/thiccpan/go-online-lawyer/controllers"
	"github.com/thiccpan/go-online-lawyer/helper"
	"github.com/thiccpan/go-online-lawyer/storage"
	"github.com/thiccpan/go-online-lawyer/usecases"
	"github.com/thiccpan/go-online-lawyer/validations"
)

func main() {
	e := echo.New()

	// DBconfDefault := gormPsqlConf.ConfigDB{
	// 	DB_Username: "thiccpan",
	// 	DB_Password: "dbpsqlpass432",
	// 	DB_Port:     "5432",
	// 	DB_Host:     "localhost",
	// 	DB_Name:     "go_online_lawyer",
	// }

	godotenv.Load(".env")
	DBconfDefault := gormPsqlConf.ConfigDB{
		DB_Username: os.Getenv("DB_USERNAME"),
		DB_Password: os.Getenv("DB_PASSWORD"),
		DB_Port:     os.Getenv("DB_PORT"),
		DB_Host:     os.Getenv("DB_HOST"),
		DB_Name:     os.Getenv("DB_NAME"),
	}

	// services
	DB := DBconfDefault.InitDB()
	tokenManager := helper.NewAuthJWT()

	// jwtConfig := echojwt.Config{
	// 	NewClaimsFunc: func(c echo.Context) jwt.Claims {
	// 		return new(helper.JwtCustomClaims)
	// 	},
	// 	SigningKey: constants.JWT_SECRET,
	// }

	// storage, usecase, and controller setup
	// pengacara
	pengacaraStorage := storage.NewPengacaraStorer(DB)
	pengacaraUseCase := usecases.NewPengacaraUsecase(pengacaraStorage)
	pengacaraController := controllers.NewPengacaraController(pengacaraUseCase)
	// user
	userStorage := storage.NewUserStorer(DB)
	userUseCase := usecases.NewUserUsecase(userStorage)
	userController := controllers.NewUserController(userUseCase, tokenManager)
	// konsultasi
	konsultasiStorage := storage.NewKonsultasiStorer(DB)
	konsultasiUsecase := usecases.NewKonsultasiUsecase(konsultasiStorage)
	konsultasiController := controllers.NewKonsultasiController(konsultasiUsecase)
	// ratingPengacara
	ratingPengacaraStorage := storage.NewRatingPengacaraStorer(DB)
	ratingPengacaraUsecase := usecases.NewRatingPengacaraUsecase(ratingPengacaraStorage)
	ratingPengacaraController := controllers.NewRatingPengacaraController(ratingPengacaraUsecase)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = &validations.CustomValidator{Validator: validator.New()}

	// Route
	e.GET("/health", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusAccepted, "online")
	})

	userJWT := e.Group("/user")
	userJWT.Use(echojwt.JWT([]byte(constants.JWT_SECRET)))
	// pengacara route
	e.GET("/pengacaras", pengacaraController.GetAll)
	e.GET("/pengacaras/:id", pengacaraController.GetById)
	e.GET("/pengacaras/filter", pengacaraController.GetWithFilter)
	e.GET("/pengacaras/category/:category", pengacaraController.GetByCategory)
	// user route
	e.POST("/register", userController.UserRegister)
	e.POST("/login", userController.UserLogin)
	// konsultasi route
	userJWT.GET("/konsultasi", konsultasiController.GetAllUserKonsultasi)
	userJWT.POST("/konsultasi", konsultasiController.CreateUserKonsultasi)
	userJWT.GET("/konsultasi/:konsultasiId", konsultasiController.GetUserKonsultasi)
	userJWT.PUT("/konsultasi/:konsultasiId", konsultasiController.EditUserKonsultasi)
	userJWT.DELETE("/konsultasi/:konsultasiId", konsultasiController.DeleteUserKonsultasi)
	// edit konsultasi status and/or link
	e.PUT("/konsultasi/:konsultasiId", konsultasiController.EditPengacaraKonsultasi)
	// rating route
	e.GET("/pengacaras/rating/:pengacaraId", ratingPengacaraController.GetAllRatingByPengacara)
	userJWT.GET("/rating", ratingPengacaraController.GetAllRatingByUser)
	userJWT.POST("/rating", ratingPengacaraController.CreateRating)
	userJWT.DELETE("/rating/:ratingId", ratingPengacaraController.DeleteRating)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
