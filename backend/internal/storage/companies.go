package storage

import "github.com/elhmn/camerdevs/pkg/models/v1beta"

//GetCompanies get companies
func (db DB) GetCompanies() ([]v1beta.Company, error) {
	companies := []v1beta.Company{}
	ret := db.c.Find(&companies)
	if ret.Error != nil {
		return companies, ret.Error
	}

	return companies, nil
}
