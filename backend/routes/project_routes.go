package routes

import (
	"LearningCampusControlContinu/controllers"
	"LearningCampusControlContinu/middlewares"

	"github.com/gin-gonic/gin"
)

func ProjectRoutes(router *gin.Engine) {
	routesGroup := router.Group("/projects")
	routesGroup.Use(middlewares.Authentication())

	{
		routesGroup.GET("/", controllers.GetProjects)
		routesGroup.GET("/:id", controllers.GetSpecificProject)
		routesGroup.POST("/", controllers.PostProject)
		routesGroup.PUT("/:id", controllers.PutProject)
		routesGroup.DELETE("/:id", controllers.DeleteProject)
	}
}
