package util

import (
	"strconv"

	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"google.golang.org/api/youtube/v3"
)

// UpdateViewCounts updates the view counts of videos based on the YoutubeVideoListResponse data
func UpdateViewCounts(cronType entities.CronType, ytVideos []*youtube.Video, videos []*entities.Video) ([]*entities.Video, error) {
	videoMap := make(map[string]*youtube.Video)
	for _, ytVideo := range ytVideos {
		videoMap[ytVideo.Id] = ytVideo
	}

	// Initialize updatedVideos to store updated videos
	updatedVideos := make([]*entities.Video, 0)

	for _, video := range videos {
		ytVideo, exists := videoMap[video.ID]
		if exists {
			// Update the views
			switch cronType {
			case entities.Daily:
				video.ViewCount.Daily = strconv.FormatUint(ytVideo.Statistics.ViewCount, 10)
			case entities.Weekly:
				video.ViewCount.Weekly = strconv.FormatUint(ytVideo.Statistics.ViewCount, 10)
			case entities.Monthly:
				video.ViewCount.Monthly = strconv.FormatUint(ytVideo.Statistics.ViewCount, 10)
			}

			video.ViewCount.Total = strconv.FormatUint(ytVideo.Statistics.ViewCount, 10)
			updatedVideos = append(updatedVideos, video)
		}
	}

	return updatedVideos, nil
}

// GetVideoIDs returns a slice of video IDs from a slice of videos.
func GetVideoIDs(videos []*entities.Video) []string {
	ids := make([]string, len(videos))
	for i, video := range videos {
		ids[i] = video.ID
	}
	return ids
}
