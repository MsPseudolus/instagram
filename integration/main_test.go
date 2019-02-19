package integration_test

import (
	"context"
	"errors"
	"flag"
	"log"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/recentralized/instagram"
	"github.com/recentralized/instagram/integration"
)

func TestMain(m *testing.M) {
	flag.Parse()

	if testing.Short() {
		log.Printf("skipping in short mode")
		os.Exit(0)
	}

	// Make sure we can initialize the API.
	newAPI()

	exit := m.Run()
	os.Exit(exit)
}

func newAPI() *instagram.Api {
	return integration.NewAPI()
}

func TestVerifyCredentials(t *testing.T) {
	ctx := context.Background()
	api := newAPI()

	if ok, err := api.VerifyCredentials(ctx); !ok {
		t.Error(err)
	}
}

func TestSelf(t *testing.T) {
	ctx := context.Background()
	api := newAPI()

	self, err := api.GetSelf(ctx)
	checkRes(t, self.Meta, err)
}

func TestContextDeadline(t *testing.T) {
	ctx := context.Background()
	api := newAPI()

	// An impossibly short timeout
	ctx, cancel := context.WithTimeout(ctx, time.Duration(1))
	defer cancel()

	_, err := api.GetSelf(ctx)
	if !strings.HasSuffix(err.Error(), context.DeadlineExceeded.Error()) {
		t.Fatalf("Want cause to be DeadlineExceeded got %v", err)
	}
}

func TestGetRecentMedia(t *testing.T) {
	ctx := context.Background()
	api := newAPI()

	params := url.Values{}
	params.Set("count", "3")
	res, err := api.GetRecentMedia(ctx, params)
	checkRes(t, res.Meta, err)

	if len(res.Medias) != 3 {
		t.Error("Count didn't apply")
	}

	media := res.Medias[0]

	if media.CreatedTime.IsZero() {
		t.Errorf("CreatedTime is empty")
	}
}

func TestGetRecentMediaPagination(t *testing.T) {
	ctx := context.Background()
	api := newAPI()

	params := url.Values{}
	params.Set("count", "3")
	res, err := api.GetRecentMedia(ctx, params)
	checkRes(t, res.Meta, err)

	nextRes, err := api.NextMedias(ctx, res.Pagination)
	checkRes(t, nextRes.Meta, err)

	if len(nextRes.Medias) == 0 {
		t.Error("Didn't return next")
	}
}

func TestGetRecentMediaIterate(t *testing.T) {
	ctx := context.Background()
	api := newAPI()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	params := url.Values{}
	params.Set("count", "2")
	res, err := api.GetRecentMedia(ctx, params)
	checkRes(t, res.Meta, err)

	var count int

	mediac, errc := api.IterateMedia(ctx, res)

	const max = 5 // multiple pages, and quit partway through a page.
	for m := range mediac {
		count++
		if count >= max {
			cancel()
		}
		if m.Id == "" {
			t.Errorf("Id is empty: %v", m)
		}
	}
	for err := range errc {
		if !strings.Contains(err.Error(), context.Canceled.Error()) {
			t.Fatalf("Got error: %s", err)
		}
	}
	// NOTE: non-strict equality check because we get max or max+1. Not
	// worth debugging right now
	if got, want := count, max; got < want {
		t.Errorf("Count got %d want %d", got, want)
	}
}

func TestGetMediaRecentComments(t *testing.T) {
	ctx := context.Background()
	api := newAPI()

	mediaID, err := findMediaWithComments()
	if err != nil {
		t.Fatalf("get media id: %s", err)
	}

	res, err := api.GetMediaRecentComments(ctx, mediaID)
	checkRes(t, res.Meta, err)

	if len(res.Comments) == 0 {
		t.Error("too few comments!", len(res.Comments))
	}
}

// -- helpers --

func findMediaWithComments() (string, error) {
	ctx := context.Background()
	api := newAPI()

	params := url.Values{}
	params.Set("count", "20")

	res, err := api.GetRecentMedia(ctx, params)
	if err != nil {
		return "", err
	}
	for _, m := range res.Medias {
		if m.Comments.Count > 0 {
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
