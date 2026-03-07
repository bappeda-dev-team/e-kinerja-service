package main

import (
	"aplikasi-internal/config"
	"aplikasi-internal/internal/helpers"
	"aplikasi-internal/routes"
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	runMigrations()

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

func runMigrations() {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	m, err := migrate.New("file://migrations", dsn)
	if err != nil {
		log.Fatal("Migration init failed:", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Migration failed:", err)
	}

	log.Println("✅ Migration sukses")
}
