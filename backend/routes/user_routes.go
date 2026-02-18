package routes

import (
	"LearningCampusControlContinu/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	routerGroup := router.Group("/users")
	{
		routerGroup.POST("/register", controllers.Register)
	}
}
