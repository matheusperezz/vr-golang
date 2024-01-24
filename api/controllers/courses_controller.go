package controllers

import (
	"github.com/gin-gonic/gin"
	"main/api/database"
	"main/api/models"
	"net/http"
)

func GetAllCourses(c *gin.Context) {
	var courses []models.Course
	if err := database.DB.Find(&courses).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if len(courses) <= 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No courses found!"})
		return
	}

	c.JSON(http.StatusOK, courses)
}

func GetCourseById(c *gin.Context) {
	var course models.Course
	courseId := c.Params.ByName("id")
	if err := database.DB.Where("id = ?", courseId).First(&course).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, course)
}

func CreateCourse(c *gin.Context) {
	var course models.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := course.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if dbErr := database.DB.Create(&course).Error; dbErr != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, course)
}

func UpdateCourse(c *gin.Context) {
	var course models.Course
	courseId := c.Params.ByName("id")
	if err := database.DB.Where("id = ?", courseId).First(&course).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err := c.ShouldBindJSON(&course); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if dbErr := database.DB.Save(&course).Error; dbErr != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, course)
}

func DeleteCourse(c *gin.Context) {
	var course models.Course
	courseId := c.Params.ByName("id")
	if err := database.DB.Where("id = ?", courseId).First(&course).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if dbErr := database.DB.Delete(&course).Error; dbErr != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully!"})
}

func GetCourseStudents(c *gin.Context) {
	courseId := c.Params.ByName("id")

	var courseStudents []models.CourseStudent
	if err := database.DB.Where("course_code = ?", courseId).Find(&courseStudents).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var students []models.Student
	for _, courseStudent := range courseStudents {
		var student models.Student
		if err := database.DB.Where("code = ?", courseStudent.StudentCode).First(&student).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		students = append(students, student)
	}

	if len(students) <= 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No students found!"})
		return
	}

	c.JSON(http.StatusOK, students)
}
