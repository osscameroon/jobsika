package docs

// swagger:route GET /jobtitles jobtitles idOfCompanyWithoutID
// Companies returns the list of jobtitles
// responses:
//   200: jobtitlesResponse
//   400: badRequestResponse
//   404: notFoundResponse

// This text will appear as description of your response body.
// swagger:response jobtitlesResponse
type JobtitlesResponseWrapper struct {
	// in:body
	Body []string
}

// swagger:response badRequestResponse
type JobtitlesBadRequestResponseWrapper struct {
	// in:body
	Body struct {
		// Example: bad request
		Error string `json:"error"`
	}
}

// swagger:response notFoundResponse
type JobtitlesNotFoundResponseWrapper struct {
	// in:body
	Body struct {
		// Example: could not found
		Error string `json:"error"`
	}
}
