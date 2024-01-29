package database

import (
	"fmt"
	"log"
	"main/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	connStr := "host=127.0.0.1 port=5432 user=myuser password=mypassword dbname=mydatabase sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	err := DB.AutoMigrate(&models.Course{}, &models.Student{}, &models.CourseStudent{})
	if err != nil {
		return
	}

	// Defining relations
	DB.Model(&models.Course{}).Association("students")
	DB.Model(&models.Student{}).Association("courses")

	fmt.Println("Tables created and relationships configured successfully!")
}
