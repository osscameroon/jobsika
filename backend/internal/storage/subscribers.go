package storage

import (
	"github.com/elhmn/jobsika/pkg/models/v1beta"
	"gorm.io/gorm"
)

const JobSubscribersTableName = "job_subscribers"

func (db DB) PostSubscribers(query v1beta.SubscribersPostQuery) (*v1beta.Subscribers, error) {
	subscriber := new(v1beta.Subscribers)

	err := db.c.Transaction(func(tx *gorm.DB) error {
		//Check if the subscriber email already exist
		err := tx.Table(JobSubscribersTableName).Where("email = ?", query.Email).First(subscriber).Error
		if err == nil {
			return nil
		}

		subscriber.Email = query.Email

		err = tx.Table(JobSubscribersTableName).Create(subscriber).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return subscriber, nil
}
