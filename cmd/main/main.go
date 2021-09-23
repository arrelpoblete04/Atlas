package main

import (
	"net/http"

	netflix "app.netflix/internal/Netflix"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	log "github.com/sirupsen/logrus"
)

func main() {

	router := gin.Default()

	router.Use(appendRequestIdLogging())
	router.GET("/netflix", netflix.GetNetflixShow())

	log.Info("Serving readiness probe at port: 8000")
	err := http.ListenAndServe(":8000", router)

	if err != nil {
		log.Fatalf("Error during readiness probe startup: %v", err)
	}

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
