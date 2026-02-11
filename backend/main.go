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

	config.ConnectDB()
	fmt.Println("Serveur démarré sur hhtp://localhost:8080")

	// Migration des tables
	config.DB.AutoMigrate(&models.Project{})

	router.Run(":8000")
	//http.ListenAndServe(":8080", nil)
}
