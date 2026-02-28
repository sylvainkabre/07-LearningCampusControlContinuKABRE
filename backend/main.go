package main

import (
	"LearningCampusControlContinu/config"
	"LearningCampusControlContinu/models"
	"LearningCampusControlContinu/routes"
	"fmt"

	_ "LearningCampusControlContinu/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Learning Campus Control Continu API
// @version 1.0
// @description API pour la gestion des projets, utilisateurs et commentaires du Learning Campus Control Continu.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {

	router := gin.Default()

	routes.ProjectRoutes(router)
	routes.UserRoutes(router)
	routes.CommentRoutes(router)
	router.Use(config.SecurityMiddleware())
	router.Use(config.CoresMiddleware())
	router.Use(config.RateLimitMiddleware(100))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	config.ConnectDB()
	fmt.Println("Serveur démarré sur hhtp://localhost:8080")

	// Migration des tables
	config.DB.AutoMigrate(&models.Project{}, &models.User{}, &models.Comment{})

	router.Run(":8000")
}
