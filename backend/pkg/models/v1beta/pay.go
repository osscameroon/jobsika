package v1beta

import (
	"errors"
	"strings"
)

type PayPostQuery struct {
	// The email of the user making the payment
	Email string `json:"email"`
	Tier  string `json:"tier"`
}

// Validate check if the mandatory fields are filled
func (p PayPostQuery) Validate() error {
	if strings.TrimSpace(p.Email) == "" {
		return errors.New("email field is mandatory")
	}

	if !IsEmailValid(p.Email) {
		return errors.New("email field is invalid")
	}

	if strings.TrimSpace(p.Tier) == "" {
		return errors.New("tier is mandatory")
	}

	return nil
}
