package factories

import (
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	entities2 "github.com/sugar-cat7/vspo-common-api/domain/entities/legacy"
	"google.golang.org/api/youtube/v3"
)

func NewVideoPtr(videoID string) *entities.Video {
	video := NewVideo(videoID)
	return &video
}

func NewVideo(videoID string) entities.Video {
	return entities.Video{
		ID:          videoID,
		Title:       "title1",
		Description: "description1",
		ViewCount: entities.Views{
			Total: "1000",
		},
		PublishedAt: time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
		Thumbnails: entities.Thumbnails{
			Default: entities.Thumbnail{
				URL:    "https://example.com/default.jpg",
				Width:  120,
				Height: 90,
			},
			Medium: entities.Thumbnail{
				URL:    "https://example.com/medium.jpg",
				Width:  320,
				Height: 180,
			},
			High: entities.Thumbnail{
				URL:    "https://example.com/high.jpg",
				Width:  480,
				Height: 360,
			},
			Standard: entities.Thumbnail{
				URL:    "https://example.com/standard.jpg",
				Width:  640,
				Height: 480,
			},
			Maxres: entities.Thumbnail{
				URL:    "https://example.com/maxres.jpg",
				Width:  1280,
				Height: 720,
			},
		},
		ChannelTitle: "channelTitle1",
		ChannelID:    "channelID1",
		Tags:         []string{"tag1", "tag2"},
	}
}

func NewChannel(channelID string) entities.Channel {
	return entities.Channel{
		ID: channelID,
		Snippet: entities.ChannelSnippet{
			Title:       "channelTitle1",
			Description: "description1",
			CustomURL:   "https://example.com/" + channelID,
			PublishedAt: time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
			Thumbnails: entities.Thumbnails{
				Default: entities.Thumbnail{
					URL:    "https://example.com/default.jpg",
					Width:  120,
					Height: 90,
				},
				Medium: entities.Thumbnail{
					URL:    "https://example.com/medium.jpg",
					Width:  320,
					Height: 180,
				},
				High: entities.Thumbnail{
					URL:    "https://example.com/high.jpg",
					Width:  480,
					Height: 360,
				},
				Standard: entities.Thumbnail{
					URL:    "https://example.com/standard.jpg",
					Width:  640,
					Height: 480,
				},
				Maxres: entities.Thumbnail{
					URL:    "https://example.com/maxres.jpg",
					Width:  1280,
					Height: 720,
				},
			},
		},
		Statistics: entities.ChannelStatistics{
			ViewCount:             "10000",
			SubscriberCount:       "5000",
			HiddenSubscriberCount: false,
			VideoCount:            "200",
		},
	}
}

// NewYoutubeVideo creates a new mock youtube.Video instance with a given ID and random data for the other fields
func NewYoutubeVideo(id string) *youtube.Video {
	viewCount := gofakeit.Number(1000000000, 9999999999) // generate a random 10-digit number for the view count
	return &youtube.Video{
		Id: id,
		Snippet: &youtube.VideoSnippet{
			Title:        gofakeit.Sentence(5),
			Description:  gofakeit.Paragraph(5, 10, 25, " "),
			ChannelId:    gofakeit.UUID(),
			ChannelTitle: gofakeit.Word(),
			CategoryId:   strconv.Itoa(gofakeit.Number(10, 99)), // generate a random 2-digit number for the category ID
			PublishedAt:  time.Now().Format(time.RFC3339),
			Thumbnails: &youtube.ThumbnailDetails{
				Default: &youtube.Thumbnail{
					Url:    gofakeit.URL(),
					Width:  120,
					Height: 90,
				},
				Medium: &youtube.Thumbnail{
					Url:    gofakeit.URL(),
					Width:  320,
					Height: 180,
				},
				High: &youtube.Thumbnail{
					Url:    gofakeit.URL(),
					Width:  480,
					Height: 360,
				},
				Standard: &youtube.Thumbnail{
					Url:    gofakeit.URL(),
					Width:  640,
					Height: 480,
				},
				Maxres: &youtube.Thumbnail{
					Url:    gofakeit.URL(),
					Width:  1280,
					Height: 720,
				},
			},
			Tags: []string{gofakeit.Word(), gofakeit.Word(), gofakeit.Word()},
		},
		ContentDetails: &youtube.VideoContentDetails{
			Duration: strconv.Itoa(gofakeit.Number(100, 999)), // generate a random 3-digit number for the duration
		},
		Statistics: &youtube.VideoStatistics{
			ViewCount: uint64(viewCount),
		},
	}
}

// NewYoutubeChannel creates a new mock youtube.Channel instance with a given ID and random data for the other fields
func NewYoutubeChannel(id string) *youtube.Channel {
	viewCount := gofakeit.Number(1000000000, 9999999999) // generate a random 10-digit number for the view count
	subscriberCount := gofakeit.Number(1000000, 9999999) // generate a random 7-digit number for the subscriber count
	videoCount := gofakeit.Number(100, 999)              // generate a random 3-digit number for the video count

	return &youtube.Channel{
		Id: id,
		Snippet: &youtube.ChannelSnippet{
			Title:       gofakeit.Sentence(5),
			Description: gofakeit.Paragraph(5, 10, 25, " "),
			CustomUrl:   gofakeit.URL(),
			PublishedAt: time.Now().Format(time.RFC3339),
			Thumbnails: &youtube.ThumbnailDetails{
				Default: &youtube.Thumbnail{
					Url:    gofakeit.ImageURL(120, 90),
					Width:  120,
					Height: 90,
				},
				Medium: &youtube.Thumbnail{
					Url:    gofakeit.ImageURL(320, 180),
					Width:  320,
					Height: 180,
				},
				High: &youtube.Thumbnail{
					Url:    gofakeit.ImageURL(480, 360),
					Width:  480,
					Height: 360,
				},
				Standard: &youtube.Thumbnail{
					Url:    gofakeit.ImageURL(640, 480),
					Width:  640,
					Height: 480,
				},
				Maxres: &youtube.Thumbnail{
					Url:    gofakeit.ImageURL(1280, 720),
					Width:  1280,
					Height: 720,
				},
			},
		},
		Statistics: &youtube.ChannelStatistics{
			ViewCount:       uint64(viewCount),
			SubscriberCount: uint64(subscriberCount),
			VideoCount:      uint64(videoCount),
		},
	}
}

func NewClip(clipID string) *entities2.OldVideo {
	return &entities2.OldVideo{
		ID:           clipID,
		Title:        "title1",
		Description:  "description1",
		ChannelID:    "channelID1",
		ChannelTitle: "channelTitle1",
		ThumbnailURL: "https://example.com/default.jpg",
		IconURL:      "https://example.com/icon.jpg",
		Platform:     "youtube",
		ViewCount:    "1000",
		// LikeCount:    "500",
		CommentCount: "200",
		NewViewCount: entities.Views{
			Total: "1000",
		},
		CreatedAt: time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
	}
}

func NewLiveStream(libestreamID string) *entities2.OldVideo {
	return &entities2.OldVideo{
		ID:           libestreamID,
		Title:        "title1",
		Description:  "description1",
		ChannelID:    "channelID1",
		ChannelTitle: "channelTitle1",
		ThumbnailURL: "https://i.ytimg.com/vi/IWYydmhCA8o/hqdefault.jpg?sqp=-oaymwEcCNACELwBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLDz6fev5X2b3gpB-QRyvP3Zue4phQ",
		IconURL:      "https://yt3.googleusercontent.com/ytc/AOPolaRJN3duCEREvTCkefZB05nmi5s6y7pd1f19exq_tA=s176-c-k-c0x00ffffff-no-rj",
		Platform:     "youtube",
		ViewCount:    "1000",
		// LikeCount:    "500",
		CommentCount: "200",
		NewViewCount: entities.Views{
			Total: "1000",
		},
		CreatedAt:          time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
		ScheduledStartTime: time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
		ActualEndTime:      time.Date(2023, 01, 02, 0, 0, 0, 0, time.UTC),
	}
}
