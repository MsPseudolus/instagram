package instagram

import (
	"fmt"
)

// Get a full list of comments on a media.
// Required Scope: comments
// Gets /media/{media-id}/comments
func (api *Api) GetMediaRecentComments(mediaId string) (res *CommentsResponse, err error) {
	res = new(CommentsResponse)
	err = api.get(fmt.Sprintf("/media/%s/comments", mediaId), nil, res)
	return
}
