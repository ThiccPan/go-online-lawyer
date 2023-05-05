package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	gormPsqlConf "github.com/thiccpan/go-online-lawyer/app/config"
	"github.com/thiccpan/go-online-lawyer/controllers"
	"github.com/thiccpan/go-online-lawyer/helper"
	"github.com/thiccpan/go-online-lawyer/storage"
	"github.com/thiccpan/go-online-lawyer/usecases"
	"github.com/thiccpan/go-online-lawyer/validations"
)

func main() {
	e := echo.New()

	DBconf := gormPsqlConf.ConfigDB{
		DB_Username: "thiccpan",
		DB_Password: "dbpsqlpass432",
		DB_Port:     "5432",
		DB_Host:     "localhost",
		DB_Name:     "go_online_lawyer",
	}
	
	// services
	DB := DBconf.InitDB()
	tokenManager := helper.NewAuthJWT()

	// storage, usecase, and controller setup
	// pengacara
	pengacaraStorage := storage.NewPengacaraStorer(DB)
	pengacaraUseCase := usecases.NewPengacaraUsecase(pengacaraStorage)
	pengacaraController := controllers.NewPengacaraController(pengacaraUseCase)
	// user
	userStorage := storage.NewUserStorer(DB)
	userUseCase:= usecases.NewUserUsecase(userStorage)
	userController := controllers.NewUserController(userUseCase, tokenManager)
	// konsultasi
	konsultasiStorage := storage.NewKonsultasiStorer(DB)
	konsultasiUsecase := usecases.NewKonsultasiUsecase(konsultasiStorage)
	konsultasiController := controllers.NewKonsultasiController(konsultasiUsecase)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = &validations.CustomValidator{Validator: validator.New()}

	// Route
	e.GET("/health", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusAccepted, "online")
	})
	// pengacara route
	e.GET("/pengacaras", pengacaraController.GetAll)
	e.GET("/pengacaras/:id", pengacaraController.GetById)
	e.GET("/pengacaras/filter", pengacaraController.GetWithFilter)
	e.GET("/pengacaras/category/:category", pengacaraController.GetByCategory)
	// user route
	e.POST("/register", userController.UserRegister)
	e.POST("/login", userController.UserLogin)
	// konsultasi route
	e.GET("/:id/konsultasi", konsultasiController.GetKonsultasiByUserId)
	e.POST("/:id/konsultasi", konsultasiController.CreateKonsultasi)
	e.PUT("/:id/konsultasi", konsultasiController.EditKonsultasi)
	e.DELETE("/:id/konsultasi", konsultasiController.DeleteKonsultasi)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
