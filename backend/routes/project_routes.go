package routes

import (
	"LearningCampusControlContinu/controllers"

	"github.com/gin-gonic/gin"
)

func ProjectRoutes(router *gin.Engine) {
	routesGroup := router.Group("/projects")

	{
		routesGroup.GET("/", controllers.GetProjects)
		routesGroup.GET("/:id", controllers.GetSpecificProject)
		routesGroup.POST("/", controllers.PostProject)
		routesGroup.PUT("/:id", controllers.PutProject)
		routesGroup.DELETE("/:id", controllers.DeleteProject)
	}
}
