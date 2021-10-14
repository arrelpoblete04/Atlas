package election

import (
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func GetCandidateInfo(logger *log.Entry, position string) (ElectionResponse, error) {

	logger.Infoln("Start - GetCandidateInfo")

	if position == "" {
		return ElectionResponse{}, errors.New("Parameter 'Position' should not be empty.")
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

	logger.Infoln("End - GetCandidateInfo")

	return electionResponse, nil
}
