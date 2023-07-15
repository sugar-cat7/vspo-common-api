package usecases

import (
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/domain/services"
)

// GetAllSongs is a use case for getting all songs from Firestore.
type GetAllSongs struct {
	songService services.SongService
}

// NewGetAllSongs creates a new GetAllSongs.
func NewGetAllSongs(songService services.SongService) *GetAllSongs {
	return &GetAllSongs{
		songService: songService,
	}
}

// Execute gets all songs from Firestore.
func (g *GetAllSongs) Execute() ([]*entities.Song, error) {
	// Get all songs from Firestore
	songs, err := g.songService.GetAllSongs()
	if err != nil {
		return nil, err
	}

	return songs, nil
}
