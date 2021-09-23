package netflix

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	log "github.com/sirupsen/logrus"
)

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

func TestGetNetflixShow(t *testing.T) {

	scenarios := []struct {
		testCase   string
		name       string
		statusCode int
	}{
		{testCase: "Valid Name", name: "Arrel", statusCode: http.StatusOK},
		{testCase: "Blank Name", name: "", statusCode: http.StatusBadRequest},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.testCase, func(t *testing.T) {

			urlPath := "/netflix"
			request, errorNewRequest := http.NewRequest(http.MethodGet, urlPath, nil)

			if errorNewRequest != nil {
				t.Errorf("Test Failed. There's an error in wrapping the request. %s", errorNewRequest.Error())
			}

			q := request.URL.Query()
			q.Add("name", scenario.name)
			request.URL.RawQuery = q.Encode()

			r := gin.Default()
			r.Use(appendRequestIdLogging())
			rr := httptest.NewRecorder()
			r.GET(urlPath, GetNetflixShow())
			r.ServeHTTP(rr, request)

			t.Log(rr.Body)

			var response NetflixResponse
			err := json.Unmarshal(rr.Body.Bytes(), &response)
			if err != nil {
				t.Error("Error while unmarshalling response")
			}

			if response.StatusCode != scenario.statusCode {
				t.Errorf("NetflixResponse status code: %v, Expected: %v", response.StatusCode,
					scenario.statusCode)
			}

		})
	}
}
