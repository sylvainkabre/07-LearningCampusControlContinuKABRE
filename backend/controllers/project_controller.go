package controllers

import (
	"LearningCampusControlContinu/config"
	"LearningCampusControlContinu/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProjects(c *gin.Context) {

	var projects []models.Project

	if err := config.DB.Find(&projects).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de récupérer au moins un projet"})
		return
	}

	c.JSON(http.StatusOK, projects)
}

func GetSpecificProject(c *gin.Context) {
	var project models.Project

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id invalide"})
	}

	if err := config.DB.First(&project, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Projet introuvable"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de récupérer le projet"})

	}
	c.JSON(http.StatusOK, project)

}

func PostProject(c *gin.Context) {
	var project models.Project

	if err := c.ShouldBindBodyWithJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"})
		return
	}

	if err := config.DB.Create(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la ccréation du projet"})
		return
	}

	c.JSON(http.StatusOK, project)
}
