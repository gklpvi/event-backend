/*
Gamegos Case Study for Backend Engineer Position
-Gökalp Akartepe
*/
package main

import (
	"log"
	"os"

	middleware "event-backend/middleware"
	model "event-backend/model"
	route "event-backend/route"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	model.ConnectDataBase()
	// Disconnect from database when the application is closed
	defer model.DisconnectDataBase()

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Setup route groups for the API
	authRoutes := router.Group("/auth")
	{
		route.AuthRouter(authRoutes)
	}

	protected := router.Group("/protected")
	protected.Use(middleware.JwtAuthMiddleware())
	{
		route.UserRouter(protected)
		route.ProfileRouter(protected)
		route.GroupRouter(protected)
		route.GroupMemberRouter(protected)
		route.GroupCategoryRouter(protected)
		route.EventRouter(protected)
	}

	// Start and run the server
	router.Run(":" + os.Getenv("DEV_PORT"))
}
