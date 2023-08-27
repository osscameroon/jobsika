package storage

import (
	"time"

	"github.com/osscameroon/jobsika/pkg/models/v1beta"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (db DB) GetJobOffers(q v1beta.GetJobOffersQuery) (v1beta.JobOffersResponse, error) {
	offset, limitInt := Paginate(q.Page, q.Limit)
	var nbHits int64

	// Build the query
	query := db.queryJobOffers().Order("jb.createdat DESC")

	if q.JobTitle != "" {
		query = query.Where("j.title LIKE ?", "%"+q.JobTitle+"%")
	}

	if q.Company != "" {
		query = query.Where("jb.company_name LIKE ?", "%"+q.Company+"%")
	}

	if q.City != "" {
		query = query.Where("jb.city LIKE ?", "%"+q.City+"%")
	}

	if q.Country != "" {
		query = query.Where("jb.country LIKE ?", "%"+q.Country+"%")
	}

	switch q.IsRemote {
	case "true":
		query = query.Where("jb.is_remote = ?", true)
	case "false":
		query = query.Where("jb.is_remote = ?", false)
	}

	query = query.Order("jb.createdat DESC").Order("jb.id DESC")

	rows, err := query.Count(&nbHits).Offset(offset).Limit(limitInt).Rows()
	if err != nil {
		return v1beta.JobOffersResponse{}, err
	}

	jobOffers := make([]v1beta.JobOfferPresenter, 0)
	for rows.Next() {
		j := v1beta.JobOfferPresenter{}
		err := db.c.ScanRows(rows, &j)
		if err != nil {
			return v1beta.JobOffersResponse{}, err
		}
		jobOffers = append(jobOffers, j)
	}

	resp := v1beta.JobOffersResponse{
		Hits:   jobOffers,
		NbHits: nbHits,
		Offset: int64(offset),
		Limit:  int64(limitInt),
	}

	return resp, nil
}

// PostJobOffer post new job offer
func (db DB) PostJobOffer(query v1beta.OfferPostQuery) (*v1beta.JobOffer, error) {
	currentTime := time.Now().UTC()
	offer := &v1beta.JobOffer{
		CompanyName:             query.CompanyName,
		CompanyEmail:            query.CompanyEmail,
		IsRemote:                query.IsRemote,
		City:                    query.City,
		Country:                 query.Country,
		Department:              query.Department,
		SalaryRangeMin:          query.SalaryRangeMin,
		SalaryRangeMax:          query.SalaryRangeMax,
		Description:             query.Description,
		Benefits:                query.Benefits,
		HowToApply:              query.HowToApply,
		ApplicationUrl:          query.ApplicationUrl,
		ApplicationEmailAddress: query.ApplicationEmailAddress,
		ApplicationPhoneNumber:  query.ApplicationPhoneNumber,
		Tags:                    query.Tags,
		CreatedAt:               currentTime,
		UpdatedAt:               currentTime,
	}

	err := db.c.Transaction(func(tx *gorm.DB) error {
		jobTitle, err := postJobTitle(tx, query.JobTitle)
		if err != nil {
			log.Error(err)
			return err
		}

		offer.TitleID = jobTitle.ID
		res := tx.Table("job_offers").Create(offer)
		if res.Error != nil {
			log.Error(res.Error)
			return res.Error
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return offer, nil
}

// PostJobOfferImage save job offer image location in the database
func (db DB) PostJobOfferImage(offerID int64, imageLocation string) error {
	currentTime := time.Now().UTC()
	jobOfferImage := v1beta.JobOfferImage{
		JobOfferID:    offerID,
		ImageLocation: imageLocation,
		CreatedAt:     currentTime,
		UpdatedAt:     currentTime,
	}

	res := db.c.Table("job_offers_image").Create(&jobOfferImage)
	if res.Error != nil {
		log.Error(res.Error)
		return res.Error
	}

	return nil
}

// DeleteJobOffer delete job offer
func (db DB) DeleteJobOffer(id int64) error {
	res := db.c.Table("job_offers").Delete(&v1beta.JobOffer{}, id)
	if res.Error != nil {
		log.Error(res.Error)
		return res.Error
	}

	return nil
}

func (db DB) queryJobOffers() *gorm.DB {
	return db.c.Table("job_offers as jb").Select(`
		jb.id,
		jb.createdat,
		jb.updatedat,
		jb.company_name,
		jb.company_email,
		jb.title_id,
		jb.is_remote,
		jb.city,
		jb.country,
		jb.department,
		jb.salary_range_min,
		jb.salary_range_max,
		jb.description,
		jb.benefits,
		jb.how_to_apply,
		jb.application_url,
		jb.application_email_address,
		jb.application_phone_number,
		jb.tags,
		jt.title as job_title,
		CASE
        	WHEN jb_image.image_location <> '' THEN 1
        	ELSE 0
    	END AS has_image
	`).
		Joins("left join jobtitles as jt on jb.title_id = jt.id").
		Joins("left join job_offers_image as jb_image on jb.id = jb_image.job_offer_id")
}

// GetJobOfferById get job offers by id
func (db DB) GetJobOfferById(id int64) (*v1beta.JobOffer, error) {
	var offer v1beta.JobOffer
	res := db.queryJobOffers().Where("jb.id = ?", id).First(&offer)
	if res.Error != nil {
		log.Error(res.Error)
		return nil, res.Error
	}

	return &offer, nil
}

// GetJobOfferImageByJobOfferId get job offer image by job offer id
func (db DB) GetJobOfferImageByJobOfferId(jobOfferID int64) (*v1beta.JobOfferImage, error) {
	var image v1beta.JobOfferImage
	res := db.c.Table("job_offers_image").Where("job_offer_id = ?", jobOfferID).First(&image)
	if res.Error != nil {
		log.Error(res.Error)
		return nil, res.Error
	}

	return &image, nil
}
