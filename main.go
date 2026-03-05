package main

import (
	"aplikasi-internal/config"
	"aplikasi-internal/routes"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

// @title Aplikasi Internal API
// @version 1.0
// @description API untuk aplikasi internal
// @host localhost:8082
// @BasePath /


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal load .env")
	}
	
	config.ConnectDB()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		MaxAge: 12 * time.Hour,
	}))
	
	routes.SetupRoutes(r)

	r.Run(":" + os.Getenv("APP_PORT"))
}
