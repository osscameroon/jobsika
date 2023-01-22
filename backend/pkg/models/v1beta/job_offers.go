package v1beta

import (
	"errors"
	"strings"
)

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

	return nil
}
