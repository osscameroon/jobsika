package storage

import (
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
		j.HasImage = j.CompanyImageLocation != ""
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
// the uploadImageFunc is a function that will be called after the job offer is created
// we made it to avoid circular dependency between storage and server package
func (db DB) PostJobOffer(query v1beta.OfferPostQuery, uploadImageFunc func(jobOfferID int64) (string, error)) (*v1beta.JobOffer, error) {
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

		if uploadImageFunc != nil {
			offer.CompanyImageLocation, err = uploadImageFunc(offer.ID)
			if err != nil {
				log.Error(err)
				return err
			}

			err = tx.Table("job_offers").Where("id = ?", offer.ID).Update("company_image_location", offer.CompanyImageLocation).Error
			if err != nil {
				log.Error(err)
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return offer, nil
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
		jb.company_image_location,
		jt.title as job_title
	`).
		Joins("left join jobtitles as jt on jb.title_id = jt.id")
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
