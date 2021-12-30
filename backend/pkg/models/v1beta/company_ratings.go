package v1beta

import "time"

//CompanyRatingQuery is a company Rating query structure
type CompanyRatingQuery struct {
	CompanyID int64 `form:"company_id"`
}

//CompanyRating defines the company rating structure
type CompanyRating struct {
	ID        int64     `json:"id" gorm:"column:id"`
	CompanyID int64     `json:"company_id" gorm:"column:company_id"`
	Rating    int64     `json:"rating" gorm:"column:rating"`
	Comment   string    `json:"comment" gorm:"column:comment"`
	CreatedAt time.Time `json:"createdat" gorm:"column:createdat"`
	UpdatedAt time.Time `json:"updatedat" gorm:"column:updatedat"`
}
