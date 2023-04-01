package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Create router with default middleware
	router := gin.Default()

	router.Run() // Start the HTTP server
}
