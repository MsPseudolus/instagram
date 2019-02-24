package instagram

import (
	"errors"
	"net/url"
	"testing"
)

func TestGetAuthorizeURL(t *testing.T) {
	oauth := OAuth{
		ClientID:    "c1",
		RedirectURI: "http://localhost/",
	}

	want := "https://api.instagram.com/oauth/authorize/?client_id=c1&redirect_uri=http%3A%2F%2Flocalhost%2F&response_type=code&state=mystate"

	state := "mystate"
	got := oauth.GetAuthorizeURL(state)

	if got != want {
		t.Errorf("Authroize URL got\n%s\nwant\n%s", got, want)
	}
}

func TestGetCodeFromRedirect(t *testing.T) {
	tests := []struct {
		desc     string
		input    string
		state    string
		wantCode string
		wantErr  error
	}{
		{
			desc:     "matching state",
			input:    "http://localhost/?code=e2b2186f34aa4f70849182be1ca340fc&state=mystate",
			state:    "mystate",
			wantCode: "e2b2186f34aa4f70849182be1ca340fc",
		},
		{
			desc:    "not matching state",
			input:   "http://localhost/?code=e2b2186f34aa4f70849182be1ca340fc&state=foo",
			state:   "mystate",
			wantErr: errors.New("state does not match"),
		},
		{
			desc:    "no code",
			input:   "http://localhost/?code=&state=mystate",
			state:   "mystate",
			wantErr: errors.New("missing code"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			oauth := OAuth{}

			u, err := url.Parse(tt.input)
			if err != nil {
				t.Fatalf("Parse: %v", err)
			}

			code, err := oauth.GetCodeFromRedirect(u, tt.state)

			if err != nil {
				if tt.wantErr == nil {
					t.Errorf("GetCodeFromRedirect: %v", err)
				} else if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error got %s want %s", err, tt.wantErr)
				}
			}

			if code != tt.wantCode {
				t.Errorf("Got code %s want %s", code, tt.wantCode)
			}
		})
	}
}
