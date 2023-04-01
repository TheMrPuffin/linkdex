package main

import (
	"github.com/TheMrPuffin/linkdex/configs"
	"github.com/TheMrPuffin/linkdex/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create router with default middleware
	router := gin.Default()

	// Run database
	configs.ConnectDb()

	//routes
	routes.LinkRoute(router)

	router.Run() // Start the HTTP server
}
