package main

import (
	"fmt"
	"net/http"
	"os"

	atlas "app.atlas/internal/Atlas"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	log "github.com/sirupsen/logrus"
)

func main() {

	router := gin.Default()

	router.Use(appendRequestIdLogging())
	router.GET("/atlas", atlas.AtlasController())

	port := os.Getenv("PORT")

	log.Infof("Serving readiness probe at port: %s", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router)

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
