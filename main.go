package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swapnika/jwt-auth/routes"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Unable to load environment variables")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "9000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.TransactionRoutes(router)
	routes.UserRoutes(router)

	router.Run(":" + port)
}
