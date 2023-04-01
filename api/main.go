package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create router with default middleware
	router := gin.Default()

	router.GET("/ping", ping)

	router.Run() // Start the HTTP server
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
