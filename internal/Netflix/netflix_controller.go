package netflix

import (
	"net/http"

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
