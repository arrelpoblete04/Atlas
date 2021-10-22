package election

type ElectionResponse struct {
	StatusCode  int        `json:"statusCode"`
	Description string     `json:"description"`
	Candidate   []Official `json:"candidate"`
}

type Official struct {
	Name  string
	Level string
}
