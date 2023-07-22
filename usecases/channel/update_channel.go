package usecases

import (
	"github.com/sugar-cat7/vspo-common-api/domain/services"
	"github.com/sugar-cat7/vspo-common-api/util"
)

// UpdateChannelsFromYoutube is a use case for updating channels in Firestore from YouTube.
type UpdateChannelsFromYoutube struct {
	youtubeService services.YouTubeService
	channelService services.ChannelService
}

// NewUpdateChannelsFromYoutube creates a new UpdateChannelsFromYoutube.
func NewUpdateChannelsFromYoutube(youtubeService services.YouTubeService, channelService services.ChannelService) *UpdateChannelsFromYoutube {
	return &UpdateChannelsFromYoutube{
		youtubeService: youtubeService,
		channelService: channelService,
	}
}

// Execute updates the channels in Firestore from YouTube.
func (u *UpdateChannelsFromYoutube) Execute(ids []string) error {
	// Fetch channel data from YouTube API
	channels, err := u.youtubeService.GetChannels(ids)

	if err != nil {
		return err
	}

	// Update the channels in Firestore
	err = u.channelService.UpdateChannelsInBatch(util.ConvertToPtrSlice(channels))
	if err != nil {
		return err
	}

	return nil
}
