package instagram

import (
	"encoding/json"
	"fmt"
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
			json:  `{"type":"","id":"","users_in_photo":null,"filter":"","tags":null,"comments":{"count":0},"caption":{"id":"","text":"","from":{"id":"","username":"","first_name":"","last_name":"","full_name":"","profile_picture":"","bio":"","website":""},"created_time":null},"likes":{"likes":0},"link":"","user":{"id":"","username":"","first_name":"","last_name":"","full_name":"","profile_picture":"","bio":"","website":""},"images":{"low_resolution":{"Id":"","Url":"","Width":0,"Height":0},"low_bandwidth":{"Id":"","Url":"","Width":0,"Height":0},"thumbnail":{"Id":"","Url":"","Width":0,"Height":0},"standard_resolution":{"Id":"","Url":"","Width":0,"Height":0}},"videos":{"low_resolution":{"Id":"","Url":"","Width":0,"Height":0},"low_bandwidth":{"Id":"","Url":"","Width":0,"Height":0},"thumbnail":{"Id":"","Url":"","Width":0,"Height":0},"standard_resolution":{"Id":"","Url":"","Width":0,"Height":0}},"carousel_media":null,"location":{"Id":"","Name":"","Latitude":0,"Longitude":0},"user_has_liked":false,"created_time":null}`,
			media: Media{},
		},
		{
			desc: "round trip",
			json: `{"type":"","id":"123","users_in_photo":null,"filter":"","tags":null,"comments":{"count":0},"caption":{"id":"","text":"","from":{"id":"","username":"","first_name":"","last_name":"","full_name":"","profile_picture":"","bio":"","website":""},"created_time":null},"likes":{"likes":0},"link":"","user":{"id":"","username":"","first_name":"","last_name":"","full_name":"","profile_picture":"","bio":"","website":""},"images":{"low_resolution":{"Id":"","Url":"","Width":0,"Height":0},"low_bandwidth":{"Id":"","Url":"","Width":0,"Height":0},"thumbnail":{"Id":"","Url":"","Width":0,"Height":0},"standard_resolution":{"Id":"","Url":"","Width":0,"Height":0}},"videos":{"low_resolution":{"Id":"","Url":"","Width":0,"Height":0},"low_bandwidth":{"Id":"","Url":"","Width":0,"Height":0},"thumbnail":{"Id":"","Url":"","Width":0,"Height":0},"standard_resolution":{"Id":"","Url":"","Width":0,"Height":0}},"carousel_media":null,"location":{"Id":"","Name":"","Latitude":0,"Longitude":0},"user_has_liked":false,"created_time":"1550173420"}`,
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
				t.Errorf("Marshal\ngot %s\nwant %s", got, want)
			}
			var gotMedia Media
			if err := json.Unmarshal([]byte(tt.json), &gotMedia); err != nil {
				t.Fatalf("Unmarshal: %v", err)
			}
			if !reflect.DeepEqual(gotMedia, tt.media) {
				fmt.Printf("%s\n%s\n", gotMedia.CreatedTime, tt.media.CreatedTime)
				t.Errorf("Unmarshal got %#v want %#v", gotMedia, tt.media)
				pretty.Ldiff(t, gotMedia, tt.media)
			}
		})
	}
}

func TestCommentJSON(t *testing.T) {
	tests := []struct {
		desc    string
		json    string
		comment Comment
	}{
		{
			desc:    "zero value",
			json:    `{"id":"","text":"","from":{"id":"","username":"","first_name":"","last_name":"","full_name":"","profile_picture":"","bio":"","website":""},"created_time":null}`,
			comment: Comment{},
		},
		{
			desc:    "null",
			json:    `null`,
			comment: Comment{},
		},
		{
			desc: "round trip",
			json: `{"id":"123","text":"","from":{"id":"","username":"","first_name":"","last_name":"","full_name":"","profile_picture":"","bio":"","website":""},"created_time":"1550173420"}`,
			comment: Comment{
				Id:          "123",
				CreatedTime: time.Date(2019, 2, 14, 19, 43, 40, 0, time.UTC),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			gotJSON, err := json.Marshal(tt.comment)
			if err != nil {
				t.Fatalf("Marshal: %v", err)
			}
			if got, want := string(gotJSON), tt.json; got != want {
				t.Errorf("Marshal\ngot  %s\nwant %s", got, want)
			}
			var gotComment Comment
			if err := json.Unmarshal([]byte(tt.json), &gotComment); err != nil {
				t.Fatalf("Unmarshal: %v", err)
			}
			if !reflect.DeepEqual(gotComment, tt.comment) {
				fmt.Printf("%s\n%s\n", gotComment.CreatedTime, tt.comment.CreatedTime)
				t.Errorf("Unmarshal got %#v want %#v", gotComment, tt.comment)
				pretty.Ldiff(t, gotComment, tt.comment)
			}
		})
	}
}
