package models

import "github.com/go-playground/validator/v10"

type Course struct {
	Code        uint            `json:"code" gorm:"primary_key auto_increment"`
	Description string          `json:"description" gorm:"size:50"`
	Syllabus    string          `json:"syllabus" gorm:"type:text"`
	Students    []CourseStudent `json:"students" gorm:"foreignKey:CourseCode"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (c *Course) Validate() error {
	return validate.Struct(c)
}
