package models

type Student struct {
	Code    uint            `json:"code" sql:"" gorm:"primary_key"`
	Name    string          `json:"name" gorm:"size:50"`
	Courses []CourseStudent `json:"courses" gorm:"foreignKey:StudentCode"`
}

func (s *Student) Validate() error {
	return validate.Struct(s)
}
