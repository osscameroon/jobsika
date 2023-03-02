package storage

import (
	"github.com/elhmn/jobsika/pkg/models/v1beta"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (db DB) GetJobOffers(page string, limit string, jobtitle string, company string, city string, isRemote string) (v1beta.JobOffersResponse, error) {
	offset, limitInt := Paginate(page, limit)
	var nbHits int64

	// Build the query
	query := db.queryJobOffers().Order("jb.createdat DESC")

	if jobtitle != "" {
		query = query.Where("j.title LIKE ?", "%"+jobtitle+"%")
	}

	if company != "" {
		query = query.Where("c.name LIKE ?", "%"+company+"%")
	}

	if city != "" {
		query = query.Where("ct.name LIKE ?", "%"+city+"%")
	}

	switch isRemote {
	case "true":
		query = query.Where("jb.seniority = ?", true)
		break
	case "false":
		query = query.Where("jb.seniority = ?", false)
		break
	}

	rows, err := query.Count(&nbHits).Offset(offset).Limit(limitInt).Rows()
	if err != nil {
		return v1beta.JobOffersResponse{}, err
	}

	var jobOffers []v1beta.JobOfferPresenter
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
