package routes

import (
	"LearningCampusControlContinu/controllers"
	"LearningCampusControlContinu/middlewares"

	"github.com/gin-gonic/gin"
)

func CommentRoutes(router *gin.Engine) {
	routesGroup := router.Group("/comments")
	routesGroup.Use(middlewares.Authentication())

	{
		routesGroup.POST("/", controllers.PostComment)
	}
}
