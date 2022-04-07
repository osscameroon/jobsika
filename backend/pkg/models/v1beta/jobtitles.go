package v1beta

import (
	"time"
)

//JobTitle defines the jobtitle structure
type JobTitle struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Title     string    `json:"title" gorm:"column:title"`
	CreatedAt time.Time `json:"createdat" gorm:"column:createdat"`
	UpdatedAt time.Time `json:"updatedat" gorm:"column:updatedat"`
}
