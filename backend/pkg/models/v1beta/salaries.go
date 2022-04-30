package v1beta

import "time"

const (
	SeniorityIntern      = "Intern"
	SeniorityEntryLevel  = "Entry-level"
	SeniorityMidLevel    = "Mid-level"
	SenioritySenior      = "Senior"
	SeniorityAboveSenior = "Above senior-level"
	SeniorityExecutive   = "Executive"
)

//Salary defines the salary structure
type Salary struct {
	ID      int64 `json:"id" gorm:"column:id"`
	TitleID int64 `json:"title" gorm:"column:title_id"`
	Salary  int64 `json:"salary" gorm:"column:salary"`
	//Seniority can be set to the values bellow
	//Intern, Entry-level, Mid-level, Senior, Above Senior level, Executive
	Seniority       string    `json:"seniority" gorm:"column:seniority"`
	CityID          int64     `json:"city_id" gorm:"column:city_id"`
	Country         string    `json:"country" gorm:"column:country"`
	CompanyID       int64     `json:"company_id" gorm:"column:company_id"`
	CompanyRatingID int64     `json:"company_rating_id" gorm:"column:company_rating_id"`
	CreatedAt       time.Time `json:"createdat" gorm:"column:createdat"`

	UpdatedAt time.Time `json:"updatedat" gorm:"column:updatedat"`
}
