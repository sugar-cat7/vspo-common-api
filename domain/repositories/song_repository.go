package repositories

import (
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
)

// SongRepository is an interface for a song repository.
type SongRepository interface {
	Create(song *entities.Song) error
	GetAll() ([]*entities.Song, error)
	GetByID(id string) (*entities.Song, error)
	Update(song *entities.Song) error
	Delete(id string) error
	UpdateInBatch(songs []*entities.Song) error
	CreateInBatch(songs []*entities.Song) error
}
