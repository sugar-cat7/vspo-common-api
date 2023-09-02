package usecases

import (
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/domain/repositories"
)

// GetChannels is a use case for getting all channels from Firestore.
type GetChannels struct {
	channelRepository repositories.ChannelRepository
}

// NewGetChannels creates a new GetChannels.
func NewGetChannels(channelRepository repositories.ChannelRepository) *GetChannels {
	return &GetChannels{
		channelRepository: channelRepository,
	}
}

// Execute gets all channels from Firestore.
func (g *GetChannels) Execute(ids []string) ([]*entities.Channel, error) {
	// Get all channels from Firestore
	channels, err := g.channelRepository.GetInBatch(ids)
	if err != nil {
		return nil, err
	}

	return channels, nil
}
