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

				User: User{
					Id:             "11073382793",
					Username:       "go_ig_test_0219",
					FullName:       "Golang Client",
					ProfilePicture: "https://scontent-sjc3-1.cdninstagram.com/vp/504ac2fa79adb1d412b31cab19be8d36/5CDDD9F1/t51.2885-19/44884218_345707102882519_2446069589734326272_n.jpg?_nc_ht=scontent-sjc3-1.cdninstagram.com",
				},

				Likes: Likes{
					Count: 1,
				},

				Comments: Comments{
					Count: 2,
				},

				Caption: Caption{
					Id:          "18002756710177046",
					CreatedTime: "1550173420",
					Text:        "Photo post #0219test",
					From: User{
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
						User: User{
							Username: "rcarver",
						},
						Position: Position{
							X: 0.57568438,
							Y: 0.7938808374,
						},
					},
				},

				Location: Location{
					Id:        float64(2.13051194e+08),
					Name:      "Oakland, California",
					Latitude:  37.8029,
					Longitude: -122.2721,
				},

				Images: Images{
					Thumbnail: Image{
						Url:    "https://scontent.cdninstagram.com/vp/fd0f484647ad37dc3caf0a2cdf37ca16/5CE59582/t51.2885-15/e35/c0.135.1080.1080/s150x150/50552544_116846169429307_872782777322498633_n.jpg?_nc_ht=scontent.cdninstagram.com",
						Width:  150,
						Height: 150,
					},
					LowResolution: Image{
						Url:    "https://scontent.cdninstagram.com/vp/0eda6589295b6fa43fd5cf2731afd691/5CF9331A/t51.2885-15/e35/p320x320/50552544_116846169429307_872782777322498633_n.jpg?_nc_ht=scontent.cdninstagram.com",
						Width:  320,
						Height: 400,
					},
					StandardResolution: Image{
						Url:    "https://scontent.cdninstagram.com/vp/bd6167c8e4469e16f2f6c900a62c51b9/5CF7EFF6/t51.2885-15/sh0.08/e35/p640x640/50552544_116846169429307_872782777322498633_n.jpg?_nc_ht=scontent.cdninstagram.com",
						Width:  640,
						Height: 800,
					},
				},
			},
		},
		{
			desc: "video media",
			fix:  "video",
			wantMedia: Media{
				Id:          "1979318157757411422_11073382793",
				Type:        "video",
				CreatedTime: "1550173160",

				Link:   "https://www.instagram.com/p/Bt382CqnHBe-UvbBQ78RvFycVFM2JDGVrd5Xfs0/",
				Filter: "Normal",

				User: User{
					Id:             "11073382793",
					Username:       "go_ig_test_0219",
					FullName:       "Golang Client",
					ProfilePicture: "https://scontent-arn2-1.cdninstagram.com/vp/10dc8a532f753c991ce068347dfc0767/5CDDD9F1/t51.2885-19/44884218_345707102882519_2446069589734326272_n.jpg?_nc_ht=scontent-arn2-1.cdninstagram.com",
				},

				Caption: Caption{
					Id:          "17878115845306299",
					CreatedTime: "1550173160",
					Text:        "Video post #0219test #videotesr",
					From: User{
						Id:             "11073382793",
						Username:       "go_ig_test_0219",
						FullName:       "Golang Client",
						ProfilePicture: "https://scontent-arn2-1.cdninstagram.com/vp/10dc8a532f753c991ce068347dfc0767/5CDDD9F1/t51.2885-19/44884218_345707102882519_2446069589734326272_n.jpg?_nc_ht=scontent-arn2-1.cdninstagram.com",
					},
				},

				Tags: []string{
					"0219test",
					"videotesr",
				},

				UsersInPhoto: []UserPosition{},

				Images: Images{
					Thumbnail: Image{
						Url:    "https://scontent.cdninstagram.com/vp/5fe061d53fd0607be846f3c5409f603a/5C683C08/t51.2885-15/e15/s150x150/50863106_2133926990002507_2324727490466410627_n.jpg?_nc_ht=scontent.cdninstagram.com",
						Width:  150,
						Height: 150,
					},
					LowResolution: Image{
						Url:    "https://scontent.cdninstagram.com/vp/6b86ebfdd4d9055e95cb392e4bba6b27/5C681EF0/t51.2885-15/e15/s320x320/50863106_2133926990002507_2324727490466410627_n.jpg?_nc_ht=scontent.cdninstagram.com",
						Width:  320,
						Height: 320,
					},
					StandardResolution: Image{
						Url:    "https://scontent.cdninstagram.com/vp/3780985d3fe523a7189a8b5d26d91b82/5C68240B/t51.2885-15/e15/s640x640/50863106_2133926990002507_2324727490466410627_n.jpg?_nc_ht=scontent.cdninstagram.com",
						Width:  640,
						Height: 640,
					},
				},

				Videos: Images{
					LowResolution: Image{
						Url:    "https://scontent.cdninstagram.com/vp/c56a21cc014e65e9e22fab2f0f25d60e/5C6890AD/t50.2886-16/52831800_292210298079609_3680295379305234432_n.mp4?_nc_ht=scontent.cdninstagram.com",
						Id:     "17848556476361628",
						Width:  480,
						Height: 480,
					},
					LowBandwidth: Image{
						Url:    "https://scontent.cdninstagram.com/vp/c56a21cc014e65e9e22fab2f0f25d60e/5C6890AD/t50.2886-16/52831800_292210298079609_3680295379305234432_n.mp4?_nc_ht=scontent.cdninstagram.com",
						Id:     "17848556476361628",
						Width:  480,
						Height: 480,
					},
					StandardResolution: Image{
						Url:    "https://scontent.cdninstagram.com/vp/fa26e894e4be773c121d92f4b1958c0b/5C6811FE/t50.2886-16/52696644_485271351879016_7856051086296088576_n.mp4?_nc_ht=scontent.cdninstagram.com",
						Id:     "17995277092175876",
						Width:  640,
						Height: 640,
					},
				},
			},
		},
		{
			desc: "carosel with image and video",
			fix:  "carousel_mixed",
			wantMedia: Media{
				Id:          "1979319391662961209_11073382793",
				Type:        "carousel",
				CreatedTime: "1550173280",

				Link:   "https://www.instagram.com/p/Bt39H_1HPI5YVhqU2BTNiI40_wBQerAXmEDoeU0/",
				Filter: "Normal",

				User: User{
					Id:             "11073382793",
					Username:       "go_ig_test_0219",
					FullName:       "Golang Client",
					ProfilePicture: "https://scontent-mia3-2.cdninstagram.com/vp/f7cc70d344dbbf503819a12b6da4800e/5CDDD9F1/t51.2885-19/44884218_345707102882519_2446069589734326272_n.jpg?_nc_ht=scontent-mia3-2.cdninstagram.com",
				},

				Caption: Caption{
					Id:          "18063037396017626",
					CreatedTime: "1550173280",
					Text:        "Carousel photo and video #0219test #carouseltest",
					From: User{
						Id:             "11073382793",
						Username:       "go_ig_test_0219",
						FullName:       "Golang Client",
						ProfilePicture: "https://scontent-mia3-2.cdninstagram.com/vp/f7cc70d344dbbf503819a12b6da4800e/5CDDD9F1/t51.2885-19/44884218_345707102882519_2446069589734326272_n.jpg?_nc_ht=scontent-mia3-2.cdninstagram.com",
					},
				},

				Tags: []string{
					"carouseltest",
					"0219test",
				},

				UsersInPhoto: []UserPosition{},

				Images: Images{
					Thumbnail: Image{
						Url:    "https://scontent.cdninstagram.com/vp/82e7dd474a6c6735614f8c306c0be3fb/5CDF80BF/t51.2885-15/e35/c0.0.1079.1079/s150x150/51287969_2359910680962588_7283258652285168337_n.jpg?_nc_ht=scontent.cdninstagram.com",
						Width:  150,
						Height: 150,
					},
					LowResolution: Image{
						Url:    "https://scontent.cdninstagram.com/vp/261dc6deb2eaa87210ee2675050dcf4c/5CDDEB8F/t51.2885-15/e35/s320x320/51287969_2359910680962588_7283258652285168337_n.jpg?_nc_ht=scontent.cdninstagram.com",
						Width:  320,
						Height: 319,
					},
					StandardResolution: Image{
						Url:    "https://scontent.cdninstagram.com/vp/df6922bf475acd96ae4b88787fddea0f/5CF61472/t51.2885-15/sh0.08/e35/s640x640/51287969_2359910680962588_7283258652285168337_n.jpg?_nc_ht=scontent.cdninstagram.com",
						Width:  640,
						Height: 639,
					},
				},

				CarouselMedias: []CarouselMedia{
					{
						Type: "image",
						Images: Images{
							Thumbnail: Image{
								Url:    "https://scontent.cdninstagram.com/vp/82e7dd474a6c6735614f8c306c0be3fb/5CDF80BF/t51.2885-15/e35/c0.0.1079.1079/s150x150/51287969_2359910680962588_7283258652285168337_n.jpg?_nc_ht=scontent.cdninstagram.com",
								Width:  150,
								Height: 150,
							},
							LowResolution: Image{
								Url:    "https://scontent.cdninstagram.com/vp/261dc6deb2eaa87210ee2675050dcf4c/5CDDEB8F/t51.2885-15/e35/s320x320/51287969_2359910680962588_7283258652285168337_n.jpg?_nc_ht=scontent.cdninstagram.com",
								Width:  320,
								Height: 319,
							},
							StandardResolution: Image{
								Url:    "https://scontent.cdninstagram.com/vp/df6922bf475acd96ae4b88787fddea0f/5CF61472/t51.2885-15/sh0.08/e35/s640x640/51287969_2359910680962588_7283258652285168337_n.jpg?_nc_ht=scontent.cdninstagram.com",
								Width:  640,
								Height: 639,
							},
						},
						UsersInPhoto: []UserPosition{},
					},
					{
						Type: "video",
						Videos: Images{
							LowResolution: Image{
								Id:     "17973623890206697",
								Url:    "https://scontent.cdninstagram.com/vp/289d214085a1d17234b1138781c0563b/5C685779/t50.2886-16/52133795_250694959199695_5368374984429273088_n.mp4?_nc_ht=scontent.cdninstagram.com",
								Width:  480,
								Height: 480,
							},
							LowBandwidth: Image{
								Id:     "17973623890206697",
								Url:    "https://scontent.cdninstagram.com/vp/289d214085a1d17234b1138781c0563b/5C685779/t50.2886-16/52133795_250694959199695_5368374984429273088_n.mp4?_nc_ht=scontent.cdninstagram.com",
								Width:  480,
								Height: 480,
							},
							StandardResolution: Image{
								Id:     "17866771429321861",
								Url:    "https://scontent.cdninstagram.com/vp/eead42a01be475b39c8099a066db344d/5C682502/t50.2886-16/52613010_405348530239650_6898777355045568512_n.mp4?_nc_ht=scontent.cdninstagram.com",
								Width:  640,
								Height: 640,
							},
						},
						UsersInPhoto: []UserPosition{},
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
