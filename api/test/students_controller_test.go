package test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"main/api/controllers"
	"main/api/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateStudent(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/students", controllers.CreateStudent)

	student := models.Student{
		Name: "Test Student",
	}

	jsonStudent, _ := json.Marshal(student)

	req, _ := http.NewRequest(http.MethodPost, "/students", bytes.NewBuffer(jsonStudent))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetAllStudents(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/students", controllers.GetAllStudents)

	req, _ := http.NewRequest(http.MethodGet, "/students", nil)

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetStudentById(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/students/:id", controllers.GetStudentById)

	req, _ := http.NewRequest(http.MethodGet, "/students/1", nil)

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
