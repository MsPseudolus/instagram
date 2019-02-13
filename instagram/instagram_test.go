package instagram

import (
	"fmt"
	"net/url"
	"testing"
)

var DoAuthorizedRequests bool
var api *Api
var ccistulli_id string = "401243155"
var ladygaga_id string = "184692323"

func init() {
	DoAuthorizedRequests = (TestConfig["access_token"] != "")
	if !DoAuthorizedRequests {
		fmt.Println("*** Authorized requests will not performed because no access_token was specified in config_test.go")
	}
	api = createApi()
}

func TestVerifyCredentials(t *testing.T) {
	authorizedRequest(t)

	if ok, err := api.VerifyCredentials(); !ok {
		t.Error(err)
	}
}

func TestSelf(t *testing.T) {
	authorizedRequest(t)

	self, err := api.GetSelf()
	checkRes(t, self.Meta, err)
}

func TestGetRecentMedia(t *testing.T) {
	params := url.Values{}
	params.Set("count", "3") // 4 images in this set
	params.Set("max_timestamp", "1466809870")
	params.Set("min_timestamp", "1396751898")
	res, err := api.GetRecentMedia(params)
	checkRes(t, res.Meta, err)

	if len(res.Medias) != 3 {
		t.Error("Count didn't apply")
	}

	nextRes, err := api.NextMedias(res.Pagination)
	checkRes(t, nextRes.Meta, err)

	if len(nextRes.Medias) != 1 {
		t.Error("Timestamps didn't apply")
	}

	if nextRes.Pagination.Pagination != nil {
		t.Error("Pagination should be not valid!", nextRes.Pagination.Pagination)
	}

	nextNextRes, err := api.NextMedias(nextRes.Pagination)
	if len(nextNextRes.Medias) > 0 {
		t.Error("Pagination returned non-nil next request after nil pagination!")
	}
}

func TestGetMediaRecentComments(t *testing.T) {
	res, err := api.GetMediaRecentComments("594914758412103315_2134762")
	checkRes(t, res.Meta, err)

	if len(res.Comments) < 10 {
		t.Error("too few comments!", len(res.Comments))
	}
}

// -- helpers --

func authorizedRequest(t *testing.T) {
	if !DoAuthorizedRequests {
		t.Skip("Access Token not provided.")
	}
}

func checkRes(t *testing.T, m *Meta, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
	if m == nil || m.Code != 200 {
		t.Error("Meta not right", m)
	}
}

func values(keyValues ...string) url.Values {
	v := url.Values{}
	for i := 0; i < len(keyValues)-1; i += 2 {
		v.Set(keyValues[i], keyValues[i+1])
	}
	return v
}

func createApi() *Api {
	return New(TestConfig["client_id"], TestConfig["client_secret"], TestConfig["access_token"], true)
}
