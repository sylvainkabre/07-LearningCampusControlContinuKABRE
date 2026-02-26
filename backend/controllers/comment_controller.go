package controllers

import (
	"LearningCampusControlContinu/config"
	"LearningCampusControlContinu/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"})
		return
	}

	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Utilisateur non authentifié"})
	}

	userIDint, ok := userID.(uint)

	comment.UserID = uint(userIDint)

	if err := config.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erreur lors de la création du commentaire"})
		return
	}

	c.JSON(http.StatusCreated, comment)
}
