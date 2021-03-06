package instagram

import (
	"context"
	"net/url"
	"strings"
)

// NextMedias returns the next page of media
func (api *API) NextMedias(ctx context.Context, mp MediaPagination) (res *PaginatedMediasResponse, err error) {
	res = new(PaginatedMediasResponse)
	err = api.next(ctx, mp.Pagination, res)
	return
}

func (api *API) next(ctx context.Context, p Pagination, res interface{}) error {
	done, uri, path, uriParams, err := p.NextPage()
	if err != nil || done == true {
		return err
	}

	// Sign params if using the secure api
	if api.EnforceSignedRequest {
		uriParams = signParams(path, uriParams, api.ClientSecret)
	}

	req, err := buildGetRequest(uri, uriParams)
	if err != nil {
		return err
	}

	return api.do(ctx, req, res)
}

// NextPage returns the next page's uri and parameters
func (p Pagination) NextPage() (done bool, uri string, path string, params url.Values, err error) {
	if p.NextURL == "" {
		// We're done. Theres no more pages
		done = true
		return
	}

	urlStruct, err := url.Parse(p.NextURL)
	if err != nil {
		return
	}

	params = urlStruct.Query()
	// Remove `sig` key that was set by the initial request
	params.Del("sig")
	urlStruct.RawQuery = ""

	done = false
	path = strings.Replace(urlStruct.Path, "/v1", "", 1)
	uri = urlStruct.String()
	return
}
