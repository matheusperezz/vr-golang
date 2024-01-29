package models

import "gorm.io/gorm"

type CourseStudent struct {
	gorm.Model
	CourseCode  uint `json:"course_code" validate:"required"`
	StudentCode uint `json:"student_code" validate:"required"`
}
