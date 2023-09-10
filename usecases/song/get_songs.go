package usecases

import (
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/domain/repositories"
)

// GetAllSongs is a use case for getting all songs from Firestore.
type GetAllSongs struct {
	songRepository repositories.SongRepository
}

// NewGetAllSongs creates a new GetAllSongs.
func NewGetAllSongs(songRepository repositories.SongRepository) *GetAllSongs {
	return &GetAllSongs{
		songRepository: songRepository,
	}
}

// Execute gets all songs from Firestore.
func (g *GetAllSongs) Execute() (entities.Videos, error) {
	// Get all songs from Firestore
	songs, err := g.songRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return songs, nil
}
