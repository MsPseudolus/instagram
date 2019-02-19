package instagram

import (
	"encoding/json"
	"strconv"
	"time"
)

// MediaJSON is Media for JSON.
type MediaJSON Media

type mediaJSONWrapper struct {
	*MediaJSON
	CreatedTime stringTime `json:"created_time"`
}

// MarshalJSON implements JSON.
func (m Media) MarshalJSON() ([]byte, error) {
	mm := MediaJSON(m)
	mj := mediaJSONWrapper{&mm, stringTime(m.CreatedTime)}
	return json.Marshal(mj)
}

// UnmarshalJSON implements JSON.
func (m *Media) UnmarshalJSON(in []byte) error {
	mj := mediaJSONWrapper{}
	if err := json.Unmarshal(in, &mj); err != nil {
		return err
	}
	*m = Media(*mj.MediaJSON)
	m.CreatedTime = time.Time(mj.CreatedTime)
	return nil
}

// CommentJSON is Comment for JSON.
type CommentJSON Comment

type commentJSONWrapper struct {
	*CommentJSON
	CreatedTime stringTime `json:"created_time"`
}

// MarshalJSON implements JSON.
func (c Comment) MarshalJSON() ([]byte, error) {
	cc := CommentJSON(c)
	cj := commentJSONWrapper{&cc, stringTime(c.CreatedTime)}
	return json.Marshal(cj)
}

// UnmarshalJSON implements JSON.
func (c *Comment) UnmarshalJSON(in []byte) error {
	cj := commentJSONWrapper{}
	if err := json.Unmarshal(in, &cj); err != nil {
		return err
	}
	if cj.CommentJSON != nil {
		*c = Comment(*cj.CommentJSON)
		c.CreatedTime = time.Time(cj.CreatedTime)
	}
	return nil
}

type stringTime time.Time

func (s stringTime) MarshalJSON() ([]byte, error) {
	t := time.Time(s)
	if t.IsZero() {
		return []byte("null"), nil
	}
	u := t.UTC().Unix()
	ss := strconv.FormatInt(u, 10)
	return json.Marshal(ss)
}
func (s *stringTime) UnmarshalJSON(in []byte) error {
	if string(in) == "null" {
		return nil
	}
	var ss string
	if err := json.Unmarshal(in, &ss); err != nil {
		return err
	}
	unix, err := strconv.ParseInt(string(ss), 10, 64)
	if err != nil {
		return err
	}
	t := time.Unix(unix, 0).UTC()
	*s = stringTime(t)
	return nil
}

// StringUnixTime is a string that's actually unix time.
type StringUnixTime string

// Time returns a time.Time
func (s StringUnixTime) Time() (t time.Time, err error) {
	unix, err := strconv.ParseInt(string(s), 10, 64)
	if err != nil {
		return
	}

	t = time.Unix(unix, 0).UTC()
	return
}

// LocationId is an ambiguous value that can become a string ID.
type LocationId interface{}

//func (l LocationId) String() string {
//switch v := l.(type) {
//case string:
//return v
//case int64:
//return fmt.Sprintf("%d", v)
//default:
//return ""
//}
//}
