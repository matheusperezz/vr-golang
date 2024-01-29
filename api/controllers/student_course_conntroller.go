package controllers

import (
	"fmt"
	"main/api/database"
	"main/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EnrollStudentInCourse(c *gin.Context) {
	var courseStudent models.CourseStudent
	if err := c.ShouldBindJSON(&courseStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("tentando converter para json")
		fmt.Println(courseStudent)
		return
	}

	var studentsInCourse int64
	database.DB.Model(&models.CourseStudent{}).Where("course_code = ?", courseStudent.CourseCode).Count(&studentsInCourse)
	if studentsInCourse >= 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Turma cheia"})
		fmt.Println("Verificando se há vagas na turma")
		fmt.Println(studentsInCourse)
		return
	}

	var coursesWithStudent int64
	database.DB.Model(&models.CourseStudent{}).Where("student_code = ?", courseStudent.StudentCode).Count(&coursesWithStudent)
	if coursesWithStudent >= 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "O aluno já está matriculado no número máximo de cursos permitidos"})
		fmt.Println("Verificando se o aluno já está matriculado no número máximo de cursos permitidos")
		fmt.Println(coursesWithStudent)
		return
	}

	// verifique se o aluno já está matriculado no curso
	var courseStudentAlreadyExists models.CourseStudent
	if err := database.DB.Where("course_code = ? AND student_code = ?", courseStudent.CourseCode, courseStudent.StudentCode).First(&courseStudentAlreadyExists).Error; err != nil {
		if err.Error() != "record not found" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			fmt.Println("tentando encontrar o registro")
			fmt.Println(courseStudent)
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "O aluno já está matriculado no curso"})
		fmt.Println("Verificando se o aluno já está matriculado no curso")
		fmt.Println(courseStudentAlreadyExists)
		return
	}

	if err := database.DB.Create(&courseStudent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("tentando criar o registro")
		fmt.Println(courseStudent)
		return
	}

	c.JSON(http.StatusOK, courseStudent)
}

func UnenrollStudentInCourse(c *gin.Context) {
	var courseStudent models.CourseStudent
	if err := c.ShouldBindJSON(&courseStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("tentando converter para json")
		fmt.Println(courseStudent)
		return
	}

	if err := database.DB.Where("course_code = ? AND student_code = ?", courseStudent.CourseCode, courseStudent.StudentCode).Delete(&courseStudent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("tentando deletar o registro")
		fmt.Println(courseStudent)
		return
	}

	c.JSON(http.StatusOK, courseStudent)
}
