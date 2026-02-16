package controllers

import (
	"LearningCampusControlContinu/config"
	"LearningCampusControlContinu/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
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

func PutProject(c *gin.Context) {
	var project models.Project

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID non valide"})
	}

	if err := config.DB.First(&project, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Projet non trouvé"})
		return
	}

	// var input map[string]interface{}
	// if err := c.ShouldBindJSON(&input); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Le format de la donnée est invalide"})
	// 	return
	// }

	var input models.ProjectUpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format de la donnée est invalide"})
		return
	}

	updates := make(map[string]interface{})

	if input.Name != nil {
		updates["name"] = *input.Name
	}

	if input.Description != nil {
		updates["description"] = *input.Description
	}

	if input.Image != nil {
		updates["image"] = *input.Image
	}

	if input.Skills != nil {
		updates["skills"] = datatypes.NewJSONSlice[string](*input.Skills)
	}

	if len(updates) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Aucun champ à mettre à jour"})
		return
	}

	if err := config.DB.Model(&project).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la mise à jour de la base de données"})
		return
	}

	c.JSON(http.StatusOK, project)
}
