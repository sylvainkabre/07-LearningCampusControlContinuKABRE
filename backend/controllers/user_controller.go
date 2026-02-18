package controllers

import (
	"LearningCampusControlContinu/config"
	"LearningCampusControlContinu/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"})
		return
	}

	var count int64
	config.DB.Model(&models.User{}).Where("email = ?", user.Email).Count(&count)
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "L'email est déjà utilisé"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur serveur"})
		return
	}

	user.Password = string(hashedPassword)

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de créer l'utilsiateur"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "utilisateur créé avec succès"})
}
