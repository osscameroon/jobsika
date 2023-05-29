package swagger

// swagger:route GET /health health idOfHealth
// Companies returns health endpoint
// responses:
//   200: healthResponse

// This text will appear as description of your response body.
// swagger:response healthResponse
type HealthResponseWrapper struct {
	// in:body
	Body struct {
		//Example: all good !
		Message string `json:"message"`
	}
}
