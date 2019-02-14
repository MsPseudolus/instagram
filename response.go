package instagram

type metaResponse struct {
	Meta *Meta
}

// UserResponse is the API response for GetSelf()
type UserResponse struct {
	metaResponse
	User *User `json:"data"`
}

// PaginatedMediasResponse is the API response for GetRecentMedia()
type PaginatedMediasResponse struct {
	metaResponse
	Medias     []Media `json:"data"`
	Pagination *MediaPagination
}

// MediaPagination will give you an easy way to request the next page of media.
type MediaPagination struct {
	*Pagination
}

// CommentsResponse is the API response for GetMediaRecentComments()
type CommentsResponse struct {
	metaResponse
	Comments []Comment `json:"data"`
}

// Pagination describes how to get the next page of results.
type Pagination struct {
	NextUrl   string `json:"next_url"`
	NextMaxId string `json:"next_max_id"`
}

// Meta is the response information.
type Meta struct {
	Code         int
	ErrorType    string `json:"error_type"`
	ErrorMessage string `json:"error_message"`
}
