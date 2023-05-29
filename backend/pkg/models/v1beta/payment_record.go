package v1beta

import "time"

type PaymentRecord struct {
	Id         int64     `json:"id" gorm:"column:id;auto_increment;primary_key"`
	Email      string    `json:"email" gorm:"column:email"`
	TierId     string    `json:"tier_id" gorm:"column:tier_id"`
	JobOfferID int64     `json:"job_offer_id" gorm:"column:job_offer_id"`
	LegacyId   int64     `json:"legacy_id" gorm:"column:legacy_id"`
	Slug       string    `json:"slug" gorm:"column:slug"`
	TierUrl    string    `json:"tier_url" gorm:"column:tier_url"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:createdat"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updatedat"`
}
