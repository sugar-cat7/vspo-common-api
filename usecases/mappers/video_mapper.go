package mappers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"google.golang.org/api/youtube/v3"
)

// MapToVideos takes a list of YouTube API video objects and maps them to custom video objects.
func MapToVideos(cronType entities.CronType, ytVideos []*youtube.Video) entities.Videos {
	videos := make(entities.Videos, len(ytVideos))
	for i, ytVideo := range ytVideos {
		videos[i] = mapToVideo(cronType, ytVideo)
	}
	return videos
}

// mapToVideo takes a YouTube API video object and maps it to a custom video object.
func mapToVideo(cronType entities.CronType, ytVideo *youtube.Video) *entities.Video {
	publishedAt, _ := time.Parse(time.RFC3339, ytVideo.Snippet.PublishedAt) // Error handling can be added here
	video := &entities.Video{
		ID:          ytVideo.Id,
		Title:       ytVideo.Snippet.Title,
		Description: ytVideo.Snippet.Description,
		ViewCount: entities.Views{
			Total: fmt.Sprintf("%d", ytVideo.Statistics.ViewCount),
		},
		PublishedAt: publishedAt,
		Thumbnails: entities.Thumbnails{
			Default: entities.Thumbnail{URL: ytVideo.Snippet.Thumbnails.Default.Url, Width: int(ytVideo.Snippet.Thumbnails.Default.Width), Height: int(ytVideo.Snippet.Thumbnails.Default.Height)},
			Medium:  entities.Thumbnail{URL: ytVideo.Snippet.Thumbnails.Medium.Url, Width: int(ytVideo.Snippet.Thumbnails.Medium.Width), Height: int(ytVideo.Snippet.Thumbnails.Medium.Height)},
			High:    entities.Thumbnail{URL: ytVideo.Snippet.Thumbnails.High.Url, Width: int(ytVideo.Snippet.Thumbnails.High.Width), Height: int(ytVideo.Snippet.Thumbnails.High.Height)},
			// Standard: entities.Thumbnail{URL: ytVideo.Snippet.Thumbnails.Standard.Url, Width: int(ytVideo.Snippet.Thumbnails.Standard.Width), Height: int(ytVideo.Snippet.Thumbnails.Standard.Height)},
			// Maxres:   entities.Thumbnail{URL: ytVideo.Snippet.Thumbnails.Maxres.Url, Width: int(ytVideo.Snippet.Thumbnails.Maxres.Width), Height: int(ytVideo.Snippet.Thumbnails.Maxres.Height)},
		},
		ChannelTitle: ytVideo.Snippet.ChannelTitle,
		ChannelID:    ytVideo.Snippet.ChannelId,
		// You may need additional logic to map ChannelIcon and Platform
		Tags: ytVideo.Snippet.Tags,
	}

	switch cronType {
	case entities.Daily:
		video.ViewCount.Daily = strconv.FormatUint(ytVideo.Statistics.ViewCount, 10)
	case entities.Weekly:
		video.ViewCount.Weekly = strconv.FormatUint(ytVideo.Statistics.ViewCount, 10)
	case entities.Monthly:
		video.ViewCount.Monthly = strconv.FormatUint(ytVideo.Statistics.ViewCount, 10)
	}

	return video
}
