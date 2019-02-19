package instagram

import (
	"encoding/json"
	"fmt"
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
	if cj.CommentJSON == nil {
		return nil
	}
	*c = Comment(*cj.CommentJSON)
	c.CreatedTime = time.Time(cj.CreatedTime)
	return nil
}

// LocationJSON is Location for JSON.
type LocationJSON Location

type locationJSONWrapper struct {
	ID interface{} `json:"id"`
	*LocationJSON
}

// MarshalJSON implements JSON.
func (l Location) MarshalJSON() ([]byte, error) {
	ll := LocationJSON(l)
	lj := locationJSONWrapper{l.Id, &ll}
	return json.Marshal(lj)
}

// UnmarshalJSON implements JSON.
func (l *Location) UnmarshalJSON(in []byte) error {
	lj := locationJSONWrapper{}
	if err := json.Unmarshal(in, &lj); err != nil {
		return err
	}
	if lj.LocationJSON == nil {
		return nil
	}
	*l = Location(*lj.LocationJSON)
	switch v := lj.ID.(type) {
	case nil:
		// ok
	case string:
		l.Id = v
	case float64:
		l.Id = fmt.Sprintf("%0.f", v)
	default:
		return fmt.Errorf("unknown type for location id: %T", lj.ID)
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
