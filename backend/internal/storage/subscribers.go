package storage

import (
	"github.com/osscameroon/jobsika/pkg/models/v1beta"
)

const JobSubscribersTableName = "job_subscribers"

func (db DB) PostSubscribers(query v1beta.SubscribersPostQuery) (*v1beta.Subscribers, error) {
	subscriber := new(v1beta.Subscribers)

	//Check if the subscriber email already exist
	err := db.c.Table(JobSubscribersTableName).Where("email = ?", query.Email).First(subscriber).Error
	if err == nil {
		return nil, err
	}

	subscriber.Email = query.Email

	err = db.c.Table(JobSubscribersTableName).Create(subscriber).Error
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return subscriber, nil
}
