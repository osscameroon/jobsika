package docs

// swagger:route GET /seniority seniority idOfSeniority
// Companies returns seniority endpoint
// responses:
//   200: seniorityResponse

// This text will appear as description of your response body.
// swagger:response seniorityResponse
type SeniorityResponseWrapper struct {
	// in:body
	//Example: ["Entry-level","Mid-level","Senior","Above senior-level","Executive"]
	Body []string
}
