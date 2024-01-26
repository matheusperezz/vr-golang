package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name    string          `json:"name" gorm:"size:50;not null" validate:"required"`
	Courses []CourseStudent `gorm:"foreignKey:StudentCode"`
}

type StudentDto struct {
	ID      uint     `json:"ID"`
	Name    string   `json:"name"`
	Courses []Course `json:"courses"`
}

func (s *Student) Validate() error {
	return validate.Struct(s)
}

func ConvertStudentToStudentDto(student Student, courses []Course) StudentDto {
	return StudentDto{
		ID:      student.ID,
		Name:    student.Name,
		Courses: courses,
	}
}
