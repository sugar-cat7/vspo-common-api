package usecases

import (
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/domain/services"
	"github.com/sugar-cat7/vspo-common-api/usecases/mappers"
)

type AddNewSong struct {
	youtubeService services.YouTubeService
	songService    services.SongService
	songMapper     *mappers.SongMapper
}

func NewAddNewSong(youtubeService services.YouTubeService, songService services.SongService, songMapper *mappers.SongMapper) *AddNewSong {
	return &AddNewSong{
		youtubeService: youtubeService,
		songService:    songService,
		songMapper:     songMapper,
	}
}

func (c *AddNewSong) Execute(playlistIDs []string) ([]*entities.Video, error) {
	exsistVideoIds, err := c.songService.GetSongIDs()
	if err != nil {
		return nil, err
	}
	playList, err := c.youtubeService.GetPlaylists(playlistIDs)
	if err != nil {
		return nil, err
	}

	videoIDs := make([]string, 0)
	for _, playlist := range playList {

		exists := false
		for _, existVideoId := range exsistVideoIds {
			if playlist.VideoId == existVideoId {
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
	songs, err := c.songMapper.MapMultiple(videos)
	if err != nil {
		return nil, err
	}

	// Save the new songs to Firestore
	err = c.songService.CreateSongsInBatch(songs)
	if err != nil {
		return nil, err
	}

	return songs, nil
}
