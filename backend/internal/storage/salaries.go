package storage

import "github.com/osscameroon/jobsika/pkg/models/v1beta"

// GetSalaries get salaries
func (db DB) GetSalaries() ([]v1beta.Salary, error) {
	salaries := []v1beta.Salary{}
	ret := db.c.Find(&salaries)
	if ret.Error != nil {
		return salaries, ret.Error
	}

	return salaries, nil
}

// GetSalaryByID get salary by `id`
func (db DB) GetSalaryByID(id int64) (v1beta.Salary, error) {
	salary := v1beta.Salary{}
	ret := db.c.First(&salary, "id = ?", id)
	if ret.Error != nil {
		return salary, ret.Error
	}

	return salary, nil
}
