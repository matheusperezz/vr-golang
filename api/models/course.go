package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	Description string          `json:"description" gorm:"size:50;not null" validate:"required"`
	Syllabus    string          `json:"syllabus" gorm:"type:text;not null" validate:"required"`
	Students    []CourseStudent `gorm:"foreignKey:CourseCode"`
}

type CourseDto struct {
	ID          uint         `json:"ID"`
	Description string       `json:"description"`
	Syllabus    string       `json:"syllabus"`
	Students    []StudentDto `json:"students"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (c *Course) Validate() error {
	return validate.Struct(c)
}

func ConvertCourseToCourseDto(course Course, students []Student) CourseDto {
	var studentDtos []StudentDto
	for _, student := range students {
		studentDtos = append(studentDtos, ConvertStudentToStudentDto(student, []Course{}))
	}
	return CourseDto{
		ID:          course.ID,
		Description: course.Description,
		Syllabus:    course.Syllabus,
		Students:    studentDtos,
	}
}
