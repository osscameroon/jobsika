package storage

import (
	"github.com/elhmn/jobsika/pkg/models/v1beta"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// PostJobOffer post new job offer
func (db DB) PostJobOffer(query v1beta.OfferPostQuery) (*v1beta.JobOffer, error) {
	offer := v1beta.JobOffer{
		Email:             query.Email,
		CompanyName:       query.CompanyName,
		IsRemote:          query.IsRemote,
		Description:       query.Description,
		HowToApply:        query.HowToApply,
		ApplyUrl:          query.ApplyUrl,
		ApplyEmailAddress: query.ApplyEmailAddress,
	}

	err := db.c.Transaction(func(tx *gorm.DB) error {
		jobTitle, err := postJobTitle(tx, query.JobTitle)
		if err != nil {
			log.Error(err)
			return err
		}

		offer.TitleID = jobTitle.ID
		res := tx.Table("job_offers").Create(&offer)
		if res.Error != nil {
			log.Error(res.Error)
			return res.Error
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &offer, nil
}

func (db DB) queryJobOffers() *gorm.DB {
	return db.c.Table("job_offers as jb").Select(`
		jb.id,
		jb.createdat,
		jb.updatedat,
		jb.email,
		jb.company_name,
		jb.title_id,
		jb.is_remote,
		jb.description,
		jb.how_to_apply,
		jb.apply_url,
		jb.apply_email_address,
		jt.title as job_title
	`).
		Joins("left join job_titles as jt on jb.title_id = jt.id")
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
