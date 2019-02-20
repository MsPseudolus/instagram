// Package instagram provides a minimialist instagram API wrapper.
package instagram

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
)

var (
	baseURL = "https://api.instagram.com/v1"
)

// API is the Instagram API.
type API struct {
	ClientID             string
	ClientSecret         string
	AccessToken          string
	EnforceSignedRequest bool

	// HTTPClient sets a custom HTTP Client used to make requests.
	HTTPClient *http.Client

	// KeepRawBody set to true will store the raw API response body
	// of the last request in RawBody.
	KeepRawBody bool
	RawBody     io.Reader

	// Header contains the raw API response header of the last request.
	Header http.Header
}

// New creates an API with either a ClientID OR an accessToken. Only one is
// required. Access tokens are preferred because they keep rate limiting down.
// If enforceSignedRequest is set to true, then clientSecret is required
func New(clientID string, clientSecret string, accessToken string, enforceSignedRequest bool) *API {
	if clientID == "" && accessToken == "" {
		panic("ClientID or AccessToken must be given to create an API")
	}

	if enforceSignedRequest && clientSecret == "" {
		panic("ClientSecret is required for signed request")
	}

	return &API{
		ClientID:             clientID,
		ClientSecret:         clientSecret,
		AccessToken:          accessToken,
		HTTPClient:           &http.Client{},
		EnforceSignedRequest: enforceSignedRequest,
	}
}

// -- Implementation of request --
func signParams(path string, params url.Values, clientSecret string) url.Values {
	message := path
	keys := []string{}

	for k := range params {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, v := range keys {
		message += "|" + v + "=" + params.Get(v)
	}

	hash := hmac.New(sha256.New, []byte(clientSecret))
	hash.Write([]byte(message))

	params.Set("sig", hex.EncodeToString(hash.Sum(nil)))
	return params
}

func buildGetRequest(urlStr string, params url.Values) (*http.Request, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	// If we are getting, then we can't merge query params
	if params != nil {
		if u.RawQuery != "" {
			return nil, fmt.Errorf("Cannot merge query params in urlStr and params")
		}
		u.RawQuery = params.Encode()
	}

	return http.NewRequest("GET", u.String(), nil)
}

func (api *API) extendParams(p url.Values) url.Values {
	if p == nil {
		p = url.Values{}
	}
	if api.AccessToken != "" {
		p.Set("access_token", api.AccessToken)
	} else {
		p.Set("client_id", api.ClientID)
	}
	return p
}

func (api *API) get(ctx context.Context, path string, params url.Values, r interface{}) error {
	params = api.extendParams(params)
	// Sign request if ForceSignedRequest is set to true
	if api.EnforceSignedRequest {
		params = signParams(path, params, api.ClientSecret)
	}

	req, err := buildGetRequest(urlify(path), params)
	if err != nil {
		return err
	}
	return api.do(ctx, req, r)
}

func (api *API) do(ctx context.Context, req *http.Request, r interface{}) error {
	req = req.WithContext(ctx)
	resp, err := api.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()

	api.Header = resp.Header

	if resp.StatusCode != 200 {
		return api.apiError(resp)
	}

	return api.decodeResponse(resp.Body, r)
}

func (api *API) decodeResponse(body io.Reader, to interface{}) error {

	var r io.Reader
	if api.KeepRawBody {
		var buf bytes.Buffer
		r = io.TeeReader(body, &buf)
		api.RawBody = &buf
	} else {
		r = body
	}

	err := json.NewDecoder(r).Decode(to)

	if err != nil {
		return fmt.Errorf("instagram: error decoding body; %s", err.Error())
	}
	return nil
}

func (api *API) apiError(resp *http.Response) error {
	m := new(metaResponse)
	if err := api.decodeResponse(resp.Body, m); err != nil {
		return err
	}

	var err MetaError
	if m.Meta != nil {
		err = MetaError(*m.Meta)
	} else {
		err = MetaError(Meta{Code: resp.StatusCode, ErrorMessage: resp.Status})
	}
	return &err
}

func urlify(path string) string {
	return baseURL + path
}

// MetaError is an error from response metadata.
type MetaError Meta

func (m *MetaError) Error() string {
	return fmt.Sprintf("Error making api call: Code %d %s %s", m.Code, m.ErrorType, m.ErrorMessage)
}

func ensureParams(v url.Values) url.Values {
	if v == nil {
		return url.Values{}
	}
	return v
}
