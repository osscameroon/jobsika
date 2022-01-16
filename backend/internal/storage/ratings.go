package storage

import (
	"errors"

	"github.com/elhmn/camerdevs/pkg/models/v1beta"
	"gorm.io/gorm"
)

func (db DB) queryRatings() *gorm.DB {
	return db.c.Table("salaries as s").Select(`
		s.id as salary_id,
		s.seniority,
		s.title as job_title,
		s.salary,
		s.city,
		s.country,
		c.id as company_id,
		c.name as company_name,
		r.id as company_rating_id,
		r.comment,
		r.rating`).
		Joins("LEFT JOIN companies c ON s.company_id = c.id").
		Joins("LEFT JOIN company_ratings r ON s.company_rating_id = r.id")
}

//GetRatings get ratings
func (db DB) GetRatings() ([]v1beta.Rating, error) {
	ratings := []v1beta.Rating{}

	rows, err := db.queryRatings().Order("salary_id").Rows()
	if err != nil {
		return ratings, err
	}

	for rows.Next() {
		r := v1beta.Rating{}
		err := db.c.ScanRows(rows, &r)
		if err != nil {
			return nil, err
		}
		ratings = append(ratings, r)
	}

	return ratings, nil
}

//GetRatingByID get rating by `id`
func (db DB) GetRatingByID(id int64) (v1beta.Rating, error) {
	rating := v1beta.Rating{}
	ret := db.queryRatings().Where("s.id = ?", id).Limit(1).Find(&rating)
	if ret.Error != nil {
		return rating, ret.Error
	}

	if rating.SalaryID == 0 {
		return rating, errors.New("record not found")
	}

	return rating, nil
}
