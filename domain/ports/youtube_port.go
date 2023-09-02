//go:generate mockgen -destination=../../mocks/ports/mock_youtube_port.go -package=mocks github.com/sugar-cat7/vspo-common-api/domain/ports YouTubeService
package ports

import (
	"google.golang.org/api/youtube/v3"
)

// YouTubeService is an interface for a YouTube implementation of a song service.
type YouTubeService interface {
	GetVideos(videoIDs []string) ([]*youtube.Video, error)
	GetPlaylists(playlistIDs []string) ([]*youtube.PlaylistItemContentDetails, error)
	GetChannels(channelIDs []string) ([]*youtube.Channel, error)
}
