package election

type ElectionResponse struct {
	StatusCode  int         `json:"statusCode"`
	Description string      `json:"description"`
	Candidate   []Candidate `json:"candidate"`
}

type Candidate struct {
	Name     string
	Position string
}
