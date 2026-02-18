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

	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"message": "pong"})
	// })

	routes.ProjectRoutes(router)
	routes.UserRoutes(router)

	config.ConnectDB()
	fmt.Println("Serveur démarré sur hhtp://localhost:8080")

	// Migration des tables
	config.DB.AutoMigrate(&models.Project{}, &models.User{})

	router.Run(":8000")
	//http.ListenAndServe(":8080", nil)
}
