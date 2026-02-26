package controllers

import (
	"LearningCampusControlContinu/config"
	"LearningCampusControlContinu/models"
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func GetProjects(c *gin.Context) {

	var projects []models.Project

	if err := config.DB.Preload("Comments").Preload("Likes").Find(&projects).Error; err != nil {
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

	if err := config.DB.Preload("Comments").Preload("Likes").First(&project, id).Error; err != nil {
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

	file, err := c.FormFile("image")

	if err == nil {
		imagePath := "uploads/" + file.Filename
		if err := c.SaveUploadedFile(file, imagePath); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Impossible d'enregistrer l'image"})
			return
		}

		img, _ := imaging.Open(imagePath)
		resize := imaging.Resize(img, 800, 0, imaging.Lanczos)
		if err := imaging.Save(resize, imagePath); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Impossible de redimensionner l'image"})
			return
		}
		project.Image = imagePath
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
	if err := c.ShouldBind(&input); err != nil {
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

	file, err := c.FormFile("image")

	if err == nil {
		imagePath := "uploads/" + file.Filename
		if err := c.SaveUploadedFile(file, imagePath); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Impossible d'enregistrer l'image"})
			return
		}

		img, _ := imaging.Open(imagePath)
		resize := imaging.Resize(img, 800, 0, imaging.Lanczos)
		if err := imaging.Save(resize, imagePath); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Impossible de redimensionner l'image"})
			return
		}

		if project.Image != "" {
			_ = os.Remove(project.Image)
		}
		updates["image"] = imagePath
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

func DeleteProject(c *gin.Context) {
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
	//Supprimer notre projet en BDD

	if err := config.DB.Delete(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la suppression du projet"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Le projet a été supprimé avec succès"})
}

func LikeProject(c *gin.Context) {
	// 1. Parse project ID
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID non valide"})
		return
	}

	// 2. Load project (sans préload inutile)
	var project models.Project
	if err := config.DB.First(&project, projectID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Projet non trouvé"})
		return
	}

	// 3. Récupération userID depuis le middleware
	rawUserID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Utilisateur introuvable"})
		return
	}

	userID, ok := rawUserID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Type user_id invalide"})
		return
	}

	// 4. Vérifier si le like existe déjà (sans charger tous les likes)
	var count int64
	err = config.DB.
		Table("project_likes").
		Where("project_id = ? AND user_id = ?", project.ID, userID).
		Count(&count).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur interne"})
		return
	}

	// 5. Toggle like
	user := models.User{ID: userID}

	if count > 0 {
		// Déjà liké → on retire
		if err := config.DB.Model(&project).Association("Likes").Delete(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la suppression du like"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Like retiré"})
		return
	}

	// Pas encore liké → on ajoute
	if err := config.DB.Model(&project).Association("Likes").Append(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de l'ajout du like"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Like ajouté"})
}
