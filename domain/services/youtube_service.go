//go:generate mockgen -destination=../../mocks/services/mock_youtube_service.go -package=mocks github.com/sugar-cat7/vspo-common-api/domain/services YouTubeService
package services

import (
	"net/http"

	"github.com/sugar-cat7/vspo-common-api/domain/entities"

	"github.com/sugar-cat7/vspo-common-api/infrastructure/youtube"
)

// YouTubeService is an interface for a YouTube implementation of a song service.
type YouTubeService interface {
	GetSongs(videoIDs []string) ([]entities.YTVideoListResponse, error)
	GetPlaylists() ([]entities.YTYouTubePlaylistResponse, error)
}

// youtubeServiceImpl is a YouTube implementation of a song service.
type youtubeServiceImpl struct {
	API *youtube.API
}

// NewYouTubeService creates a new YouTubeService.
func NewYouTubeService(client *http.Client) YouTubeService {
	api := youtube.NewAPI(client)
	return &youtubeServiceImpl{API: api}
}

// GetSongs returns a slice of Song models.
func (s *youtubeServiceImpl) GetSongs(videoIDs []string) ([]entities.YTVideoListResponse, error) {
	// Fetch video data from YouTube API
	videoData, err := s.API.GetVideos(videoIDs)
	if err != nil {
		return nil, err
	}

	return videoData, nil
}

// GetPlaylists returns a slice of YTYouTubePlaylistResponse models.
func (s *youtubeServiceImpl) GetPlaylists() ([]entities.YTYouTubePlaylistResponse, error) {
	return s.API.GetPlaylists()
}
