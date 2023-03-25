package storage

import "github.com/osscameroon/jobsika/pkg/models/v1beta"

//GetCompanyRatings get company ratings
func (db DB) GetCompanyRatings(query v1beta.CompanyRatingQuery) ([]v1beta.CompanyRating, error) {
	if query.CompanyID != 0 {
		ratings := []v1beta.CompanyRating{}
		ret := db.c.Where("company_id = ?", query.CompanyID).Find(&ratings)
		if ret.Error != nil {
			return ratings, ret.Error
		}

		return ratings, nil
	}

	ratings := []v1beta.CompanyRating{}
	ret := db.c.Find(&ratings)
	if ret.Error != nil {
		return ratings, ret.Error
	}

	return ratings, nil
}

//GetCompanyRatingsByID get company ratings by ID
func (db DB) GetCompanyRatingsByID(id int64) (v1beta.CompanyRating, error) {
	rating := v1beta.CompanyRating{}
	ret := db.c.First(&rating, "id = ?", id)
	if ret.Error != nil {
		return rating, ret.Error
	}

	return rating, nil
}
