package storage

import "github.com/osscameroon/jobsika/pkg/models/v1beta"

func (db *DB) CreatePaymentRecord(record *v1beta.PaymentRecord) error {
	return db.c.Create(record).Error
}
