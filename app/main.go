package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/healthchek", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusAccepted, "online")
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}