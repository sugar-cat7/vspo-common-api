package usecases

import (
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/domain/ports"
	"github.com/sugar-cat7/vspo-common-api/domain/repositories"
	"github.com/sugar-cat7/vspo-common-api/usecases/mappers"
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

	videoIDs := allVideos.GetIDs()
	// Fetch video data from YouTube API
	ytVideos, err := u.youtubeService.GetVideos(videoIDs)
	if err != nil {
		return err
	}
	mappedVideos := mappers.MapToVideos(cronType, ytVideos)

	// Update the songs in Firestore
	err = u.songRepository.UpdateInBatch(mappedVideos)
	if err != nil {
		return err
	}

	return nil
}
