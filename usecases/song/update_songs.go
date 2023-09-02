package usecases

import (
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/domain/ports"
	"github.com/sugar-cat7/vspo-common-api/domain/repositories"
	"github.com/sugar-cat7/vspo-common-api/util"
)

// UpdateSongs is a use case for updating songs in Firestore from YouTube.
type UpdateSongs struct {
	youtubeService ports.YouTubeService
	songRepository repositories.SongRepository
}

// NewUpdateSongs creates a new UpdateSongs.
func NewUpdateSongs(youtubeService ports.YouTubeService, songRepository repositories.SongRepository) *UpdateSongs {
	return &UpdateSongs{
		youtubeService: youtubeService,
		songRepository: songRepository,
	}
}

// Execute updates the songs in Firestore from YouTube.
func (u *UpdateSongs) Execute(cronType entities.CronType) error {
	// Get all song IDs from Firestore
	allVideos, err := u.songRepository.GetAll()
	if err != nil {
		return err
	}

	videoIDs := util.GetVideoIDs(allVideos)
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
	err = u.songRepository.UpdateInBatch(updatedSongs)
	if err != nil {
		return err
	}

	return nil
}
