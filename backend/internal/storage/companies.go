package storage

import "github.com/elhmn/camerdevs/pkg/models/v1beta"

//GetCompanies get companies
func (db DB) GetCompanies() ([]v1beta.Company, error) {
	companies := []v1beta.Company{}
	ret := db.c.Table("companies").Order("name").Find(&companies)
	if ret.Error != nil {
		return companies, ret.Error
	}

	return companies, nil
}

//GetCompanyByID get company by `id`
func (db DB) GetCompanyByID(id int64) (v1beta.Company, error) {
	company := v1beta.Company{}
	ret := db.c.First(&company, "id = ?", id)
	if ret.Error != nil {
		return company, ret.Error
	}

	return company, nil
}
