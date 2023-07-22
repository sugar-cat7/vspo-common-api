package usecases

import (
	"github.com/sugar-cat7/vspo-common-api/domain/services"
	"github.com/sugar-cat7/vspo-common-api/util"
)

// CreateChannel is a use case for creating channels in Firestore from YouTube.
type CreateChannel struct {
	youtubeService services.YouTubeService
	channelService services.ChannelService
}

// NewCreateChannel creates a new CreateChannel.
func NewCreateChannel(youtubeService services.YouTubeService, channelService services.ChannelService) *CreateChannel {
	return &CreateChannel{
		youtubeService: youtubeService,
		channelService: channelService,
	}
}

// Execute creates new channels in Firestore from YouTube.
func (c *CreateChannel) Execute(ids []string) error {
	// Fetch channel data from YouTube API
	channels, err := c.youtubeService.GetChannels(ids)

	if err != nil {
		return err
	}

	// Save the new channels to Firestore
	err = c.channelService.CreateChannelsInBatch(util.ConvertToPtrSlice(channels))
	if err != nil {
		return err
	}

	return nil
}
