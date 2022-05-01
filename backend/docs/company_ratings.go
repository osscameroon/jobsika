package docs

import (
	"github.com/elhmn/jobsika/pkg/models/v1beta"
)

// swagger:route GET /company-ratings company-ratings idOfCompanyRatingsWithoutID
// CompanyRatings returns the list of company-ratings
// responses:
//   200: company-ratingsResponse
//   400: badRequestResponse
//   404: notFoundResponse

// swagger:route GET /company-ratings/{id} company-ratings idOfCompanyRatings
// CompanyRatings returns the list of company-ratings
// responses:
//   200: companyRatingResponse
//   400: badRequestResponse
//   404: notFoundResponse

// This text will appear as description of your response body.
// swagger:response company-ratingsResponse
type CompanyRatingsResponseWrapper struct {
	// in:body
	Body []v1beta.Company
}

// swagger:response badRequestResponse
type CompanyRatingsBadRequestResponseWrapper struct {
	// in:body
	Body struct {
		// Example: bad request
		Error string `json:"error"`
	}
}

// swagger:response notFoundResponse
type CompanyRatingsNotFoundResponseWrapper struct {
	// in:body
	Body struct {
		// Example: could not found
		Error string `json:"error"`
	}
}

// swagger:parameters idOfCompanyRatingsWithoutID
type CompanyRatingsQueryParam struct {
	//in:query
	//unique: true
	CompanyID int `json:"company_id"`
}

// swagger:parameters idOfCompanyRatings
type CompanyRatingsParam struct {
	//in:path
	//Example: 1
	ID string `json:"id"`
}
