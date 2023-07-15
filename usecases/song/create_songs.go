package usecases

import (
	"github.com/sugar-cat7/vspo-common-api/domain/services"
	"github.com/sugar-cat7/vspo-common-api/util"
)

// CreateSong is a use case for creating songs in Firestore from YouTube.
type CreateSong struct {
	youtubeService services.YouTubeService
	songService    services.SongService
}

// NewCreateSong creates a new CreateSong.
func NewCreateSong(youtubeService services.YouTubeService, songService services.SongService) *CreateSong {
	return &CreateSong{
		youtubeService: youtubeService,
		songService:    songService,
	}
}

// Execute creates new songs in Firestore from YouTube.
func (c *CreateSong) Execute() error {
	// Fetch playlist data from YouTube API
	playlistsData, err := c.youtubeService.GetPlaylists()
	if err != nil {
		return err
	}

	// Extract video IDs from the playlists
	extractedVideoIDs := make(map[string]struct{})
	for _, playlist := range playlistsData {
		for _, item := range playlist.Items {
			extractedVideoIDs[item.ResourceID.VideoID] = struct{}{}
		}
	}

	// Get all song IDs from Firestore
	existingVideoIDs, err := c.songService.GetSongIDs()
	if err != nil {
		return err
	}

	// Convert existingVideoIDs slice to map for efficient lookup
	existingVideoIDsMap := make(map[string]struct{})
	for _, id := range existingVideoIDs {
		existingVideoIDsMap[id] = struct{}{}
	}

	// Extract video IDs not present in Firestore
	newVideoIDs := []string{}
	for id := range extractedVideoIDs {
		if _, exists := existingVideoIDsMap[id]; !exists {
			newVideoIDs = append(newVideoIDs, id)
		}
	}

	// Fetch video data for new videos from YouTube API
	newVideoData, err := c.youtubeService.GetSongs(newVideoIDs)
	if err != nil {
		return err
	}

	// Convert the video data to Song models
	newSongs, err := util.ConvertToSongs(newVideoData)
	if err != nil {
		return err
	}

	// Save the new songs to Firestore
	err = c.songService.CreateSongsInBatch(newSongs)
	if err != nil {
		return err
	}

	return nil
}
