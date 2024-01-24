package models

type CourseStudent struct {
	Code        uint `json:"code" gorm:"primary_key"`
	CourseCode  uint `json:"course_code"`
	StudentCode uint `json:"student_code"`
}
