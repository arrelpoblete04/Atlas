package election

import (
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func GetCandidateInfo(logger *log.Entry, position string) (ElectionResponse, error) {

	logger.Infoln("Start - GetCandidateInfo")

	if position == "" {
		return ElectionResponse{}, errors.New("Position should be included.")
	}

	candidateList := []Candidate{
		{Name: "Leni Robredo", Position: "President"},
		{Name: "Isko Moreno", Position: "President"},
		{Name: "Manny Pacquiao", Position: "President"},
		{Name: "Panfilo Lacson", Position: "President"},
	}

	electionResponse := ElectionResponse{
		StatusCode:  http.StatusOK,
		Description: "Request is completed",
		Candidate:   candidateList,
	}

	logger.Infoln("End - GetCandidateInfo")

	return electionResponse, nil
}
