package instagram

import (
	"fmt"
	"net/url"
)

// GetSelf returns basic information about the authenticated user.
// REST API: GET /users/self
func (api *Api) GetSelf() (res *UserResponse, err error) {
	res = new(UserResponse)
	err = api.get("/users/self", nil, res)
	return
}

// VerifyCredentials checks client keys and user tokens by making a small
// request.
func (api *Api) VerifyCredentials() (ok bool, err error) {
	_, err = api.GetSelf()
	return err == nil, err
}

// GetRecentMedia the most recent media published by the authenticated user. May return a mix of types.
// REST API: GET /users/self/media/recent
func (api *Api) GetRecentMedia(params url.Values) (res *PaginatedMediasResponse, err error) {
	res = new(PaginatedMediasResponse)
	err = api.get("/users/self/media/recent", params, res)
	return
}

// GetMediaRecentComments returns a list of recent comments on a media.
// Requires scope: comments.
// REST API: GET /media/{media-id}/comments
func (api *Api) GetMediaRecentComments(mediaID string) (res *CommentsResponse, err error) {
	res = new(CommentsResponse)
	err = api.get(fmt.Sprintf("/media/%s/comments", mediaID), nil, res)
	return
}
