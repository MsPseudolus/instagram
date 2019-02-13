package instagram_test

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"testing"

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
	client       *instagram.Api
)

func TestMain(m *testing.M) {

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

	exit := m.Run()
	os.Exit(exit)
}

func newAPI() *instagram.Api {
	return instagram.New(clientID, clientSecret, accessToken, true)
}

func TestVerifyCredentials(t *testing.T) {
	api := newAPI()

	if ok, err := api.VerifyCredentials(); !ok {
		t.Error(err)
	}
}

func TestSelf(t *testing.T) {
	api := newAPI()

	self, err := api.GetSelf()
	checkRes(t, self.Meta, err)
}

func TestGetRecentMedia(t *testing.T) {
	api := newAPI()

	params := url.Values{}
	params.Set("count", "3")
	res, err := api.GetRecentMedia(params)
	checkRes(t, res.Meta, err)

	if len(res.Medias) != 3 {
		t.Error("Count didn't apply")
	}

	//nextRes, err := api.NextMedias(res.Pagination)
	//checkRes(t, nextRes.Meta, err)

	//if nextRes.Pagination.Pagination != nil {
	//t.Error("Pagination should be not valid!", nextRes.Pagination.Pagination)
	//}

	//nextNextRes, err := api.NextMedias(nextRes.Pagination)
	//if len(nextNextRes.Medias) > 0 {
	//t.Error("Pagination returned non-nil next request after nil pagination!")
	//}
}

func TestGetMediaRecentComments(t *testing.T) {
	api := newAPI()

	mediaID, err := findMediaWithComments()
	if err != nil {
		t.Fatalf("get media id: %s", err)
	}

	res, err := api.GetMediaRecentComments(mediaID)
	checkRes(t, res.Meta, err)

	if len(res.Comments) == 0 {
		t.Error("too few comments!", len(res.Comments))
	}
}

// -- helpers --

func findMediaWithComments() (string, error) {
	api := newAPI()

	params := url.Values{}
	params.Set("count", "20")

	res, err := api.GetRecentMedia(params)
	if err != nil {
		return "", err
	}
	for _, m := range res.Medias {
		if m.Comments != nil && m.Comments.Count > 0 {
			log.Printf("M C %d", m.Comments.Count)
			return m.Id, nil
		}
	}
	return "", errors.New("No media with comments")
}

func checkRes(t *testing.T, m *instagram.Meta, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
	if m == nil || m.Code != 200 {
		t.Error("Meta not right", m)
	}
}
