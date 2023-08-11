package mappers

import (
	"time"

	"github.com/sugar-cat7/vspo-common-api/domain/entities"
)

type ChannelSnippetResponse struct {
	Title       string             `json:"title"`
	Description string             `json:"description"`
	CustomURL   string             `json:"customUrl"`
	PublishedAt time.Time          `json:"publishedAt"`
	Thumbnails  ThumbnailsResponse `json:"thumbnails"`
}

type ChannelStatisticsResponse struct {
	ViewCount             string `json:"viewCount"`
	SubscriberCount       string `json:"subscriberCount"`
	HiddenSubscriberCount bool   `json:"hiddenSubscriberCount"`
	VideoCount            string `json:"videoCount"`
}

type ChannelResponse struct {
	ID         string                    `json:"id"`
	Snippet    ChannelSnippetResponse    `json:"snippet"`
	Statistics ChannelStatisticsResponse `json:"statistics"`
}

type ChannelsResponse struct {
	Channels []ChannelResponse `json:"channels"`
}

func MapChannelSnippetToResponse(snippet entities.ChannelSnippet) ChannelSnippetResponse {
	return ChannelSnippetResponse{
		Title:       snippet.Title,
		Description: snippet.Description,
		CustomURL:   snippet.CustomURL,
		PublishedAt: snippet.PublishedAt,
		Thumbnails:  MapThumbnailsToResponse(snippet.Thumbnails),
	}
}

func MapChannelStatisticsToResponse(statistics entities.ChannelStatistics) ChannelStatisticsResponse {
	return ChannelStatisticsResponse{
		ViewCount:             statistics.ViewCount,
		SubscriberCount:       statistics.SubscriberCount,
		HiddenSubscriberCount: statistics.HiddenSubscriberCount,
		VideoCount:            statistics.VideoCount,
	}
}

func MapChannelToResponse(channel *entities.Channel) ChannelResponse {
	return ChannelResponse{
		ID:         channel.ID,
		Snippet:    MapChannelSnippetToResponse(channel.Snippet),
		Statistics: MapChannelStatisticsToResponse(channel.Statistics),
	}
}

func MapChannelsToResponse(channels []*entities.Channel) ChannelsResponse {
	var channelResponses []ChannelResponse
	for _, channel := range channels {
		channelResponses = append(channelResponses, MapChannelToResponse(channel))
	}
	return ChannelsResponse{Channels: channelResponses}
}
