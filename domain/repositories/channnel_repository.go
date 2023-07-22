package repositories

import (
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
)

// ChannelRepository is an interface for a channel repository.
type ChannelRepository interface {
	Create(channel *entities.Channel) error
	GetAll() ([]*entities.Channel, error)
	GetByID(id string) (*entities.Channel, error)
	Update(channel *entities.Channel) error
	Delete(id string) error
	GetInBatch(ids []string) ([]*entities.Channel, error)
	UpdateInBatch(channels []*entities.Channel) error
	CreateInBatch(channels []*entities.Channel) error
}
