package election

import (
	"errors"
	"net/http"
)

func GetCandidateInfo(position string) (error, ElectionResponse) {

	if position == "" {
		return errors.New("Position should be included"), ElectionResponse{}
	}

	candidate := Candidate{
		Name:     "Leni Robredo",
		Position: "President",
	}

	candidateList := []Candidate{candidate}

	electionResponse := ElectionResponse{
		StatusCode:  http.StatusOK,
		Description: "Request is completed",
		Candidate:   candidateList,
	}

	return nil, electionResponse
}
