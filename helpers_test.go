package instagram

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/kr/pretty"
)

func TestMediaJSON(t *testing.T) {
	tests := []struct {
		desc  string
		json  string
		media Media
	}{
		{
			desc:  "zero value",
			json:  `{"type":"","id":"","users_in_photo":null,"filter":"","tags":null,"comments":{"count":0},"caption":{"id":"","text":"","from":{},"created_time":null},"likes":{"count":0},"link":"","user":{},"images":{},"videos":{},"carousel_media":null,"location":{"id":""},"user_has_liked":false,"created_time":null}`,
			media: Media{},
		},
		{
			desc: "round trip",
			json: `{"type":"","id":"123","users_in_photo":null,"filter":"","tags":null,"comments":{"count":0},"caption":{"id":"","text":"","from":{},"created_time":null},"likes":{"count":0},"link":"","user":{},"images":{},"videos":{},"carousel_media":null,"location":{"id":""},"user_has_liked":false,"created_time":"1550173420"}`,
			media: Media{
				Id:          "123",
				CreatedTime: time.Date(2019, 2, 14, 19, 43, 40, 0, time.UTC),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			gotJSON, err := json.Marshal(tt.media)
			if err != nil {
				t.Fatalf("Marshal: %v", err)
			}
			if got, want := string(gotJSON), tt.json; got != want {
				t.Errorf("Marshal\ngot  %s\nwant %s", got, want)
			}
			var got Media
			if err := json.Unmarshal([]byte(tt.json), &got); err != nil {
				t.Fatalf("Unmarshal: %v", err)
			}
			if got, want := got, tt.media; !reflect.DeepEqual(got, want) {
				t.Errorf("Unmarshal got %#v want %#v", got, want)
				pretty.Ldiff(t, got, want)
			}
		})
	}
}

func TestCommentJSON(t *testing.T) {
	tests := []struct {
		desc          string
		json          string
		comment       Comment
		unmarshalOnly bool
	}{
		{
			desc:    "zero value",
			json:    `{"id":"","text":"","from":{},"created_time":null}`,
			comment: Comment{},
		},
		{
			desc:          "null",
			json:          `null`,
			comment:       Comment{},
			unmarshalOnly: true,
		},
		{
			desc: "round trip",
			json: `{"id":"123","text":"","from":{},"created_time":"1550173420"}`,
			comment: Comment{
				Id:          "123",
				CreatedTime: time.Date(2019, 2, 14, 19, 43, 40, 0, time.UTC),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if !tt.unmarshalOnly {
				gotJSON, err := json.Marshal(tt.comment)
				if err != nil {
					t.Fatalf("Marshal: %v", err)
				}
				if got, want := string(gotJSON), tt.json; got != want {
					t.Errorf("Marshal\ngot  %s\nwant %s", got, want)
				}
			}
			var got Comment
			if err := json.Unmarshal([]byte(tt.json), &got); err != nil {
				t.Fatalf("Unmarshal: %v", err)
			}
			if got, want := got, tt.comment; !reflect.DeepEqual(got, want) {
				t.Errorf("Unmarshal got %#v want %#v", got, want)
				pretty.Ldiff(t, got, want)
			}
		})
	}
}

func TestLocationJSON(t *testing.T) {
	tests := []struct {
		desc          string
		json          string
		location      Location
		unmarshalOnly bool
	}{
		{
			desc:     "zero value",
			json:     `{"id":""}`,
			location: Location{},
		},
		{
			desc:          "null",
			json:          `null`,
			location:      Location{},
			unmarshalOnly: true,
		},
		{
			desc: "round trip with string",
			json: `{"id":"123","name":"florida"}`,
			location: Location{
				Id:   "123",
				Name: "florida",
			},
		},
		{
			desc: "round trip with number",
			json: `{"id":123,"name":"florida"}`,
			location: Location{
				Id:   "123",
				Name: "florida",
			},
			unmarshalOnly: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if !tt.unmarshalOnly {
				gotJSON, err := json.Marshal(tt.location)
				if err != nil {
					t.Fatalf("Marshal: %v", err)
				}
				if got, want := string(gotJSON), tt.json; got != want {
					t.Errorf("Marshal\ngot  %s\nwant %s", got, want)
				}
			}
			var got Location
			if err := json.Unmarshal([]byte(tt.json), &got); err != nil {
				t.Fatalf("Unmarshal: %v", err)
			}
			if got, want := got, tt.location; !reflect.DeepEqual(got, want) {
				t.Errorf("Unmarshal got %#v want %#v", got, want)
				pretty.Ldiff(t, got, want)
			}
		})
	}
}
