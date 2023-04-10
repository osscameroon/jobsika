package v1beta

import (
	"errors"
	"strings"
	"time"
)

type Subscribers struct {
	ID        int64     `json:"id" gorm:"autoIncrement:true"`
	Email     string    `json:"email" gorm:"primaryKey;column:email"`
	CreatedAt time.Time `json:"createdat" gorm:"column:createdat"`
	UpdatedAt time.Time `json:"updatedat" gorm:"column:updatedat"`
}

type SubscribersPostQuery struct {
	Email string `json:"email"`
}

func (query *SubscribersPostQuery) Validate() error {
	if !IsEmailValid(query.Email) {
		return errors.New("email is not a valid email address")
	}

	query.Email = strings.ToLower(query.Email)

	return nil
}
