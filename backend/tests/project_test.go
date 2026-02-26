package tests

import (
	"LearningCampusControlContinu/config"
	"LearningCampusControlContinu/controllers"
	"LearningCampusControlContinu/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setuptestDB() *gorm.DB {

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

	config.DB = setuptestDB()

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
