package factories

import (
	"time"

	"github.com/sugar-cat7/vspo-common-api/domain/entities"
)

func NewYTVideoListResponse(videoID string) entities.YTVideoListResponse {
	videoListResponse := entities.YTVideoListResponse{
		Kind: "youtube#videoListResponse",
		Etag: "etag1",
		Items: []entities.YTVideo{
			{
				Kind: "youtube#video",
				Etag: "etag2",
				ID:   videoID,
				Snippet: entities.YTVideoSnippet{
					PublishedAt: "2023-01-01T00:00:00Z",
					ChannelID:   "channelID1",
					Title:       "title1",
					Description: "description1",
					Thumbnails: entities.YTThumbnails{
						Default: entities.YTThumbnail{
							URL:    "https://example.com/default.jpg",
							Width:  120,
							Height: 90,
						},
						Medium: entities.YTThumbnail{
							URL:    "https://example.com/medium.jpg",
							Width:  320,
							Height: 180,
						},
						High: entities.YTThumbnail{
							URL:    "https://example.com/high.jpg",
							Width:  480,
							Height: 360,
						},
						Standard: entities.YTThumbnail{
							URL:    "https://example.com/standard.jpg",
							Width:  640,
							Height: 480,
						},
						Maxres: entities.YTThumbnail{
							URL:    "https://example.com/maxres.jpg",
							Width:  1280,
							Height: 720,
						},
					},
					ChannelTitle: "channelTitle1",
					Tags:         []string{"tag1", "tag2"},
					CategoryID:   "categoryID1",
				},
				Statistics: entities.YTStatistics{
					ViewCount:     "1000",
					LikeCount:     "500",
					FavoriteCount: "200",
					CommentCount:  "100",
				},
			},
		},
	}

	return videoListResponse
}

func NewYTPlayListListResponse(videoID string) entities.YTYouTubePlaylistResponse {
	playlist := entities.YTYouTubePlaylistResponse{
		Kind: "youtube#playlistResponse",
		Items: []struct {
			entities.YTPlayListSnippet `json:"snippet"`
		}{
			{
				YTPlayListSnippet: entities.YTPlayListSnippet{
					PublishedAt: "2023-01-01T00:00:00Z",
					ChannelID:   "channelID1",
					Title:       "title1",
					Description: "description1",
					Thumbnails: entities.YTThumbnails{
						Default: entities.YTThumbnail{
							URL:    "https://example.com/default.jpg",
							Width:  120,
							Height: 90,
						},
						Medium: entities.YTThumbnail{
							URL:    "https://example.com/medium.jpg",
							Width:  320,
							Height: 180,
						},
						High: entities.YTThumbnail{
							URL:    "https://example.com/high.jpg",
							Width:  480,
							Height: 360,
						},
						Standard: entities.YTThumbnail{
							URL:    "https://example.com/standard.jpg",
							Width:  640,
							Height: 480,
						},
						Maxres: entities.YTThumbnail{
							URL:    "https://example.com/maxres.jpg",
							Width:  1280,
							Height: 720,
						},
					},
					ChannelTitle:           "channelTitle1",
					PlaylistID:             "playlistID1",
					Position:               1,
					ResourceID:             entities.YTResourceID{Kind: "youtube#video", VideoID: videoID},
					VideoOwnerChannelTitle: "videoOwnerChannelTitle1",
					VideoOwnerChannelID:    "videoOwnerChannelID1",
				},
			},
		},
		PageInfo: entities.YTPageInfo{
			TotalResults:   1,
			ResultsPerPage: 1,
		},
	}

	return playlist
}

func NewSong(videoID string) entities.Song {
	return entities.Song{
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
