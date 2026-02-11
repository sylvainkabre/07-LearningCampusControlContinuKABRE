package controllers

import (
	"LearningCampusControlContinu/config"
	"LearningCampusControlContinu/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProjects(c *gin.Context) {

	var projects []models.Project

	if err := config.DB.Find(&projects).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de récupérer au moins un projet"})

		return
	}

	c.JSON(http.StatusOK, projects)
}
