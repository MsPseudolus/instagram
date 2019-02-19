package instagram

import "time"

// User is the instagram user. This struct is used in several contexts, each of
// which may populate different data.
type User struct {
	Id             string      `json:"id,omitempty"`
	Username       string      `json:"username,omitempty"`
	FirstName      string      `json:"first_name,omitempty"`
	LastName       string      `json:"last_name,omitempty"`
	FullName       string      `json:"full_name,omitempty"`
	ProfilePicture string      `json:"profile_picture,omitempty"`
	Bio            string      `json:"bio,omitempty"`
	Website        string      `json:"website,omitempty"`
	Counts         *UserCounts `json:"counts,omitempty"`
}

// UserCounts is information about the user's counts.
type UserCounts struct {
	Media      int64 `json:"media"`
	Follows    int64 `json:"follows"`
	FollowedBy int64 `json:"followed_by"`
}

// Media is the overall wrapper for any kind of Instagram media.
type Media struct {
	Type           string          `json:"type"`
	Id             string          `json:"id"`
	UsersInPhoto   []UserPosition  `json:"users_in_photo"`
	Filter         string          `json:"filter"`
	Tags           []string        `json:"tags"`
	Comments       Comments        `json:"comments"`
	Caption        Comment         `json:"caption"`
	Likes          Likes           `json:"likes"`
	Link           string          `json:"link"`
	User           User            `json:"user"`
	CreatedTime    time.Time       `json:"created_time"`
	Images         MediaVariants   `json:"images"`
	Videos         MediaVariants   `json:"videos"`
	CarouselMedias []CarouselMedia `json:"carousel_media"`
	Location       Location        `json:"location"`
	UserHasLiked   bool            `json:"user_has_liked"`
	Attribution    *Attribution    `json:"attribution,omitempty"`
}

// UserPosition describes a user tagged in a media.
type UserPosition struct {
	User     User     `json:"user"`
	Position Position `json:"position"`
}

// Position is the position of a tagged user.
type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Comments decribes the comments on a media.
type Comments struct {
	Count int64 `json:"count"`
}

// Comment is a comment on a media.
type Comment struct {
	Id          string    `json:"id"`
	Text        string    `json:"text"`
	From        User      `json:"from"`
	CreatedTime time.Time `json:"created_time"`
}

// Likes describes the likes on a media.
type Likes struct {
	Count int64 `json:"count"`
}

// MediaVariants is the set of represetations of the media.
type MediaVariants struct {
	LowResolution      *MediaVariant `json:"low_resolution,omitempty"`
	LowBandwidth       *MediaVariant `json:"low_bandwidth,omitempty"`
	Thumbnail          *MediaVariant `json:"thumbnail,omitempty"`
	StandardResolution *MediaVariant `json:"standard_resolution,omitempty"`
}

// MediaVariant is a specific variant of the media.
type MediaVariant struct {
	Id     string `json:"id,omitempty"`
	Url    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

// CarouselMedia describes a carousel media type.
type CarouselMedia struct {
	Type         string         `json:"type"`
	Images       MediaVariants  `json:"images"`
	Videos       MediaVariants  `json:"videos"`
	UsersInPhoto []UserPosition `json:"users_in_photo"`
}

// Location is the location of a media.
type Location struct {
	Id        string  `json:"id,omitempty"`
	Name      string  `json:"name,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

// Attribution is if another app uploaded the media, then this is the place it
// is given. As of 11/2013, Hipstamic is the only allowed app
type Attribution struct {
	Website   string
	ItunesUrl string
	Name      string
}
