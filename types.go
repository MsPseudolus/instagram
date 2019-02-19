package instagram

import "time"

// Instagram User Object. Note that user objects are not always fully returned.
// Be sure to see the descriptions on the instagram documentation for any given endpoint.
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

// Instagram User Counts object. Returned on User objects
type UserCounts struct {
	Media      int64 `json:"media"`
	Follows    int64 `json:"follows"`
	FollowedBy int64 `json:"followed_by"`
}

// Instagram Media object
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
	Images         Images          `json:"images"`
	Videos         Images          `json:"videos"`
	CarouselMedias []CarouselMedia `json:"carousel_media"`
	Location       Location        `json:"location"`
	UserHasLiked   bool            `json:"user_has_liked"`
	Attribution    *Attribution    `json:"attribution,omitempty"`
}

// A pair of user object and position
type UserPosition struct {
	User     User     `json:"user"`
	Position Position `json:"position"`
}

// A position in a media
type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Comments struct {
	Count int64     `json:"count"`
	Data  []Comment `json:"data,omitempty"`
}

type Comment struct {
	Id          string    `json:"id"`
	Text        string    `json:"text"`
	From        User      `json:"from"`
	CreatedTime time.Time `json:"created_time"`
}

type Likes struct {
	Count int64  `json:"count"`
	Data  []User `json:"data,omitempty"`
}

type Images struct {
	LowResolution      *Image `json:"low_resolution,omitempty"`
	LowBandwidth       *Image `json:"low_bandwidth,omitempty"`
	Thumbnail          *Image `json:"thumbnail,omitempty"`
	StandardResolution *Image `json:"standard_resolution,omitempty"`
}

type Image struct {
	Id     string `json:"id,omitempty"`
	Url    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

type CarouselMedia struct {
	Type         string         `json:"type"`
	Images       Images         `json:"images"`
	Videos       Images         `json:"videos"`
	UsersInPhoto []UserPosition `json:"users_in_photo"`
}

type Location struct {
	Id        string  `json:"id,omitempty"`
	Name      string  `json:"name,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

// If another app uploaded the media, then this is the place it is given. As of 11/2013, Hipstamic is the only allowed app
type Attribution struct {
	Website   string
	ItunesUrl string
	Name      string
}
