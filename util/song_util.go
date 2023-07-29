package util

import (
	"fmt"
	"time"

	"github.com/sugar-cat7/vspo-common-api/domain/entities"
)

// ConvertToSongs converts a slice of YTVideoListResponses to a slice of Songs and updates the view count based on the provided cronType.
func ConvertToSongs(data []entities.YTVideoListResponse) ([]*entities.Song, error) {
	var songs []*entities.Song

	for _, responseData := range data {
		for _, video := range responseData.Items {
			publishedAt, err := time.Parse(time.RFC3339, video.Snippet.PublishedAt)
			if err != nil {
				return nil, fmt.Errorf("error parsing published date: %v", err)
			}
			song := &entities.Song{
				ID:          video.ID,
				Title:       video.Snippet.Title,
				Description: video.Snippet.Description,
				ViewCount: entities.Views{
					Total: video.Statistics.ViewCount,
				},
				PublishedAt: publishedAt,
				Thumbnails: entities.Thumbnails{
					Default: entities.Thumbnail{
						URL:    video.Snippet.Thumbnails.Default.URL,
						Width:  video.Snippet.Thumbnails.Default.Width,
						Height: video.Snippet.Thumbnails.Default.Height,
					},
					Medium: entities.Thumbnail{
						URL:    video.Snippet.Thumbnails.Medium.URL,
						Width:  video.Snippet.Thumbnails.Medium.Width,
						Height: video.Snippet.Thumbnails.Medium.Height,
					},
					High: entities.Thumbnail{
						URL:    video.Snippet.Thumbnails.High.URL,
						Width:  video.Snippet.Thumbnails.High.Width,
						Height: video.Snippet.Thumbnails.High.Height,
					},
					Standard: entities.Thumbnail{
						URL:    video.Snippet.Thumbnails.Standard.URL,
						Width:  video.Snippet.Thumbnails.Standard.Width,
						Height: video.Snippet.Thumbnails.Standard.Height,
					},
					Maxres: entities.Thumbnail{
						URL:    video.Snippet.Thumbnails.Maxres.URL,
						Width:  video.Snippet.Thumbnails.Maxres.Width,
						Height: video.Snippet.Thumbnails.Maxres.Height,
					},
				},
				ChannelTitle: video.Snippet.ChannelTitle,
				ChannelID:    video.Snippet.ChannelID,
				Tags:         video.Snippet.Tags,
			}
			// oldTotalViewCount, err := strconv.Atoi(song.ViewCount.Total)
			// if err != nil {
			// 	return nil, fmt.Errorf("error converting daily view count to int: %v", err)
			// }
			// // Update the views based on cronType
			// switch cronType {
			// case entities.Daily:
			// 	song.ViewCount.Daily = video.Statistics.ViewCount
			// case entities.Weekly:
			// 	song.ViewCount.Weekly = video.Statistics.ViewCount
			// case entities.Monthly:
			// 	song.ViewCount.Monthly = video.Statistics.ViewCount
			// }

			songs = append(songs, song)
		}
	}

	return songs, nil
}

// UpdateViewCounts updates the view counts of songs based on the YoutubeVideoListResponse data
func UpdateViewCounts(cronType entities.CronType, videoLists []entities.YTVideoListResponse, songs []*entities.Song) error {
	videoMap := make(map[string]entities.YTVideo)
	for _, videoList := range videoLists {
		for _, video := range videoList.Items {
			videoMap[video.ID] = video
		}
	}

	// updatedSongs := make([]entities.Song, len(songs))
	for _, song := range songs {
		video, exists := videoMap[song.ID]
		if !exists {
			return fmt.Errorf("no video data for song with ID %s", song.ID)
		}

		// Update the views
		switch cronType {
		case entities.Daily:
			song.ViewCount.Daily = video.Statistics.ViewCount
		case entities.Weekly:
			song.ViewCount.Weekly = video.Statistics.ViewCount
		case entities.Monthly:
			song.ViewCount.Monthly = video.Statistics.ViewCount
		}

		song.ViewCount.Total = video.Statistics.ViewCount

		// updatedSongs[i] = song
	}

	return nil
}

// GetSongIDs returns a slice of song IDs from a slice of songs.
func GetSongIDs(songs []*entities.Song) []string {
	ids := make([]string, len(songs))
	for i, song := range songs {
		ids[i] = song.ID
	}
	return ids
}
