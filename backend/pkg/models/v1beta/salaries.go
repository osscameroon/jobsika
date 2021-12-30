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
	ID     int64  `json:"id" gorm:"column:id"`
	Title  string `json:"title" gorm:"column:title"`
	Salary int64  `json:"salary" gorm:"column:salary"`
	//Seniority can be set to the values bellow
	//Intern, Entry-level, Mid-level, Senior, Above Senior level, Executive
	Seniority string `json:"seniority" gorm:"column:seniority"`
	City      string `json:"city" gorm:"column:city"`
	//The Company field should be shown only if there are more than 4 salary
	//records with the same company
	Company   string    `json:"company" gorm:"column:company"`
	CreatedAt time.Time `json:"createdat" gorm:"column:createdat"`
	UpdatedAt time.Time `json:"updatedat" gorm:"column:updatedat"`
}
