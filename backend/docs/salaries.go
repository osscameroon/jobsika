package docs

import (
	"github.com/elhmn/camerdevs/pkg/models/v1beta"
)

// swagger:route GET /salaries salaries idOfSalaryWithoutID
// Salaries returns the list of salaries
// responses:
//   200: salariesResponse
//   400: badRequestResponse
//   404: notFoundResponse

// swagger:route GET /salaries/{id} salaries idOfSalary
// Salaries returns the list of salaries
// responses:
//   200: salariesResponse
//   400: badRequestResponse
//   404: notFoundResponse

// This text will appear as description of your response body.
// swagger:response salariesResponse
type SalariesResponseWrapper struct {
	// in:body
	Body []v1beta.Salary
}

// swagger:response badRequestResponse
type SalariesBadRequestResponseWrapper struct {
	// in:body
	Body struct {
		// Example: bad request
		Error string `json:"error"`
	}
}

// swagger:response notFoundResponse
type SalariesNotFoundResponseWrapper struct {
	// in:body
	Body struct {
		// Example: could not found
		Error string `json:"error"`
	}
}

// swagger:parameters idOfSalary
type SalaryParam struct {
	//in:path
	//example: 1
	ID string `json:"id"`
}
