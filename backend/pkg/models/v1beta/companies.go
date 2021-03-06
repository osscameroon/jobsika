package v1beta

import (
	"time"
)

//Company defines the company structure
type Company struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name      string    `json:"name" gorm:"column:name"`
	Rating    int64     `json:"rating" gorm:"column:rating"`
	CreatedAt time.Time `json:"createdat" gorm:"column:createdat"`
	UpdatedAt time.Time `json:"updatedat" gorm:"column:updatedat"`
}
