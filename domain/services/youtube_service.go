//go:generate mockgen -destination=../../mocks/services/mock_youtube_service.go -package=mocks github.com/sugar-cat7/vspo-common-api/domain/services YouTubeService
package services

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/sugar-cat7/vspo-common-api/util"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

// YouTubeService is an interface for a YouTube implementation of a song service.
type YouTubeService interface {
	GetVideos(videoIDs []string) ([]*youtube.Video, error)
	GetPlaylists(playlistIDs []string) ([]*youtube.PlaylistItemContentDetails, error)
	GetChannels(channelIDs []string) ([]*youtube.Channel, error)
}

type youtubeServiceImpl struct {
	Service *youtube.Service
}

// NewYouTubeService creates a new YouTubeService.
func NewYouTubeService() (YouTubeService, error) {
	apiKey, ok := os.LookupEnv("YOUTUBE_API_KEY")
	if !ok {
		return nil, fmt.Errorf("YOUTUBE_API_KEY not set")
	}

	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	return &youtubeServiceImpl{Service: service}, nil
}

const chunkSize = 50

func (s *youtubeServiceImpl) getChunks(ids []string) ([][]string, error) {
	chunks, err := util.Chunk(ids, chunkSize)
	if err != nil {
		return nil, fmt.Errorf("error splitting ids into chunks: %v", err)
	}
	return chunks, nil
}

func (s *youtubeServiceImpl) GetVideos(videoIDs []string) ([]*youtube.Video, error) {
	var videos []*youtube.Video
	videoIDChunks, err := s.getChunks(videoIDs)
	if err != nil {
		return nil, err
	}
	for _, chunk := range videoIDChunks {
		call := s.Service.Videos.List([]string{"snippet", "liveStreamingDetails", "statistics"}).Id(strings.Join(chunk, ","))
		response, err := call.Do()
		if err != nil {
			return nil, fmt.Errorf("error making Videos.List call: %v", err)
		}
		videos = append(videos, response.Items...)
	}
	return videos, nil
}

func (s *youtubeServiceImpl) GetPlaylists(playlistIDs []string) ([]*youtube.PlaylistItemContentDetails, error) {
	var contentDetails []*youtube.PlaylistItemContentDetails
	for _, playlistID := range playlistIDs {
		call := s.Service.PlaylistItems.List([]string{"contentDetails"}).PlaylistId(playlistID).MaxResults(50)
		response, err := call.Do()
		if err != nil {
			return nil, fmt.Errorf("error making PlaylistItems.List call for playlist ID %v: %v", playlistID, err)
		}

		for _, item := range response.Items {
			contentDetails = append(contentDetails, item.ContentDetails)
		}
	}
	return contentDetails, nil
}

func (s *youtubeServiceImpl) GetChannels(channelIDs []string) ([]*youtube.Channel, error) {
	var channels []*youtube.Channel
	channelIDChunks, err := s.getChunks(channelIDs)
	if err != nil {
		return nil, err
	}
	for _, chunk := range channelIDChunks {
		call := s.Service.Channels.List([]string{"snippet", "statistics"}).Id(strings.Join(chunk, ","))
		response, err := call.Do()
		if err != nil {
			return nil, fmt.Errorf("error making Channels.List call: %v", err)
		}
		channels = append(channels, response.Items...)
	}
	return channels, nil
}
