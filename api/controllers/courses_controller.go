package controllers

import (
	"main/api/database"
	"main/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
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

	var coursesDto []models.CourseDto
	for _, course := range courses {
		var studentsCourse []models.CourseStudent
		if err := database.DB.Where("course_code = ?", course.ID).Find(&studentsCourse).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		var students []models.Student
		for _, studentCourse := range studentsCourse {
			var student models.Student
			if err := database.DB.Where("id = ?", studentCourse.StudentCode).First(&student).Error; err != nil {
				c.AbortWithStatus(http.StatusNotFound)
				return
			}
			students = append(students, student)
		}

		courseDto := models.ConvertCourseToCourseDto(course, students)
		coursesDto = append(coursesDto, courseDto)
	}

	c.JSON(http.StatusOK, coursesDto)
}

func GetCourseById(c *gin.Context) {
	var course models.Course
	courseId := c.Params.ByName("id")
	if err := database.DB.Where("id = ?", courseId).First(&course).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var studentsCourse []models.CourseStudent
	if err := database.DB.Where("course_code = ?", courseId).Find(&studentsCourse).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var students []models.Student
	for _, studentCourse := range studentsCourse {
		var student models.Student
		if err := database.DB.Where("id = ?", studentCourse.StudentCode).First(&student).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		students = append(students, student)
	}

	courseDto := models.ConvertCourseToCourseDto(course, students)

	c.JSON(http.StatusOK, courseDto)
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

	if len(course.Students) == 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "O curso não pode ter mais de 10 alunos!"})
		return
	}

	if dbErr := database.DB.Create(&course).Error; dbErr != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, course)
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

	if len(course.Students) == 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "O curso não pode ter mais de 10 alunos!"})
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

	// Remove todos os estudantes do curso
	var courseStudents []models.CourseStudent
	if err := database.DB.Where("course_code = ?", courseId).Find(&courseStudents).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	for _, courseStudent := range courseStudents {
		if err := database.DB.Delete(&courseStudent).Error; err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
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
		if err := database.DB.Where("id = ?", courseStudent.StudentCode).First(&student).Error; err != nil {
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
