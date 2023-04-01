package main

import (
	"github.com/TheMrPuffin/linkdex/configs"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create router with default middleware
	router := gin.Default()

	configs.ConnectDb()

	router.GET("/ping", ping)

	router.Run() // Start the HTTP server
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
