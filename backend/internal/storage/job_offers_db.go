package storage

import (
	"github.com/elhmn/jobsika/pkg/models/v1beta"
	"gorm.io/gorm"
	"strings"
)

// PostJobOffer post new job offer
func (db DB) PostJobOffer(query v1beta.OfferPostQuery) error {
	query.CompanyName = strings.Title(strings.ToLower(query.CompanyName))
	query.JobTitle = strings.Title(strings.ToLower(query.JobTitle))

	return db.c.Transaction(func(tx *gorm.DB) error {

		return nil
	})
}
