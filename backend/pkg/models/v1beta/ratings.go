package v1beta

import "time"

//RatingQuery is a Rating query structure
type RatingQuery struct {
	ID int64 `form:"rating_id"`
}

//Rating defines the rating structure
//This structure combines a salary and a company rating entry
type Rating struct {
	SalaryID        int64     `json:"salary_id" gorm:"column:salary_id"`
	CompanyID       int64     `json:"company_id" gorm:"column:company_id"`
	CompanyRatingID int64     `json:"company_rating_id" gorm:"column:company_rating_id"`
	Rating          int64     `json:"rating" gorm:"column:rating"`
	Salary          int64     `json:"salary" gorm:"column:salary"`
	CompanyName     string    `json:"company_name" gorm:"column:company_name"`
	Seniority       string    `json:"seniority" gorm:"column:seniority"`
	Comment         string    `json:"comment" gorm:"column:comment"`
	JobTitle        string    `json:"job_title" gorm:"column:job_title"`
	Country         string    `json:"country" gorm:"column:country"`
	City            string    `json:"city" gorm:"column:city"`
	CreatedAt       time.Time `json:"createdat" gorm:"column:createdat"`
}

type RatingResponse struct {
	Hits   []Rating `json:"hits"`
	Limit  int64    `json:"limit"`
	NBHits int64    `json:"nbHits"`
	Offset int64    `json:"offset"`
}

//AverageRating defines the average rating structure
type AverageRating struct {
	Rating int64 `json:"rating" gorm:"column:rating"`
	Salary int64 `json:"salary" gorm:"column:salary"`
}

//RatingPostQuery defines the body object used to
//create a new query
type RatingPostQuery struct {
	CompanyName string `json:"company_name"`
	JobTitle    string `json:"job_title"`
	Salary      int64  `json:"salary"`
	City        string `json:"city"`
	Comment     string `json:"comment"`
	Seniority   string `json:"seniority"`
	Rating      int64  `json:"rating"`
}
