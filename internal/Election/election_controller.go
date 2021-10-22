package election

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func ElectionController() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := c.Value("requestLogger").(*log.Entry)

		logger.Infoln("Start Election Controller")

		candidateResponse, errCandidateResp := GetCandidateInfo(logger)

		configVariable := os.Getenv("MY_CONFIG")
		log.Infof("Config Variable: %s", configVariable)

		if errCandidateResp != nil {
			logger.Errorf("Something went wrong in the request: %s", errCandidateResp.Error())
			errNet := ElectionResponse{
				StatusCode:  http.StatusBadRequest,
				Description: errCandidateResp.Error(),
				Candidate:   []Official{},
			}

			c.JSON(errNet.StatusCode, errNet)
			return
		}

		logger.Infoln("End Election Controller")

		c.JSON(candidateResponse.StatusCode, candidateResponse)
		return
	}
}
