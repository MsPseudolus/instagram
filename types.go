package instagram

import "time"

// Instagram User Object. Note that user objects are not always fully returned.
// Be sure to see the descriptions on the instagram documentation for any given endpoint.
type User struct {
	Id             string      `json:"id"`
	Username       string      `json:"username"`
	FirstName      string      `json:"first_name"`
	LastName       string      `json:"last_name"`
	FullName       string      `json:"full_name"`
	ProfilePicture string      `json:"profile_picture"`
	Bio            string      `json:"bio"`
	Website        string      `json:"website"`
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
	User     User
	Position Position
}

// A position in a media
type Position struct {
	X float64
	Y float64
}

// Instagram tag
type Tag struct {
	MediaCount int64 `json:"media_count"`
	Name       string
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
	Count int64  `json:"likes"`
	Data  []User `json:"data,omitempty"`
}

type Images struct {
	LowResolution      Image `json:"low_resolution"`
	LowBandwidth       Image `json:"low_bandwidth"`
	Thumbnail          Image `json:"thumbnail"`
	StandardResolution Image `json:"standard_resolution"`
}

type Image struct {
	Id     string
	Url    string
	Width  int64
	Height int64
}

type CarouselMedia struct {
	Type         string
	Images       Images
	Videos       Images
	UsersInPhoto []UserPosition `json:"users_in_photo"`
}

type Location struct {
	Id        string
	Name      string
	Latitude  float64
	Longitude float64
}

// If another app uploaded the media, then this is the place it is given. As of 11/2013, Hipstamic is the only allowed app
type Attribution struct {
	Website   string
	ItunesUrl string
	Name      string
}
