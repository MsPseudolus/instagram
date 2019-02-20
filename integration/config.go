package integration

import (
	"fmt"
	"os"

	"github.com/recentralized/instagram"
)

const (
	envClientID     = "TEST_INSTAGRAM_CLIENT_ID"
	envClientSecret = "TEST_INSTAGRAM_CLIENT_SECRET"
	envAccessToken  = "TEST_INSTAGRAM_ACCESS_TOKEN"
)

var (
	clientID     string
	clientSecret string
	accessToken  string
)

// NewAPI initializes an API configured with environment variables.
func NewAPI() *instagram.API {
	loadConfig()
	return instagram.New(clientID, clientSecret, accessToken, true)
}

func loadConfig() {
	if clientID != "" {
		return
	}

	clientID = os.Getenv(envClientID)
	clientSecret = os.Getenv(envClientSecret)
	accessToken = os.Getenv(envAccessToken)

	if clientID == "" {
		fmt.Printf("Missing %s\n", envClientID)
		os.Exit(1)
	}
	if clientSecret == "" {
		fmt.Printf("Missing %s\n", envClientSecret)
		os.Exit(1)
	}
	if accessToken == "" {
		fmt.Printf("Missing %s\n", envAccessToken)
		os.Exit(1)
	}
}
