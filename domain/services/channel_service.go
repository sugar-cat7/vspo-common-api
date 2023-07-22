//go:generate mockgen -destination=../../mocks/services/mock_channel_service.go -package=mocks github.com/sugar-cat7/vspo-common-api/domain/services ChannelService
package services

import (
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/domain/repositories"
)

// ChannelService is an interface for a channel service.
type ChannelService interface {
	CreateChannel(channel *entities.Channel) error
	GetAllChannels() ([]*entities.Channel, error)
	GetChannels(ids []string) ([]*entities.Channel, error)
	GetChannelByID(id string) (*entities.Channel, error)
	GetChannelIDs() ([]string, error)
	UpdateChannel(channel *entities.Channel) error
	DeleteChannel(id string) error
	CreateChannelsInBatch(channels []*entities.Channel) error
	UpdateChannelsInBatch(channels []*entities.Channel) error
}
type channelService struct {
	repo repositories.ChannelRepository
}

// NewChannelService creates a new ChannelService.
func NewChannelService(repo repositories.ChannelRepository) ChannelService {
	return &channelService{repo: repo}
}

// CreateChannel creates a new channel.
func (s *channelService) CreateChannel(channel *entities.Channel) error {
	return s.repo.Create(channel)
}

// GetAllChannels retrieves all channels.
func (s *channelService) GetAllChannels() ([]*entities.Channel, error) {
	return s.repo.GetAll()
}

// GetChannelIDs retrieves all channel IDs.
func (s *channelService) GetChannelIDs() ([]string, error) {
	channels, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	var channelIDs []string
	for _, channel := range channels {
		channelIDs = append(channelIDs, channel.ID)
	}
	return channelIDs, nil
}

// GetChannelByID retrieves a channel by its ID.
func (s *channelService) GetChannelByID(id string) (*entities.Channel, error) {
	return s.repo.GetByID(id)
}

// UpdateChannel updates a channel.
func (s *channelService) UpdateChannel(channel *entities.Channel) error {
	return s.repo.Update(channel)
}

// DeleteChannel deletes a channel.
func (s *channelService) DeleteChannel(id string) error {
	return s.repo.Delete(id)
}

// UpdateChannelsInBatch updates multiple channels.
func (s *channelService) UpdateChannelsInBatch(channels []*entities.Channel) error {
	return s.repo.UpdateInBatch(channels)
}

// UpdateChannelsInBatch updates multiple channels.
func (s *channelService) CreateChannelsInBatch(channels []*entities.Channel) error {
	return s.repo.CreateInBatch(channels)
}

// GetAllChannels retrieves all channels.
func (s *channelService) GetChannels(ids []string) ([]*entities.Channel, error) {
	return s.repo.GetInBatch(ids)
}
