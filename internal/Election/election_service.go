package election

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func GetCandidateInfo(logger *log.Entry) (ElectionResponse, error) {

	logger.Infoln("Start - GetCandidateInfo")

	candidateList := []Official{
		{Name: "President", Level: "National"},
		{Name: "Vice President", Level: "National"},
		{Name: "Senator", Level: "National"},
		{Name: "Governor", Level: "Local"},
	}

	electionResponse := ElectionResponse{
		StatusCode:  http.StatusOK,
		Description: "Request is completed",
		Candidate:   candidateList,
	}

	logger.Infoln("End - GetCandidateInfo")

	return electionResponse, nil
}
