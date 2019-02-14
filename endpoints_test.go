package instagram

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/kr/pretty"
)

func TestMediaTypes(t *testing.T) {
	tests := []struct {
		desc      string
		fix       string
		wantMedia Media
	}{
		{
			desc: "image media",
			fix:  "image",
			wantMedia: Media{
				Id:          "1979320569926821011_11073382793",
				Type:        "image",
				CreatedTime: "1550173420",

				Link:   "https://www.instagram.com/p/Bt39ZJLHKSTFwXShw402xx8W9loUPHTyH5BsqY0/",
				Filter: "Crema",

				User: &User{
					Id:             "11073382793",
					Username:       "go_ig_test_0219",
					FullName:       "Golang Client",
					ProfilePicture: "https://scontent-sjc3-1.cdninstagram.com/vp/504ac2fa79adb1d412b31cab19be8d36/5CDDD9F1/t51.2885-19/44884218_345707102882519_2446069589734326272_n.jpg?_nc_ht=scontent-sjc3-1.cdninstagram.com",
				},

				Likes: &Likes{
					Count: 1,
				},

				Comments: &Comments{
					Count: 2,
				},

				Caption: &Caption{
					Id:          "18002756710177046",
					CreatedTime: "1550173420",
					Text:        "Photo post #0219test",
					From: &User{
						Id:             "11073382793",
						Username:       "go_ig_test_0219",
						FullName:       "Golang Client",
						ProfilePicture: "https://scontent-sjc3-1.cdninstagram.com/vp/504ac2fa79adb1d412b31cab19be8d36/5CDDD9F1/t51.2885-19/44884218_345707102882519_2446069589734326272_n.jpg?_nc_ht=scontent-sjc3-1.cdninstagram.com",
					},
				},

				Tags: []string{
					"0219test",
				},

				UsersInPhoto: []UserPosition{
					{
						User: &User{
							Username: "rcarver",
						},
						Position: &Position{
							X: 0.57568438,
							Y: 0.7938808374,
						},
					},
				},

				Location: &Location{
					Id:        float64(2.13051194e+08),
					Name:      "Oakland, California",
					Latitude:  37.8029,
					Longitude: -122.2721,
				},

				Images: &Images{
					Thumbnail: &Image{
						Url:    "https://scontent.cdninstagram.com/vp/fd0f484647ad37dc3caf0a2cdf37ca16/5CE59582/t51.2885-15/e35/c0.135.1080.1080/s150x150/50552544_116846169429307_872782777322498633_n.jpg?_nc_ht=scontent.cdninstagram.com",
						Width:  150,
						Height: 150,
					},
					LowResolution: &Image{
						Url:    "https://scontent.cdninstagram.com/vp/0eda6589295b6fa43fd5cf2731afd691/5CF9331A/t51.2885-15/e35/p320x320/50552544_116846169429307_872782777322498633_n.jpg?_nc_ht=scontent.cdninstagram.com",
						Width:  320,
						Height: 400,
					},
					StandardResolution: &Image{
						Url:    "https://scontent.cdninstagram.com/vp/bd6167c8e4469e16f2f6c900a62c51b9/5CF7EFF6/t51.2885-15/sh0.08/e35/p640x640/50552544_116846169429307_872782777322498633_n.jpg?_nc_ht=scontent.cdninstagram.com",
						Width:  640,
						Height: 800,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			api, server := newTestAPIServer(t, func(r *http.Request) string {
				return tt.fix
			})
			defer server.Close()

			ctx := context.Background()
			resp, err := api.GetRecentMedia(ctx, nil)
			if err != nil {
				t.Fatalf("GetRecentMedia: %s", err)
			}

			if got, want := len(resp.Medias), 1; got != want {
				t.Fatalf("Want %d media back, got %d", want, got)
			}
			got := resp.Medias[0]

			if !reflect.DeepEqual(got, tt.wantMedia) {
				t.Errorf("Medias not equal\nhave: %#v\nwant: %#v", got, tt.wantMedia)
				pretty.Ldiff(t, got, tt.wantMedia)
			}
		})
	}
}
