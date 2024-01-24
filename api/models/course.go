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

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (c *Course) Validate() error {
	return validate.Struct(c)
}
