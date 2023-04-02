package routes

import (
	"github.com/TheMrPuffin/linkdex/controllers"
	"github.com/gin-gonic/gin"
)

func LinkRoute(router *gin.Engine) {
	router.GET("/ping", controllers.Ping)
	router.GET("/links", controllers.GetAllLinks())
	router.GET("/link/:linkName", controllers.GetLinkByName())
	router.POST("/link", controllers.CreateLink())
	router.DELETE("/link/:linkName", controllers.DeleteLink())
}
