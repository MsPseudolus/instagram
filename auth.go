package instagram

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// NewOAuth initializes a new OAuth flow.
func NewOAuth(clientID string, clientSecret string, redirectURI string) OAuth {
	return OAuth{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURI:  redirectURI,
		HTTPClient:   &http.Client{},
	}
}

// OAuth is an OAuth flow for Instagram. Use this to get an AccessToken.
type OAuth struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	HTTPClient   *http.Client
}

// GetAuthorizeURL returns a URL to send a user to.
func (o OAuth) GetAuthorizeURL(state string) string {
	authURL, _ := url.Parse("https://api.instagram.com/oauth/authorize/")
	query := authURL.Query()
	query.Set("client_id", o.ClientID)
	query.Set("redirect_uri", o.RedirectURI)
	query.Set("response_type", "code")
	if state != "" {
		query.Set("state", state)
	}
	log.Printf("E %s", query.Encode())
	authURL.RawQuery = query.Encode()

	log.Printf("U %s", authURL.String())
	return authURL.String()
}

// GetCodeFromRedirect parses the redirect URL. It validates the state and
// returns the code.
func (o OAuth) GetCodeFromRedirect(url *url.URL, state string) (string, error) {
	q := url.Query()
	c := q.Get("code")
	s := q.Get("state")

	if state != "" && state != s {
		return "", errors.New("state does not match")
	}
	if c == "" {
		return "", errors.New("missing code")
	}
	return c, nil
}

// GetAccessToken trades the code for an AccessToken.
func (o OAuth) GetAccessToken(ctx context.Context, code string) (string, error) {
	form := url.Values{}
	form.Set("client_id", o.ClientID)
	form.Set("client_secret", o.ClientSecret)
	form.Set("redirect_uri", o.RedirectURI)
	form.Set("grant_type", "authorization_code")
	form.Set("code", code)

	req, err := http.NewRequest("POST", "https://api.instagram.com/oauth/access_token", strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}
	req = req.WithContext(ctx)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := o.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}

	type response struct {
		AccessToken string `json:"access_token"`
	}

	var tokenResponse response
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return "", err
	}
	return tokenResponse.AccessToken, nil
}
