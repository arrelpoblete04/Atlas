package main

import (
	netflix "app.netflix/internal/Netflix"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	log "github.com/sirupsen/logrus"
)

func main() {

	router := gin.Default()

	router.Use(appendRequestIdLogging())
	router.GET("/netflix", netflix.GetNetflixShow())

	router.Run(":8000")

}

func appendRequestIdLogging() gin.HandlerFunc {
	log.Println("Appending x-request-id")
	return func(c *gin.Context) {
		id := c.Request.Header.Get("x-request-id")
		if id == "" {
			log.Println("Creating new x-request-id")
			id = uuid.New().String()
		}

		entry := log.WithField("x-request-id", id)
		c.Set("requestLogger", entry)
	}
}
