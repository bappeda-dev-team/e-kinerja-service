package main

import (
	"aplikasi-internal/config"
	"aplikasi-internal/routes"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal load .env")
	}
	
	config.ConnectDB()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":" + os.Getenv("APP_PORT"))
}
