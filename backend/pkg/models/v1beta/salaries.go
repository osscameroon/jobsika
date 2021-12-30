package v1beta

import "time"

//Salary defines the salary structure
type Salary struct {
	Title  string `json:"title" gorm:"column:title"`
	Salary int64  `json:"salary" gorm:"column:salary"`
	City   string `json:"city" gorm:"column:city"`
	//The Company field should be shown only if there are more than 4 salary
	//records with the same company
	Company   string    `json:"company" gorm:"column:company"`
	CreatedAt time.Time `json:"createdat" gorm:"column:createdat"`
	UpdatedAt time.Time `json:"updatedat" gorm:"column:updatedat"`
}
