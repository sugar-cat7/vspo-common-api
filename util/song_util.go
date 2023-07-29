package util

import (
	"strconv"

	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"google.golang.org/api/youtube/v3"
)

// UpdateViewCounts updates the view counts of songs based on the YoutubeVideoListResponse data
func UpdateViewCounts(cronType entities.CronType, videos []*youtube.Video, songs []*entities.Song) ([]*entities.Song, error) {
	videoMap := make(map[string]*youtube.Video)
	for _, video := range videos {
		videoMap[video.Id] = video
	}

	// Initialize updatedSongs to store updated songs
	updatedSongs := make([]*entities.Song, 0)

	for _, song := range songs {
		video, exists := videoMap[song.ID]
		if exists {
			// Update the views
			switch cronType {
			case entities.Daily:
				song.ViewCount.Daily = strconv.FormatUint(video.Statistics.ViewCount, 10)
			case entities.Weekly:
				song.ViewCount.Weekly = strconv.FormatUint(video.Statistics.ViewCount, 10)
			case entities.Monthly:
				song.ViewCount.Monthly = strconv.FormatUint(video.Statistics.ViewCount, 10)
			}

			song.ViewCount.Total = strconv.FormatUint(video.Statistics.ViewCount, 10)
			updatedSongs = append(updatedSongs, song)
		}
	}

	return updatedSongs, nil
}

// GetSongIDs returns a slice of song IDs from a slice of songs.
func GetSongIDs(songs []*entities.Song) []string {
	ids := make([]string, len(songs))
	for i, song := range songs {
		ids[i] = song.ID
	}
	return ids
}
