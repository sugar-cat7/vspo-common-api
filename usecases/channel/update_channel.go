package usecases

import (
	"github.com/sugar-cat7/vspo-common-api/domain/services"
	"github.com/sugar-cat7/vspo-common-api/usecases/mappers"
)

// UpdateChannelsFromYoutube is a use case for updating channels in Firestore from YouTube.
type UpdateChannelsFromYoutube struct {
	youtubeService services.YouTubeService
	channelService services.ChannelService
	channelMapper  *mappers.ChannelMapper
}

// NewUpdateChannelsFromYoutube creates a new UpdateChannelsFromYoutube.
func NewUpdateChannelsFromYoutube(youtubeService services.YouTubeService, channelService services.ChannelService, channelMapper *mappers.ChannelMapper) *UpdateChannelsFromYoutube {
	return &UpdateChannelsFromYoutube{
		youtubeService: youtubeService,
		channelService: channelService,
		channelMapper:  channelMapper,
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
	channels, err := u.channelMapper.MapMultiple(ytChannels)
	if err != nil {
		return err
	}

	// Update the channels in Firestore
	err = u.channelService.UpdateChannelsInBatch(channels)
	if err != nil {
		return err
	}

	return nil
}
