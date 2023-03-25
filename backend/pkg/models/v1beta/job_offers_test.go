package v1beta

import (
	"errors"
	"testing"
)

func TestValidateJobOffer(t *testing.T) {
	testCases := []struct {
		name          string
		jobOffer      OfferPostQuery
		expectedError error
	}{
		{"invalid company email", OfferPostQuery{
			CompanyName:             "test",
			CompanyEmail:            "test@osscameroon",
			JobTitle:                "test",
			IsRemote:                true,
			City:                    "test",
			Country:                 "test",
			Department:              "test",
			SalaryRangeMin:          1,
			SalaryRangeMax:          10,
			Description:             "test",
			Benefits:                "test",
			HowToApply:              "test",
			ApplicationUrl:          "test",
			ApplicationEmailAddress: "test@osscameroon.com",
			ApplicationPhoneNumber:  "+49123456789",
			Tags:                    "test, test",
		}, errors.New("company email is not a valid email address")},

		{"invalid application email", OfferPostQuery{
			CompanyName:             "test",
			CompanyEmail:            "test@osscameroon.com",
			JobTitle:                "test",
			IsRemote:                true,
			City:                    "test",
			Country:                 "test",
			Department:              "test",
			SalaryRangeMin:          1,
			SalaryRangeMax:          10,
			Description:             "test",
			Benefits:                "test",
			HowToApply:              "test",
			ApplicationUrl:          "test",
			ApplicationEmailAddress: "test@osscameroon",
			ApplicationPhoneNumber:  "+49123456789",
			Tags:                    "test, test",
		}, errors.New("application email address is not a valid email address")},

		{"invalid job offer", OfferPostQuery{
			CompanyName:             "test",
			CompanyEmail:            "test@osscameroon.com",
			JobTitle:                "test",
			IsRemote:                true,
			City:                    "test",
			Country:                 "test",
			Department:              "test",
			SalaryRangeMin:          1,
			SalaryRangeMax:          10,
			Description:             "test",
			Benefits:                "test",
			HowToApply:              "test",
			ApplicationUrl:          "test",
			ApplicationEmailAddress: "test@osscameroon.com",
			ApplicationPhoneNumber:  "I am not a phone number",
			Tags:                    "test, test",
		}, errors.New("application phone number is not a valid phone number")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.jobOffer.Validate()
			if err == nil || err != nil && tc.expectedError.Error() != err.Error() {
				t.Errorf("expected error [%v], got [%v]", tc.expectedError, err)
			}
		})
	}
}
