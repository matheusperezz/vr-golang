package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name    string          `json:"name" gorm:"size:50;not null" validate:"required"`
	Courses []CourseStudent `gorm:"foreignKey:StudentCode"`
}

func (s *Student) Validate() error {
	return validate.Struct(s)
}
