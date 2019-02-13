package instagram

import (
	"net/url"
	"testing"
)

func TestIterate_GetUserRecentMedia(t *testing.T) {
	params := url.Values{}
	params.Set("count", "2") // 5 images in this set. Get them 2 at time
	params.Set("max_timestamp", "1384161094")
	params.Set("min_timestamp", "1382656250")
	res, err := api.GetRecentMedia(params)
	checkRes(t, res.Meta, err)

	mediaChan, errChan := api.IterateMedia(res, nil)
	for media := range mediaChan {
		if media.User.Username != "ladygaga" {
			t.Error("Got a media with wrong username?", media.User)
		}
	}
	if err := <-errChan; err != nil {
		t.Error(err)
	}
}
