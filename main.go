package main

import (
	"log"
	"os"

	middleware "event-backend/middleware"
	route "event-backend/route"
	model "event-backend/model"

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
	auth_routes := router.Group("/auth")
	{
		route.LoginRouter(auth_routes)
		route.RegisterRouter(auth_routes)
	}

	protected := router.Group("/protected")
	protected.Use(middleware.BlacklistTokenMiddleware())
	protected.Use(middleware.JwtAuthMiddleware())
	{
		route.UserRouter(protected)
		route.ProfileRouter(protected)
		route.GroupRouter(protected)
		route.GroupMemberRouter(protected)
		route.GroupCategoryRouter(protected)
	}

	// Start and run the server
	router.Run(":" + os.Getenv("DEV_PORT"))
}
