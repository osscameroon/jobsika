package storage

//GetJobTitles get jobtitles
func (db DB) GetJobTitles() ([]string, error) {
	jobTitles := []string{}
	ret := db.c.Table("salaries").Select("title").Distinct("title").Order("title").Find(&jobTitles)
	if ret.Error != nil {
		return jobTitles, ret.Error
	}

	return jobTitles, nil
}
