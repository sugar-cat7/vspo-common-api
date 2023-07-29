package usecases

import (
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/domain/services"
	"github.com/sugar-cat7/vspo-common-api/util"
)

// UpdateSongs is a use case for updating songs in Firestore from YouTube.
type UpdateSongs struct {
	youtubeService services.YouTubeService
	songService    services.SongService
}

// NewUpdateSongs creates a new UpdateSongs.
func NewUpdateSongs(youtubeService services.YouTubeService, songService services.SongService) *UpdateSongs {
	return &UpdateSongs{
		youtubeService: youtubeService,
		songService:    songService,
	}
}

// Execute updates the songs in Firestore from YouTube.
func (u *UpdateSongs) Execute(cronType entities.CronType) error {
	// Get all song IDs from Firestore
	videos, err := u.songService.GetAllSongs()
	if err != nil {
		return err
	}

	videoIDs := util.GetSongIDs(videos)
	// Fetch video data from YouTube API
	videoData, err := u.youtubeService.GetSongs(videoIDs)
	if err != nil {
		return err
	}

	// Convert the video data to Song models
	err = util.UpdateViewCounts(cronType, videoData, videos)
	if err != nil {
		return err
	}

	// Update the songs in Firestore
	err = u.songService.UpdateSongsInBatch(videos)
	if err != nil {
		return err
	}

	return nil
}
