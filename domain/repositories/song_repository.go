//go:generate mockgen -destination=../../mocks/repositories/mock_song_repository.go -package=mocks github.com/sugar-cat7/vspo-common-api/domain/repositories SongRepository
package repositories

import (
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
)

// SongRepository is an interface for a song repository.
type SongRepository interface {
	Create(song *entities.Video) error
	GetAll() ([]*entities.Video, error)
	GetByID(id string) (*entities.Video, error)
	Update(song *entities.Video) error
	Delete(id string) error
	UpdateInBatch(songs []*entities.Video) error
	CreateInBatch(songs []*entities.Video) error
}
