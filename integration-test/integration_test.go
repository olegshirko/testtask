package integration_test

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	. "github.com/Eun/go-hit"
)

const (
	// Attempts connection
	host       = "app:8080"
	healthPath = "https://" + host + "/health"
	attempts   = 20

	// HTTP REST
	basePath = "https://" + host + "/api"
)

var (
	// Define the base client that will be used for making HTTP requests in tests
	client *http.Client
)

func init() {
	// Setup the custom HTTP client
	client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // Disable SSL certificate verification
			},
		},
	}
}

func TestMain(m *testing.M) {
	err := healthCheck(attempts)
	if err != nil {
		log.Fatalf("Integration tests: host %s is not available: %s", host, err)
	}

	log.Printf("Integration tests: host %s is available", host)

	code := m.Run()
	os.Exit(code)
}

func healthCheck(attempts int) error {
	var err error

	for attempts > 0 {
		err = Do(Get(healthPath),
			HTTPClient(client),
			Expect().Status().Equal(http.StatusOK))
		if err == nil {
			return nil
		}

		log.Printf("Integration tests: url %s is not available, attempts left: %d", healthPath, attempts)

		time.Sleep(time.Second)

		attempts--
	}

	return err
}

// HTTP GET: /api/auth.
func TestAuth(t *testing.T) {
	body := `{
 		"login": "alice",
 		"password": "secret"
	}`
	Test(t,
		Description("Auth Success"),
		HTTPClient(client),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Post(basePath+"/auth"),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().String().Contains(`{"token":"`),
	)
}

// HTTP GET: /api/history.
func TestHTTPHistory(t *testing.T) {
	Test(t,
		Description("History Unauthorized"),
		HTTPClient(client),
		Send().Headers("Authorization").Add("Bearer 0234"),
		Get(basePath+"/history"),
		Expect().Status().Equal(http.StatusUnauthorized),
		Expect().Body().String().Contains(`{"error":"Unauthorized: No valid bearer token provided"`),
	)
}
