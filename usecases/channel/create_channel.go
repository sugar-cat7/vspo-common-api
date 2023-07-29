package usecases

import (
	"github.com/sugar-cat7/vspo-common-api/domain/services"
	"github.com/sugar-cat7/vspo-common-api/usecases/mappers"
)

// CreateChannel is a use case for creating channels in Firestore from YouTube.
type CreateChannel struct {
	youtubeService services.YouTubeService
	channelService services.ChannelService
	channelMapper  *mappers.ChannelMapper
}

// NewCreateChannel creates a new CreateChannel.
func NewCreateChannel(youtubeService services.YouTubeService, channelService services.ChannelService, channelMapper *mappers.ChannelMapper) *CreateChannel {
	return &CreateChannel{
		youtubeService: youtubeService,
		channelService: channelService,
		channelMapper:  channelMapper,
	}
}

// Execute creates new channels in Firestore from YouTube.
func (c *CreateChannel) Execute(ids []string) error {
	// Fetch channel data from YouTube API
	ytChannels, err := c.youtubeService.GetChannels(ids)
	if err != nil {
		return err
	}

	// Map ytChannels to domain entities
	channels, err := c.channelMapper.MapMultiple(ytChannels)
	if err != nil {
		return err
	}

	// Save the new channels to Firestore
	err = c.channelService.CreateChannelsInBatch(channels)
	if err != nil {
		return err
	}

	return nil
}
