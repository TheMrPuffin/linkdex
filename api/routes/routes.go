package routes

import (
	"github.com/TheMrPuffin/linkdex/controllers"
	"github.com/gin-gonic/gin"
)

func LinkRoute(router *gin.Engine) {
	router.GET("/ping", controllers.Ping)
}
