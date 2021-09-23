package netflix

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetNetflixShow() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := c.Value("requestLogger").(*log.Entry)

		logger.Infoln("Welcome to this sample Netflix show!")

		name := c.Request.URL.Query().Get("name")

		logger.Infof("Hello to the app: %s", name)

		errNetflix, netflixResponse := GetFavoriteShow(name)

		sampleVariable := os.Getenv("SAMPLE_VARIABLE")
		secretVariable := os.Getenv("MY_SECRET")
		configVariable := os.Getenv("MY_CONFIG")

		log.Infof("Sample Variabl: %s, Secret Variable: %s, Config Variable: %s", sampleVariable, secretVariable, configVariable)

		if errNetflix != nil {
			logger.Errorf("Something went wrong in the request: %s", errNetflix.Error())
			errNet := NetflixResponse{
				StatusCode:  http.StatusBadRequest,
				Description: errNetflix.Error(),
				NetflixShow: NetflixShow{},
				Person:      Person{},
			}

			c.JSON(errNet.StatusCode, errNet)
			return
		}

		c.JSON(netflixResponse.StatusCode, netflixResponse)
		return
	}
}
