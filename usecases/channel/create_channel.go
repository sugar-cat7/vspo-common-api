package usecases

import (
	"github.com/sugar-cat7/vspo-common-api/domain/ports"
	"github.com/sugar-cat7/vspo-common-api/domain/repositories"
	"github.com/sugar-cat7/vspo-common-api/usecases/mappers"
)

// CreateChannel is a use case for creating channels in Firestore from YouTube.
type CreateChannel struct {
	youtubeService    ports.YouTubeService
	channelRepository repositories.ChannelRepository
}

// NewCreateChannel creates a new CreateChannel.
func NewCreateChannel(youtubeService ports.YouTubeService, channelRepository repositories.ChannelRepository) *CreateChannel {
	return &CreateChannel{
		youtubeService:    youtubeService,
		channelRepository: channelRepository,
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
	channels, err := mappers.ChannelMapMultiple(ytChannels)
	if err != nil {
		return err
	}

	// Save the new channels to Firestore
	err = c.channelRepository.CreateInBatch(channels)
	if err != nil {
		return err
	}

	return nil
}
