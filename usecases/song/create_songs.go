package usecases

import (
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/domain/ports"
	"github.com/sugar-cat7/vspo-common-api/domain/repositories"
	"github.com/sugar-cat7/vspo-common-api/usecases/mappers"
)

type CreateSong struct {
	youtubeService ports.YouTubeService
	songRepository repositories.SongRepository
}

func NewCreateSong(youtubeService ports.YouTubeService, songRepository repositories.SongRepository) *CreateSong {
	return &CreateSong{
		youtubeService: youtubeService,
		songRepository: songRepository,
	}
}

func (c *CreateSong) Execute(videoIDs []string) error {
	// Fetch video data from YouTube API
	videos, err := c.youtubeService.GetVideos(videoIDs)
	if err != nil {
		return err
	}

	// Map the video data to Song models
	songs := mappers.MapToVideos(entities.None, videos)
	if err != nil {
		return err
	}

	// Save the new songs to Firestore
	err = c.songRepository.CreateInBatch(songs)
	if err != nil {
		return err
	}

	return nil
}
