package main

import (
	"fmt"
	"net/http"
	"os"

	election "app.election/internal/Election"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	log "github.com/sirupsen/logrus"
)

func main() {

	// db_user := url.QueryEscape(os.Getenv("DB_USERNAME"))
	// db_password := url.QueryEscape(os.Getenv("DB_PASSWORD"))
	// s := database.DB(db_user, db_password)

	// errorPing := s.Ping()

	// if errorPing != nil {
	// 	log.Errorf("Error in Ping")
	// } else {
	// 	log.Infof("No Error in Ping")
	// }

	// log.Info("aa" + s.Ping().Error())
	router := gin.Default()

	router.Use(appendRequestIdLogging())
	router.GET("/election2022", election.ElectionController())

	log.Info("ArgoCD Tekton Github - PR2")
	log.Info("Serving readiness probe at port: 8000")
	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router)

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
