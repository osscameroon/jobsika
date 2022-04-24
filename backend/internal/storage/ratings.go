package storage

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/elhmn/camerdevs/pkg/models/v1beta"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//maxEntryBeforeDisplay The maximum number of entries before we can display the
//company name, and any other sensitive data
const maxEntryBeforeDisplay = 3

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
	return db.c.Table("salaries as s").Select(fmt.Sprintf(`
		s.id as salary_id,
		s.seniority,
		j.title as job_title,
		s.salary,
		s.city,
		s.country,
		r.id as company_rating_id,
		r.createdat,
		case
           when (
                    Select count(ss.id)
                    from salaries ss
                    where ss.company_id = s.company_id
						and ss.title_id = s.title_id) < %[1]d then ''
           else comment end as comment,
		case
           when (
                    Select count(ss.id)
                    from salaries ss
                    where ss.company_id = s.company_id
						and ss.title_id = s.title_id) < %[1]d then ''
           else c.name end as company_name,
		case
           when (
                    Select count(ss.id)
                    from salaries ss
                    where ss.company_id = s.company_id
						and ss.title_id = s.title_id) < %[1]d then 0
           else c.id end as company_id,
		case
           when (
                    Select count(ss.id)
                    from salaries ss
                    where ss.company_id = s.company_id
						and ss.title_id = s.title_id) < %[1]d then NULL
           else r.rating end as rating
		`, maxEntryBeforeDisplay)).
		Joins("LEFT JOIN companies c ON s.company_id = c.id").
		Joins("LEFT JOIN jobtitles j ON s.title_id = j.id").
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
		query = query.Where("j.title = ?", jobtitle)
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
func (db DB) GetAverageRating(jobtitle string, company string) (v1beta.AverageRating, error) {
	r := struct {
		AVGRating string `gorm:"column:rating"`
		AVGSalary string `gorm:"column:salary"`
	}{}

	query := db.queryRatings()
	if company != "" {
		query = query.Where("c.name = ?", company)
	}
	if jobtitle != "" {
		query = query.Where("j.title = ?", jobtitle)
	}

	ret := query.Select("AVG(s.salary) as salary, AVG(NULLIF (r.rating, 0)) as rating").Find(&r)
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

//PostRatings post new rating
func (db DB) PostRatings(query v1beta.RatingPostQuery) error {
	query.CompanyName = strings.Title(strings.ToLower(query.CompanyName))
	query.JobTitle = strings.Title(strings.ToLower(query.JobTitle))
	query.City = strings.Title(strings.ToLower(query.City))

	return db.c.Transaction(func(tx *gorm.DB) error {
		company, err := postCompany(tx, query)
		if err != nil {
			log.Error(err)
			return err
		}

		jobTitle, err := postJobTitle(tx, query)
		if err != nil {
			log.Error(err)
			return err
		}

		//Create company_rating
		companyRating := v1beta.CompanyRating{
			CompanyID: company.ID,
			Rating:    query.Rating,
			Comment:   query.Comment,
		}
		ret := tx.Create(&companyRating)
		if ret.Error != nil {
			log.Error(ret.Error)
			return ret.Error
		}

		//Create a salary
		salary := v1beta.Salary{
			TitleID:         jobTitle.ID,
			Salary:          query.Salary,
			Seniority:       query.Seniority,
			City:            query.City,
			Country:         "Cameroon",
			CompanyID:       company.ID,
			CompanyRatingID: companyRating.ID,
		}
		ret = tx.Create(&salary)
		if ret.Error != nil {
			log.Error(ret.Error)
			return ret.Error
		}

		return nil
	})
}

func postCompany(tx *gorm.DB, query v1beta.RatingPostQuery) (v1beta.Company, error) {
	company := v1beta.Company{}

	//Check if the company already exist
	ret := tx.Table("companies").Where("name = ?", query.CompanyName).Find(&company)
	if ret.Error != nil {
		return company, ret.Error
	}

	if company.Name == "" {
		//if we did not find the company, we create a new company entry
		company = v1beta.Company{
			Name:   query.CompanyName,
			Rating: query.Rating,
		}
		ret := tx.Table("companies").Create(&company)
		if ret.Error != nil {
			return company, ret.Error
		}
	}

	return company, nil
}

func postJobTitle(tx *gorm.DB, query v1beta.RatingPostQuery) (v1beta.JobTitle, error) {
	jobTitle := v1beta.JobTitle{}

	//Check if the jobTitle already exist
	ret := tx.Table("jobtitles").Where("title = ?", query.JobTitle).Find(&jobTitle)
	if ret.Error != nil {
		return jobTitle, ret.Error
	}

	if jobTitle.Title == "" {
		//if we did not find the jobTitle, we create a new jobTitle entry
		jobTitle = v1beta.JobTitle{
			Title: query.JobTitle,
		}
		ret := tx.Table("jobtitles").Create(&jobTitle)
		if ret.Error != nil {
			return jobTitle, ret.Error
		}
	}

	return jobTitle, nil
}
