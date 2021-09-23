package netflix

import (
	"errors"
	"net/http"
)

func GetFavoriteShow(name string) (error, NetflixResponse) {

	if name == "" {
		return errors.New("Name should be included"), NetflixResponse{}
	}

	netflixShow := NetflixShow{
		Show:  "Squid Game",
		Genre: "Thriller",
	}

	person := Person{
		Name: name,
	}

	netflixResponse := NetflixResponse{
		StatusCode:  http.StatusOK,
		Description: "Request is completed",
		NetflixShow: netflixShow,
		Person:      person,
	}

	return nil, netflixResponse
}
