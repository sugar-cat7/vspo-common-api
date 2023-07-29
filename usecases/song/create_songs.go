package usecases

import (
	"github.com/sugar-cat7/vspo-common-api/domain/services"
	"github.com/sugar-cat7/vspo-common-api/usecases/mappers"
)

type CreateSong struct {
	youtubeService services.YouTubeService
	songService    services.SongService
	songMapper     *mappers.SongMapper
}

func NewCreateSong(youtubeService services.YouTubeService, songService services.SongService, songMapper *mappers.SongMapper) *CreateSong {
	return &CreateSong{
		youtubeService: youtubeService,
		songService:    songService,
		songMapper:     songMapper,
	}
}

func (c *CreateSong) Execute(videoIDs []string) error {
	// Fetch video data from YouTube API
	videos, err := c.youtubeService.GetVideos(videoIDs)
	if err != nil {
		return err
	}

	// Map the video data to Song models
	songs, err := c.songMapper.MapMultiple(videos)
	if err != nil {
		return err
	}

	// Save the new songs to Firestore
	err = c.songService.CreateSongsInBatch(songs)
	if err != nil {
		return err
	}

	return nil
}
