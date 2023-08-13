package mappers

import (
	"strconv"
	"time"

	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"google.golang.org/api/youtube/v3"
)

// SongMapper maps a YouTube API video to a domain Song.
type SongMapper struct{}

// Map maps a YouTube API video to a domain Song.
func (sm *SongMapper) Map(ytVideo *youtube.Video) (*entities.Video, error) {
	pubAt, err := time.Parse(time.RFC3339, ytVideo.Snippet.PublishedAt)
	if err != nil {
		return nil, err
	}
	viewCount := strconv.FormatUint(ytVideo.Statistics.ViewCount, 10)
	return &entities.Video{
		ID:          ytVideo.Id,
		Title:       ytVideo.Snippet.Title,
		Description: ytVideo.Snippet.Description,
		ViewCount: entities.Views{
			Daily:   viewCount,
			Weekly:  viewCount,
			Monthly: viewCount,
			Total:   viewCount,
		},
		PublishedAt: pubAt,
		Thumbnails: entities.Thumbnails{
			Default: entities.Thumbnail{
				URL:    ytVideo.Snippet.Thumbnails.Default.Url,
				Width:  int(ytVideo.Snippet.Thumbnails.Default.Width),
				Height: int(ytVideo.Snippet.Thumbnails.Default.Height),
			},
			Medium: entities.Thumbnail{
				URL:    ytVideo.Snippet.Thumbnails.Medium.Url,
				Width:  int(ytVideo.Snippet.Thumbnails.Medium.Width),
				Height: int(ytVideo.Snippet.Thumbnails.Medium.Height),
			},
			High: entities.Thumbnail{
				URL:    ytVideo.Snippet.Thumbnails.High.Url,
				Width:  int(ytVideo.Snippet.Thumbnails.High.Width),
				Height: int(ytVideo.Snippet.Thumbnails.High.Height),
			},
			// Standard: entities.Thumbnail{
			// 	URL:    ytVideo.Snippet.Thumbnails.Standard.Url,
			// 	Width:  int(ytVideo.Snippet.Thumbnails.Standard.Width),
			// 	Height: int(ytVideo.Snippet.Thumbnails.Standard.Height),
			// },
			// Maxres: entities.Thumbnail{
			// 	URL:    ytVideo.Snippet.Thumbnails.Maxres.Url,
			// 	Width:  int(ytVideo.Snippet.Thumbnails.Maxres.Width),
			// 	Height: int(ytVideo.Snippet.Thumbnails.Maxres.Height),
			// },
		},
		ChannelTitle: ytVideo.Snippet.ChannelTitle,
		ChannelID:    ytVideo.Snippet.ChannelId,
		Tags:         ytVideo.Snippet.Tags,
	}, nil
}

// MapMultiple maps multiple YouTube API videos to domain Songs.
func (sm *SongMapper) MapMultiple(ytVideos []*youtube.Video) ([]*entities.Video, error) {
	songs := make([]*entities.Video, len(ytVideos))
	for i, ytVideo := range ytVideos {
		song, err := sm.Map(ytVideo)
		if err != nil {
			return nil, err
		}
		songs[i] = song
	}

	return songs, nil
}
