package mappers

import (
	"time"

	"github.com/sugar-cat7/vspo-common-api/domain/entities"
)

// ChannelSnippetResponse represents the snippet section of a YouTube channel response.
type ChannelSnippetResponse struct {
	Title       string             `json:"title" example:"花芽なずな / Nazuna Kaga"`
	Description string             `json:"description" example:"ぶいすぽ所属　最年少！５歳可愛い担当花芽なずなです♡\n\n好きなゲームはFPS全般！"`
	CustomURL   string             `json:"customUrl" example:"@nazunakaga"`
	PublishedAt time.Time          `json:"publishedAt" example:"2018-09-20T11:41:24Z"`
	Thumbnails  ThumbnailsResponse `json:"thumbnails"`
}

// ChannelStatisticsResponse represents the statistics section of a YouTube channel response.
type ChannelStatisticsResponse struct {
	ViewCount             string `json:"viewCount" example:"59373115"`
	SubscriberCount       string `json:"subscriberCount" example:"357000"`
	HiddenSubscriberCount bool   `json:"hiddenSubscriberCount" example:"false"`
	VideoCount            string `json:"videoCount" example:"1183"`
}

// ChannelResponse represents the main structure of a YouTube channel response.
type ChannelResponse struct {
	ID         string                    `json:"id" example:"UCiMG6VdScBabPhJ1ZtaVmbw"`
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
