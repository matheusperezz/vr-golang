package controllers

import (
	"main/api/database"
	"main/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {
	var students []models.Student
	if err := database.DB.Find(&students).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if len(students) <= 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No students found!"})
		return
	}

	var studentsDto []models.StudentDto
	for _, student := range students {
		var coursesStudents []models.CourseStudent
		if err := database.DB.Where("student_code = ?", student.ID).Find(&coursesStudents).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		var courses []models.Course
		for _, courseStudent := range coursesStudents {
			var course models.Course
			if err := database.DB.Where("id = ?", courseStudent.CourseCode).First(&course).Error; err != nil {
				c.AbortWithStatus(http.StatusNotFound)
				return
			}
			courses = append(courses, course)
		}

		if courses == nil {
			courses = []models.Course{}
		}

		studentDto := models.ConvertStudentToStudentDto(student, courses)
		studentsDto = append(studentsDto, studentDto)
	}

	c.JSON(http.StatusOK, studentsDto)
}

func GetStudentById(c *gin.Context) {
	var student models.Student
	studentId := c.Params.ByName("id")
	if err := database.DB.Where("id = ?", studentId).First(&student).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var coursesStudents []models.CourseStudent
	if err := database.DB.Where("student_code = ?", studentId).Find(&coursesStudents).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var courses []models.Course
	for _, courseStudent := range coursesStudents {
		var course models.Course
		if err := database.DB.Where("id = ?", courseStudent.CourseCode).First(&course).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		courses = append(courses, course)
	}

	studentDto := models.ConvertStudentToStudentDto(student, courses)

	c.JSON(http.StatusOK, studentDto)
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := student.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(student.Courses) == 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Estudantes não podem se matricular em mais de 3 cursos"})
		return
	}

	if dbErr := database.DB.Create(&student).Error; dbErr != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, student)
}

func UpdateStudent(c *gin.Context) {
	var student models.Student
	studentId := c.Params.ByName("id")
	if err := database.DB.Where("id = ?", studentId).First(&student).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if len(student.Courses) == 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Estudantes não podem se matricular em mais de 3 cursos"})
		return
	}

	if dbErr := database.DB.Save(&student).Error; dbErr != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	var student models.Student
	studentId := c.Params.ByName("id")
	if err := database.DB.Where("id = ?", studentId).First(&student).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// Remove o estudante de todos os cursos
	var courseStudents []models.CourseStudent
	if err := database.DB.Where("student_code = ?", studentId).Find(&courseStudents).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	for _, courseStudent := range courseStudents {
		if err := database.DB.Delete(&courseStudent).Error; err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	if err := database.DB.Delete(&student).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully!"})
}

func GetStudentCourses(c *gin.Context) {
	studentId := c.Params.ByName("id")

	var courseStudents []models.CourseStudent
	if err := database.DB.Where("student_code = ?", studentId).Find(&courseStudents).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var courses []models.Course
	for _, courseStudent := range courseStudents {
		var course models.Course
		if err := database.DB.Where("id = ?", courseStudent.CourseCode).First(&course).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		courses = append(courses, course)
	}

	if len(courses) <= 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No courses found!"})
		return
	}

	c.JSON(http.StatusOK, courses)
}
