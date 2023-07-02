package swagger

import "github.com/osscameroon/jobsika/pkg/models/v1beta"

// swagger:route GET /jobs jobs idOfJobWithoutID
// Jobs returns the list of jobs
// responses:
//   200: jobsResponse
//   400: badRequestResponse
//   404: notFoundResponse

// swagger:route POST /jobs jobs idOfJobPost
// Jobs create a new job item
// responses:
//   201: postJobsResponse
//   400: badRequestResponse

// swagger:route GET /jobs/{id}/image jobs idOfJobImage
// Jobs return the image of the company posting a job
// responses:
//   200: jobImageResponse
//   400: badRequestResponse
//   404: notFoundResponse

// This text will appear as description of your response body.
// swagger:response jobsResponse
type JobsResponseWrapper struct {
	// in:body
	Body v1beta.JobOffersResponse
}

// swagger:response postJobsResponse
type PostJobsResponseWrapper struct {
	// in:body
	Body v1beta.JobOfferPresenter
}

// swagger:parameters idOfJobWithoutID
type JobsParams struct {
	// in:query
	//example: 1
	Page string `json:"page"`

	//in:query
	//example: 20
	Limit string `json:"limit"`

	//in:query
	//example: Jaxbean
	Company string `json:"company"`

	//in:query
	//example: Recruiting Manager
	JobTitle string `json:"jobtitle"`

	//in:query
	//example: true
	IsRemote string `json:"isremote"`

	//in:query
	//example: Douala
	City string `json:"city"`

	//in:query
	//example: Cameroon
	Country string `json:"country"`
}

// swagger:parameters idOfJobPost
type PostJobsParams struct {
	// in:body
	Body v1beta.OfferPostQuery
}

// swagger:parameters idOfJobImage
type JobImageParams struct {
	// in:path
	//example: 1
	ID string `json:"id"`
}

// This text will appear as description of your response body.
// swagger:response jobImageResponse
type JobImageResponseWrapper struct {
	// in:body
	Body []byte
}

// swagger:response badRequestResponse
type JobsBadRequestResponseWrapper struct {
	// in:body
	Body struct {
		// Example: bad request
		Error string `json:"error"`
	}
}

// swagger:response notFoundResponse
type JobsNotFoundResponseWrapper struct {
	// in:body
	Body struct {
		// Example: could not found
		Error string `json:"error"`
	}
}
