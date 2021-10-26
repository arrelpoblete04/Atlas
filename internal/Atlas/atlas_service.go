package atlas

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func GetCapitalInfo(logger *log.Entry) (AtlasResponse, error) {

	logger.Infoln("Start - GetCapitalInfo")

	countryList := []Country{
		{Name: "Philippines", Capital: "Manila City"},
		{Name: "Indonesia", Capital: "Jakarta"},
		{Name: "Vietnam", Capital: "Hanoi"},
		{Name: "Thailand", Capital: "Bangkok"},
		{Name: "Malaysia", Capital: "Kuala Lumpur"},
	}

	atlasResponse := AtlasResponse{
		StatusCode:  http.StatusOK,
		Description: "Request is completed",
		CountryList: countryList,
	}

	logger.Infoln("End - GetCapitalInfo")

	return atlasResponse, nil
}
