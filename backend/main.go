package main

import (
	"log"

	"github.com/ericengeng/tasks-app/backend/config"
	"github.com/ericengeng/tasks-app/backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	client := config.ConnectDB()

	// Ensure disconnection when application closes
	defer client.Disconnect(config.Ctx)

	// Initialize the Gin router
	router := gin.Default()

	router.Use(cors.Default())

	// Setup routes
	routes.RegisterRoutes(router)

	// Start the server
	log.Println("Server running on http://localhost:8080")
	router.Run(":8080")
}
