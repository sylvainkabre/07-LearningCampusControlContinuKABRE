package routes

import (
	"LearningCampusControlContinu/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	routesGroup := router.Group("/users")
	{
		routesGroup.POST("/register", controllers.Register)
		routesGroup.POST("/login", controllers.Login)
		routesGroup.GET("/all", controllers.GetAllUsers)
	}
}
