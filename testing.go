package instagram

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func newTestAPIServer(t *testing.T, fix func(*http.Request) string) (*Api, *httptest.Server) {
	httpClient, mux, server := initTestServer()

	mux.HandleFunc("/v1/", func(w http.ResponseWriter, r *http.Request) {
		fixture := fix(r)
		if fixture == "" {
			t.Fatalf("no fixture for request: %s", r.URL)
			return
		}

		path := fmt.Sprintf("testdata/%s.json", fixture)
		f, err := os.Open(path)
		if err != nil {
			t.Fatalf("failed to open %s: %s", path, err)
			return
		}
		defer f.Close()

		w.Header().Set("Content-Type", "application/json")
		io.Copy(w, f)
	})

	api := &Api{
		EnforceSignedRequest: true,
		HTTPClient:           httpClient,
	}

	return api, server
}

func initTestServer() (*http.Client, *http.ServeMux, *httptest.Server) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	serverURL, _ := url.Parse(server.URL)
	transport := &rewriteTransport{serverURL, &http.Transport{}}
	client := &http.Client{Transport: transport}
	return client, mux, server
}

type rewriteTransport struct {
	serverURL *url.URL
	Transport http.RoundTripper
}

func (t *rewriteTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.URL.Scheme = "http"
	r.URL.Host = t.serverURL.Host
	return t.Transport.RoundTrip(r)
}
