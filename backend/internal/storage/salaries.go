package storage

import "github.com/elhmn/camerdevs/pkg/models/v1beta"

//GetSalaries get salaries
func (db DB) GetSalaries() ([]v1beta.Salary, error) {
	salaries := []v1beta.Salary{}
	ret := db.c.Find(&salaries)
	if ret.Error != nil {
		return salaries, ret.Error
	}

	return salaries, nil
}
