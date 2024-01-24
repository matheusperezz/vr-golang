package routes

import (
	"github.com/gin-gonic/gin"
	"main/api/controllers"
)

func SetupRoutes(router *gin.Engine) *gin.Engine {
	home := router.Group("/")
	{
		students := home.Group("/students")
		{
			students.GET("/", controllers.GetAllStudents)
			students.GET("/:id", controllers.GetStudentById)
			students.GET("/:id/courses", controllers.GetStudentCourses)
			students.POST("/", controllers.CreateStudent)
			students.PUT("/:id", controllers.UpdateStudent)
			students.DELETE("/:id", controllers.DeleteStudent)
		}

		courses := home.Group("/courses")
		{
			courses.GET("/", controllers.GetAllCourses)
			courses.GET("/:id", controllers.GetCourseById)
			courses.GET("/:id/students", controllers.GetCourseStudents)
			courses.POST("/", controllers.CreateCourse)
			courses.PUT("/:id", controllers.UpdateCourse)
			courses.DELETE("/:id", controllers.DeleteCourse)
		}

		class := home.Group("/class")
		{
			class.POST("/", controllers.EnrollStudentInCourse)
		}
	}
	return router
}
