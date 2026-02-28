package tests

import (
	"LearningCampusControlContinu/config"
	"LearningCampusControlContinu/controllers"
	"LearningCampusControlContinu/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetuptestDB() *gorm.DB {

	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.Project{}, &models.User{}, &models.Comment{})

	project := models.Project{Name: "Test Project", Description: "Test projet pour les tests"}
	db.Create(&project)

	comment := models.Comment{ProjectID: project.ID, Content: "Test comentaire pour les tests"}
	db.Create(&comment)

	return db
}

func TestGetProjects(t *testing.T) {
	gin.SetMode(gin.TestMode)

	config.DB = SetuptestDB()

	router := gin.Default()
	router.GET("/projects", controllers.GetProjects)

	req, _ := http.NewRequest(http.MethodGet, "/projects", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	body := resp.Body.String()
	assert.Contains(t, body, "Test Project")
	assert.Contains(t, body, "Test projet pour les tests")
}

func TestPostProject(t *testing.T) {
	gin.SetMode(gin.TestMode)

	config.DB = SetuptestDB()

	router := gin.Default()
	router.POST("/projects", controllers.PostProject)

	project := map[string]interface{}{
		"name":        "Nouveau Test Projet",
		"description": "Nouvelle description projet pour les tests",
		"skills":      []string{"Go", "Testing", "SQLite"},
	}

	data, _ := json.Marshal(project)

	req, _ := http.NewRequest(http.MethodPost, "/projects", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	assert.Contains(t, resp.Body.String(), "Nouveau Test Projet")
	assert.Contains(t, resp.Body.String(), "Nouvelle description projet pour les tests")
	assert.Contains(t, resp.Body.String(), "Go")
	assert.Contains(t, resp.Body.String(), "Testing")
	assert.Contains(t, resp.Body.String(), "SQLite")
}
