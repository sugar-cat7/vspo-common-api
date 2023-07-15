//go:generate mockgen -destination=../../mocks/services/mock_song_service.go -package=mocks github.com/sugar-cat7/vspo-common-api/domain/services SongService
package services

import (
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/domain/repositories"
)

// SongService is an interface for a song service.
type SongService interface {
	CreateSong(song *entities.Song) error
	GetAllSongs() ([]*entities.Song, error)
	GetSongByID(id string) (*entities.Song, error)
	GetSongIDs() ([]string, error)
	UpdateSong(song *entities.Song) error
	DeleteSong(id string) error
	CreateSongsInBatch(songs []*entities.Song) error
	UpdateSongsInBatch(songs []*entities.Song) error
}
type songService struct {
	repo repositories.SongRepository
}

// NewSongService creates a new SongService.
func NewSongService(repo repositories.SongRepository) SongService {
	return &songService{repo: repo}
}

// CreateSong creates a new song.
func (s *songService) CreateSong(song *entities.Song) error {
	return s.repo.Create(song)
}

// GetAllSongs retrieves all songs.
func (s *songService) GetAllSongs() ([]*entities.Song, error) {
	return s.repo.GetAll()
}

// GetSongIDs retrieves all song IDs.
func (s *songService) GetSongIDs() ([]string, error) {
	songs, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	var songIDs []string
	for _, song := range songs {
		songIDs = append(songIDs, song.ID)
	}
	return songIDs, nil
}

// GetSongByID retrieves a song by its ID.
func (s *songService) GetSongByID(id string) (*entities.Song, error) {
	return s.repo.GetByID(id)
}

// UpdateSong updates a song.
func (s *songService) UpdateSong(song *entities.Song) error {
	return s.repo.Update(song)
}

// DeleteSong deletes a song.
func (s *songService) DeleteSong(id string) error {
	return s.repo.Delete(id)
}

// UpdateSongsInBatch updates multiple songs.
func (s *songService) UpdateSongsInBatch(songs []*entities.Song) error {
	return s.repo.UpdateInBatch(songs)
}

// UpdateSongsInBatch updates multiple songs.
func (s *songService) CreateSongsInBatch(songs []*entities.Song) error {
	return s.repo.CreateInBatch(songs)
}
