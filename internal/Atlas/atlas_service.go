package atlas

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func GetCapitalInfo(logger *log.Entry) (AtlasResponse, error) {

	logger.Infoln("Start - GetCapitalInfo")

	countryList := []Country{
		{Name: "Philippines", Capital: "Manila City"},
		{Name: "Indonesia", Capital: "Jakarta City"},
		{Name: "Vietnam", Capital: "Hanoi City"},
		{Name: "Thailand", Capital: "Bangkok City"},
		{Name: "Malaysia", Capital: "Kuala Lumpur City"},
		{Name: "Singapore", Capital: "Singapore City"},
	}

	atlasResponse := AtlasResponse{
		StatusCode:  http.StatusOK,
		Description: "Request is completed",
		CountryList: countryList,
	}

	logger.Infoln("End - GetCapitalInfo")

	return atlasResponse, nil
}
