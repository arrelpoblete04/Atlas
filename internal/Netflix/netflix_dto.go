package netflix

type NetflixResponse struct {
	StatusCode  int         `json:"statusCode"`
	Description string      `json:"description"`
	NetflixShow NetflixShow `json:"netflixShow"`
	Person      Person      `json:"person"`
}

type NetflixShow struct {
	Show  string
	Genre string
}

type Person struct {
	Name string
}
