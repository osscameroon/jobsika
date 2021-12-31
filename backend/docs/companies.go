package docs

import (
	"github.com/elhmn/camerdevs/pkg/models/v1beta"
)

// swagger:route GET /companies companies idOfCompanyWithoutID
// Companies returns the list of companies
// responses:
//   200: companiesResponse
//   400: badRequestResponse
//   404: notFoundResponse

// swagger:route GET /companies/{id} companies idOfCompany
// Companies returns the list of companies
// responses:
//   200: companiesResponse
//   400: badRequestResponse
//   404: notFoundResponse

// This text will appear as description of your response body.
// swagger:response companiesResponse
type CompaniesResponseWrapper struct {
	// in:body
	Body []v1beta.Company
}

// swagger:response badRequestResponse
type CompaniesBadRequestResponseWrapper struct {
	// in:body
	Body struct {
		// Example: bad request
		Error string `json:"error"`
	}
}

// swagger:response notFoundResponse
type CompaniesNotFoundResponseWrapper struct {
	// in:body
	Body struct {
		// Example: could not found
		Error string `json:"error"`
	}
}

// swagger:parameters idOfCompany
type CompanyParam struct {
	//in:path
	//example: 1
	ID string `json:"id"`
}
