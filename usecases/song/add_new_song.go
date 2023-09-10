package usecases

import (
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/domain/ports"
	"github.com/sugar-cat7/vspo-common-api/domain/repositories"
	"github.com/sugar-cat7/vspo-common-api/usecases/mappers"
)

type AddNewSong struct {
	youtubeService ports.YouTubeService
	songRepository repositories.SongRepository
}

func NewAddNewSong(youtubeService ports.YouTubeService, songRepository repositories.SongRepository) *AddNewSong {
	return &AddNewSong{
		youtubeService: youtubeService,
		songRepository: songRepository,
	}
}

func (c *AddNewSong) Execute(playlistIDs []string) (entities.Videos, error) {

	playList, err := c.youtubeService.GetPlaylists(playlistIDs)
	if err != nil {
		return nil, err
	}

	exVideos, err := c.songRepository.GetAll()
	if err != nil {
		return nil, err
	}

	existVideoIDs := make([]string, 0, len(exVideos))
	for _, exVideo := range exVideos {
		existVideoIDs = append(existVideoIDs, exVideo.ID)
	}

	videoIDs := make([]string, 0, len(playList))
	for _, playlist := range playList {

		exists := false
		for _, existVideoID := range existVideoIDs {
			if playlist.VideoId == existVideoID {
				exists = true
				break
			}
		}

		if exists {
			continue
		}

		videoIDs = append(videoIDs, playlist.VideoId)
	}
	// Fetch video data from YouTube API
	videos, err := c.youtubeService.GetVideos(videoIDs)
	if err != nil {
		return nil, err
	}

	// Map the video data to Song models
	songs := mappers.MapToVideos(entities.None, videos)

	// Save the new songs to Firestore
	err = c.songRepository.CreateInBatch(songs)
	if err != nil {
		return nil, err
	}

	return songs, nil
}
