package storage

import "github.com/osscameroon/jobsika/pkg/models/v1beta"

func (db *DB) CreatePaymentRecord(record *v1beta.PaymentRecord) error {
	return db.c.Create(record).Error
}

// GetPaymentRecordByID get salary by `id`
func (db DB) GetPaymentRecordByID(id int64) (v1beta.PaymentRecord, error) {
	paymentRecord := v1beta.PaymentRecord{}
	ret := db.c.First(&paymentRecord, "legacy_id = ?", id)
	if ret.Error != nil {
		return paymentRecord, ret.Error
	}

	return paymentRecord, nil
}
