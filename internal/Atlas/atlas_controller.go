package atlas

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func AtlasController() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := c.Value("requestLogger").(*log.Entry)

		logger.Infoln("Start Atlas Controller")

		atlasResponse, errAtlasResponse := GetCapitalInfo(logger)

		configVariable := os.Getenv("MY_CONFIG")
		log.Infof("Config Variable: %s", configVariable)

		if errAtlasResponse != nil {
			logger.Errorf("Something went wrong in the request: %s", errAtlasResponse.Error())
			errNet := AtlasResponse{
				StatusCode:  http.StatusBadRequest,
				Description: errAtlasResponse.Error(),
				CountryList: []Country{},
			}

			c.JSON(errNet.StatusCode, errNet)
			return
		}

		logger.Infoln("End Atlas Controller")

		c.JSON(atlasResponse.StatusCode, atlasResponse)
	}
}
