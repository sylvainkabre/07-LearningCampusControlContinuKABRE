package main

import (
	"LearningCampusControlContinu/config"
	"LearningCampusControlContinu/models"
	"LearningCampusControlContinu/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	routes.ProjectRoutes(router)
	routes.UserRoutes(router)
	routes.CommentRoutes(router)
	router.Use(config.SecurityMiddleware())
	router.Use(config.CoresMiddleware())
	router.Use(config.RateLimitMiddleware(100))

	config.ConnectDB()
	fmt.Println("Serveur démarré sur hhtp://localhost:8080")

	// Migration des tables
	config.DB.AutoMigrate(&models.Project{}, &models.User{}, &models.Comment{})

	router.Run(":8000")
	//http.ListenAndServe(":8080", nil)
}
