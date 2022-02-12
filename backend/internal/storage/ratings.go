package storage

import (
	"errors"
	"math"
	"strconv"

	"github.com/elhmn/camerdevs/pkg/models/v1beta"
	"gorm.io/gorm"
)

//Paginate returns a query scoped to a certain page
func Paginate(pageStr, limitStr string) (offset, limit int) {
	page, _ := strconv.Atoi(pageStr)
	if page == 0 {
		page = 1
	}

	limit, _ = strconv.Atoi(limitStr)
	switch {
	case limit <= 0:
		limit = 20
	}

	offset = (page - 1) * limit
	return offset, limit
}

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
		r.rating,
		r.createdat`).
		Joins("LEFT JOIN companies c ON s.company_id = c.id").
		Joins("LEFT JOIN company_ratings r ON s.company_rating_id = r.id")
}

//GetRatings get ratings
func (db DB) GetRatings(page, limit, jobtitle, company string) (v1beta.RatingResponse, error) {
	offset, limitInt := Paginate(page, limit)
	var nbHits int64

	query := db.queryRatings().Order("salary_id")
	if company != "" {
		query = query.Where("c.name = ?", company)
	}
	if jobtitle != "" {
		query = query.Where("s.title = ?", jobtitle)
	}

	rows, err := query.Count(&nbHits).Offset(offset).Limit(limitInt).Rows()
	if err != nil {
		return v1beta.RatingResponse{}, err
	}

	ratings := []v1beta.Rating{}
	for rows.Next() {
		r := v1beta.Rating{}
		err := db.c.ScanRows(rows, &r)
		if err != nil {
			return v1beta.RatingResponse{}, err
		}
		ratings = append(ratings, r)
	}

	resp := v1beta.RatingResponse{
		Hits:   ratings,
		NBHits: nbHits,
		Offset: int64(offset),
		Limit:  int64(limitInt),
	}

	return resp, nil
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

//GetAverageRating get average rating
func (db DB) GetAverageRating() (v1beta.AverageRating, error) {
	r := struct {
		AVGRating string `gorm:"column:rating"`
		AVGSalary string `gorm:"column:salary"`
	}{}

	ret := db.queryRatings().Select("AVG(s.salary) as salary, AVG(r.rating) as rating").Find(&r)
	if ret.Error != nil {
		return v1beta.AverageRating{}, ret.Error
	}

	vr, _ := strconv.ParseFloat(r.AVGRating, 64)
	vs, _ := strconv.ParseFloat(r.AVGSalary, 64)

	rating := v1beta.AverageRating{
		Rating: int64(math.Round(vr)),
		Salary: int64(math.Round(vs)),
	}

	return rating, nil
}
