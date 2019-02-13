package instagram

import (
	"net/url"
)

// Get basic information about authenticated user.
// Gets /users/self
func (api *Api) GetSelf() (res *UserResponse, err error) {
	res = new(UserResponse)
	err = api.get("/users/self", nil, res)
	return
}

// Get the most recent media published by a user. May return a mix of both image and video types.
// Gets /users/{user-id}/media/recent
func (api *Api) GetRecentMedia(string, params url.Values) (res *PaginatedMediasResponse, err error) {
	res = new(PaginatedMediasResponse)
	err = api.get("/users/self/media/recent", params, res)
	return
}

// Verify a valid client keys and user tokens by making a small request
func (api *Api) VerifyCredentials() (ok bool, err error) {
	_, err = api.GetSelf()
	return err == nil, err
}
