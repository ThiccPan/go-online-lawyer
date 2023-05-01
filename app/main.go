package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	gormPsqlConf "github.com/thiccpan/go-online-lawyer/app/config"
)

func main() {
	e := echo.New()

	DB := gormPsqlConf.ConfigDB{
		DB_Username: "thiccpan",
		DB_Password: "dbpsqlpass432",
		DB_Port:     "5432",
		DB_Host:     "localhost",
		DB_Name:     "go_online_lawyer",
	}

	DB.InitDB()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/health", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusAccepted, "online")
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
