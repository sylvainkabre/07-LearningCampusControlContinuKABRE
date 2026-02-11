package routes

import (
	"LearningCampusControlContinu/controllers"

	"github.com/gin-gonic/gin"
)

func ProjectRoutes(router *gin.Engine) {
	routesGroup := router.Group("/projects")

	{
		routesGroup.GET("/", controllers.GetProjects)
	}
}
