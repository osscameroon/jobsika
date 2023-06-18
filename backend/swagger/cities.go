package swagger

// swagger:route GET /cities cities idOfCities
// Companies returns cities endpoint
// responses:
//   200: citiesResponse

// This text will appear as description of your response body.
// swagger:response citiesResponse
type CitiesResponseWrapper struct {
	// in:body
	//Example: ["Douala","Yaoundé","Garoua","Bamenda","Maroua","Nkongsamba","Bafoussam","Ngaoundéré"]
	Body []string
}
