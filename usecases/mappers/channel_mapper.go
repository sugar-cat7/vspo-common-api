package mappers

import (
	"time"

	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"google.golang.org/api/youtube/v3"
)

// ChannelMap maps a YouTube API channel to a domain Channel.
func ChannelMap(ytChannel *youtube.Channel) (*entities.Channel, error) {
	pubAt, err := time.Parse(time.RFC3339, ytChannel.Snippet.PublishedAt)
	if err != nil {
		return nil, err
	}

	return &entities.Channel{
		ID: ytChannel.Id,
		Snippet: entities.ChannelSnippet{
			Title:       ytChannel.Snippet.Title,
			Description: ytChannel.Snippet.Description,
			CustomURL:   ytChannel.Snippet.CustomUrl,
			PublishedAt: pubAt,
			Thumbnails: entities.Thumbnails{
				Default: entities.Thumbnail{
					URL:    ytChannel.Snippet.Thumbnails.Default.Url,
					Width:  int(ytChannel.Snippet.Thumbnails.Default.Width),
					Height: int(ytChannel.Snippet.Thumbnails.Default.Height),
				},
				Medium: entities.Thumbnail{
					URL:    ytChannel.Snippet.Thumbnails.Medium.Url,
					Width:  int(ytChannel.Snippet.Thumbnails.Medium.Width),
					Height: int(ytChannel.Snippet.Thumbnails.Medium.Height),
				},
				//FIXME:Nullの場合がある
				// High: entities.Thumbnail{
				// 	URL:    ytChannel.Snippet.Thumbnails.High.Url,
				// 	Width:  int(ytChannel.Snippet.Thumbnails.High.Width),
				// 	Height: int(ytChannel.Snippet.Thumbnails.High.Height),
				// },
				// Standard: entities.Thumbnail{
				// 	URL:    ytChannel.Snippet.Thumbnails.Standard.Url,
				// 	Width:  int(ytChannel.Snippet.Thumbnails.Standard.Width),
				// 	Height: int(ytChannel.Snippet.Thumbnails.Standard.Height),
				// },
				// Maxres: entities.Thumbnail{
				// 	URL:    ytChannel.Snippet.Thumbnails.Maxres.Url,
				// 	Width:  int(ytChannel.Snippet.Thumbnails.Maxres.Width),
				// 	Height: int(ytChannel.Snippet.Thumbnails.Maxres.Height),
				// },
			},
		},
		// Statistics: entities.ChannelStatistics{
		// 	ViewCount:             strconv.FormatUint(ytChannel.Statistics.ViewCount, 10),
		// 	SubscriberCount:       strconv.FormatUint(ytChannel.Statistics.SubscriberCount, 10),
		// HiddenSubscriberCount: ytChannel.Statistics.HiddenSubscriberCount,
		// 	VideoCount:            strconv.FormatUint(ytChannel.Statistics.VideoCount, 10),
		// },
	}, nil
}

// ChannelMapMultiple maps multiple YouTube API channels to domain Channels.
func ChannelMapMultiple(ytChannels []*youtube.Channel) ([]*entities.Channel, error) {
	channels := make([]*entities.Channel, len(ytChannels))
	for i, ytChannel := range ytChannels {
		channel, err := ChannelMap(ytChannel)
		if err != nil {
			return nil, err
		}
		channels[i] = channel
	}
	return channels, nil
}
