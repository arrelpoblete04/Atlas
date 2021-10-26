package atlas

type AtlasResponse struct {
	StatusCode  int       `json:"statusCode"`
	Description string    `json:"description"`
	CountryList []Country `json:"countries"`
}

type Country struct {
	Name    string
	Capital string
}
