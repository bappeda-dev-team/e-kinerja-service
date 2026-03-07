package main

import (
	"aplikasi-internal/config"
	"aplikasi-internal/internal/helpers"
	"aplikasi-internal/routes"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @title Aplikasi Internal API
// @version 1.0
// @description API untuk aplikasi internal
// @host localhost:8082
// @BasePath /

func main() {
	godotenv.Load()

	config.ConnectDB()

	e := echo.New()

	e.Validator = &helpers.CustomValidator{
		Validator: validator.New(),
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			echo.GET,
			echo.POST,
			echo.PUT,
			echo.DELETE,
			echo.OPTIONS,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAuthorization,
		},
	}))

	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
