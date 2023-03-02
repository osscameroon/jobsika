package v1beta

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

// JobOffer defines the job_offers structure
type JobOffer struct {
	ID        int64     `json:"id" gorm:"column:id"`
	CreatedAt time.Time `json:"createdat" gorm:"column:createdat"`
	UpdatedAt time.Time `json:"updatedat" gorm:"column:updatedat"`

	Email             string `json:"email" gorm:"column:email"`
	CompanyName       string `json:"company_name" gorm:"column:company_name"`
	TitleID           int64  `json:"title" gorm:"column:title_id"`
	IsRemote          bool   `json:"is_remote" gorm:"column:is_remote"`
	Description       string `json:"description" gorm:"column:description"`
	HowToApply        string `json:"how_to_apply" gorm:"column:how_to_apply"`
	ApplyUrl          string `json:"apply_url" gorm:"column:apply_url"`
	ApplyEmailAddress string `json:"apply_email_address" gorm:"column:apply_email_address"`
}

// OfferPostQuery defines the body object used to create a new jb offer on a POST query
type OfferPostQuery struct {
	Email             string `json:"email"`
	CompanyName       string `json:"company_name"`
	JobTitle          string `json:"job_title"`
	IsRemote          bool   `json:"is_remote"`
	Description       string `json:"description"`
	HowToApply        string `json:"how_to_apply"`
	ApplyUrl          string `json:"apply_url"`
	ApplyEmailAddress string `json:"apply_email_address"`
}

// Validate check if the mandatory fields are filled
func (r OfferPostQuery) Validate() error {
	if strings.TrimSpace(r.Email) == "" {
		return errors.New("email is mandatory")
	}

	emailRegex := regexp.MustCompile(`(?i)[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}`)
	if !emailRegex.MatchString(r.Email) {
		return errors.New("email is not a valid email address")
	}

	if strings.TrimSpace(r.CompanyName) == "" {
		return errors.New("company name is mandatory")
	}

	if strings.TrimSpace(r.JobTitle) == "" {
		return errors.New("job title is mandatory")
	}

	if strings.TrimSpace(r.Description) == "" {
		return errors.New("job description is mandatory")
	}

	if strings.TrimSpace(r.ApplyUrl) == "" && strings.TrimSpace(r.ApplyEmailAddress) == "" {
		return errors.New("apply url or apply email address is mandatory")
	}

	if strings.TrimSpace(r.ApplyEmailAddress) != "" {
		if !emailRegex.MatchString(r.ApplyEmailAddress) {
			return errors.New("apply email address is not a valid email address")
		}
	}

	return nil
}

type JobOfferPresenter struct {
	ID        int64     `json:"id" gorm:"column:id"`
	CreatedAt time.Time `json:"createdat" gorm:"column:createdat"`
	UpdatedAt time.Time `json:"updatedat" gorm:"column:updatedat"`

	Email             string `json:"email"`
	CompanyName       string `json:"company_name"`
	JobTitle          string `json:"job_title"`
	IsRemote          bool   `json:"is_remote"`
	Description       string `json:"description"`
	HowToApply        string `json:"how_to_apply"`
	ApplyUrl          string `json:"apply_url"`
	ApplyEmailAddress string `json:"apply_email_address"`
}

type JobOffersResponse struct {
	Hits   []JobOfferPresenter `json:"hits"`
	NbHits int64               `json:"nbHits"`
	Offset int64               `json:"offset"`
	Limit  int64               `json:"limit"`
}

type GetJobOffersQuery struct {
	Page     string
	Offset   string
	Limit    string
	JobTitle string
	Company  string
	City     string
	IsRemote string
}
