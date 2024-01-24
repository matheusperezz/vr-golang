package controllers

import (
	"github.com/gin-gonic/gin"
	"main/api/database"
	"main/api/models"
	"net/http"
)

func EnrollStudentInCourse(c *gin.Context) {
	var courseStudent models.CourseStudent
	if err := c.ShouldBindJSON(&courseStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var studentsInCourse int64
	database.DB.Model(&models.CourseStudent{}).Where("course_code = ?", courseStudent.CourseCode).Count(&studentsInCourse)
	if studentsInCourse >= 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Turma cheia"})
		return
	}

	var coursesWithStudent int64
	database.DB.Model(&models.CourseStudent{}).Where("student_code = ?", courseStudent.StudentCode).Count(&coursesWithStudent)
	if coursesWithStudent >= 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "O aluno já está matriculado no número máximo de cursos permitidos"})
		return
	}

	if err := database.DB.Create(&courseStudent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, courseStudent)
}
