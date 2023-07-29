package usecases

import (
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/domain/services"
	"github.com/sugar-cat7/vspo-common-api/usecases/mappers"
	"github.com/sugar-cat7/vspo-common-api/util"
)

// UpdateSongs is a use case for updating songs in Firestore from YouTube.
type UpdateSongs struct {
	youtubeService services.YouTubeService
	songService    services.SongService
	songMapper     *mappers.SongMapper
}

// NewUpdateSongs creates a new UpdateSongs.
func NewUpdateSongs(youtubeService services.YouTubeService, songService services.SongService, songMapper *mappers.SongMapper) *UpdateSongs {
	return &UpdateSongs{
		youtubeService: youtubeService,
		songService:    songService,
		songMapper:     songMapper,
	}
}

// Execute updates the songs in Firestore from YouTube.
func (u *UpdateSongs) Execute(cronType entities.CronType) error {
	// Get all song IDs from Firestore
	allVideos, err := u.songService.GetAllSongs()
	if err != nil {
		return err
	}

	videoIDs := util.GetSongIDs(allVideos)
	// Fetch video data from YouTube API
	ytVideos, err := u.youtubeService.GetVideos(videoIDs)
	if err != nil {
		return err
	}

	// Convert the video data to Song models
	updatedSongs, err := util.UpdateViewCounts(cronType, ytVideos, allVideos)
	if err != nil {
		return err
	}

	// Update the songs in Firestore
	err = u.songService.UpdateSongsInBatch(updatedSongs)
	if err != nil {
		return err
	}

	return nil
}
