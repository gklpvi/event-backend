package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Setup route group for the API
	// api := router.Group("/api")
	// {
		
	// }

	// Start and run the server
	router.Run(":" + os.Getenv("PORT"))
}
