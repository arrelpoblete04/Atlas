package election

import (
	"errors"
	"net/http"
)

func GetCandidateInfo(position string) (error, ElectionResponse) {

	if position == "" {
		return errors.New("Position should be included"), ElectionResponse{}
	}

	candidateList := []Candidate{
		{Name: "Leni Robredo", Position: "President"},
		{Name: "Isko Moreno", Position: "President"},
		{Name: "Manny Pacquiao", Position: "President"},
	}

	electionResponse := ElectionResponse{
		StatusCode:  http.StatusOK,
		Description: "Request is completed",
		Candidate:   candidateList,
	}

	return nil, electionResponse
}
