package main

import (
	"aplikasi-internal/config"
	"aplikasi-internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8082")
}
