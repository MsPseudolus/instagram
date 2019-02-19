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
			json:  `{"Type":"","Id":"","UsersInPhoto":null,"Filter":"","Tags":null,"Comments":{"Count":0,"Data":null},"Caption":{"created_time":"","Text":"","From":{"id":"","username":"","first_name":"","last_name":"","full_name":"","profile_picture":"","Bio":"","Website":"","Counts":null},"Id":""},"Likes":{"Count":0,"Data":null},"Link":"","User":{"id":"","username":"","first_name":"","last_name":"","full_name":"","profile_picture":"","Bio":"","Website":"","Counts":null},"Images":{"low_resolution":{"Id":"","Url":"","Width":0,"Height":0},"low_bandwidth":{"Id":"","Url":"","Width":0,"Height":0},"Thumbnail":{"Id":"","Url":"","Width":0,"Height":0},"standard_resolution":{"Id":"","Url":"","Width":0,"Height":0}},"Videos":{"low_resolution":{"Id":"","Url":"","Width":0,"Height":0},"low_bandwidth":{"Id":"","Url":"","Width":0,"Height":0},"Thumbnail":{"Id":"","Url":"","Width":0,"Height":0},"standard_resolution":{"Id":"","Url":"","Width":0,"Height":0}},"carousel_media":null,"Location":{"Id":null,"Name":"","Latitude":0,"Longitude":0},"user_has_liked":false,"Attribution":null,"created_time":null}`,
			media: Media{},
		},
		{
			desc: "round trip",
			json: `{"Type":"","Id":"123","UsersInPhoto":null,"Filter":"","Tags":null,"Comments":{"Count":0,"Data":null},"Caption":{"created_time":"","Text":"","From":{"id":"","username":"","first_name":"","last_name":"","full_name":"","profile_picture":"","Bio":"","Website":"","Counts":null},"Id":""},"Likes":{"Count":0,"Data":null},"Link":"","User":{"id":"","username":"","first_name":"","last_name":"","full_name":"","profile_picture":"","Bio":"","Website":"","Counts":null},"Images":{"low_resolution":{"Id":"","Url":"","Width":0,"Height":0},"low_bandwidth":{"Id":"","Url":"","Width":0,"Height":0},"Thumbnail":{"Id":"","Url":"","Width":0,"Height":0},"standard_resolution":{"Id":"","Url":"","Width":0,"Height":0}},"Videos":{"low_resolution":{"Id":"","Url":"","Width":0,"Height":0},"low_bandwidth":{"Id":"","Url":"","Width":0,"Height":0},"Thumbnail":{"Id":"","Url":"","Width":0,"Height":0},"standard_resolution":{"Id":"","Url":"","Width":0,"Height":0}},"carousel_media":null,"Location":{"Id":null,"Name":"","Latitude":0,"Longitude":0},"user_has_liked":false,"Attribution":null,"created_time":"1550173420"}`,
			media: Media{
				Id:          "123",
				CreatedTime: time.Date(2019, 2, 14, 19, 43, 40, 0, time.UTC),
			},
		},
	}
	for _, tt := range tests {
		gotJSON, err := json.Marshal(tt.media)
		if err != nil {
			t.Fatalf("Marshal: %v", err)
		}
		if got, want := string(gotJSON), tt.json; got != want {
			t.Errorf("Marshal got\n  %s\nwant %s", got, want)
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
	}
}
