package usecases

import (
	"github.com/sugar-cat7/vspo-common-api/domain/ports"
	"github.com/sugar-cat7/vspo-common-api/domain/repositories"
	"github.com/sugar-cat7/vspo-common-api/usecases/mappers"
)

// UpdateChannelsFromYoutube is a use case for updating channels in Firestore from YouTube.
type UpdateChannelsFromYoutube struct {
	youtubeService    ports.YouTubeService
	channelRepository repositories.ChannelRepository
}

// NewUpdateChannelsFromYoutube creates a new UpdateChannelsFromYoutube.
func NewUpdateChannelsFromYoutube(youtubeService ports.YouTubeService, channelRepository repositories.ChannelRepository) *UpdateChannelsFromYoutube {
	return &UpdateChannelsFromYoutube{
		youtubeService:    youtubeService,
		channelRepository: channelRepository,
	}
}

// Execute updates the channels in Firestore from YouTube.
func (u *UpdateChannelsFromYoutube) Execute(ids []string) error {
	// Fetch channel data from YouTube API
	ytChannels, err := u.youtubeService.GetChannels(ids)

	if err != nil {
		return err
	}

	// Map ytChannels to domain entities
	channels, err := mappers.ChannelMapMultiple(ytChannels)
	if err != nil {
		return err
	}

	// Update the channels in Firestore
	err = u.channelRepository.UpdateInBatch(channels)
	if err != nil {
		return err
	}

	return nil
}
