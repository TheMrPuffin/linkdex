package main

import (
	"github.com/TheMrPuffin/linkdex/configs"
	"github.com/TheMrPuffin/linkdex/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// Create router with default middleware
	router := gin.Default()

	// COR configuration for GIN router
	// Currently allows all origins
	router.Use(cors.Default()) // This needs to be before routes declaration

	// Run database
	configs.ConnectDb()

	//routes
	routes.LinkRoute(router)

	router.Run() // Start the HTTP server
}
