package usecases

import (
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/domain/services"
)

// GetChannels is a use case for getting all channels from Firestore.
type GetChannels struct {
	channelService services.ChannelService
}

// NewGetChannels creates a new GetChannels.
func NewGetChannels(channelService services.ChannelService) *GetChannels {
	return &GetChannels{
		channelService: channelService,
	}
}

// Execute gets all channels from Firestore.
func (g *GetChannels) Execute(ids []string) ([]*entities.Channel, error) {
	// Get all channels from Firestore
	channels, err := g.channelService.GetChannels(ids)
	if err != nil {
		return nil, err
	}

	return channels, nil
}
