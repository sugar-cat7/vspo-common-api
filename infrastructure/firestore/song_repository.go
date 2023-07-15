package firestore

import (
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
)

// SongRepository is a Firestore implementation of a song repository.
type SongRepository struct {
	repo *Repository
}

var collectionName = "songs"

// NewSongRepository creates a new SongRepository.
func NewSongRepository(client FirestoreClient) *SongRepository {
	return &SongRepository{
		repo: NewRepository(client, collectionName),
	}
}

// Create creates a new song in Firestore.
func (r *SongRepository) Create(song *entities.Song) error {
	return r.repo.Create(song)
}

// GetAll retrieves all songs from Firestore.
func (r *SongRepository) GetAll() ([]*entities.Song, error) {
	var songs []*entities.Song
	err := r.repo.GetAll(&songs)
	if err != nil {
		return nil, err
	}
	return songs, nil
}

// GetByID retrieves a song by its ID from Firestore.
func (r *SongRepository) GetByID(id string) (*entities.Song, error) {
	var song entities.Song
	err := r.repo.GetByID(id, &song)
	if err != nil {
		return nil, err
	}
	return &song, nil
}

// Update updates a song in Firestore.
func (r *SongRepository) Update(song *entities.Song) error {
	return r.repo.Update(song)
}

// Delete deletes a song from Firestore.
func (r *SongRepository) Delete(id string) error {
	return r.repo.Delete(id)
}

// UpdateInBatch updates multiple songs in Firestore.
func (r *SongRepository) UpdateInBatch(songs []*entities.Song) error {
	var docs []Doc
	for _, song := range songs {
		docs = append(docs, song)
	}
	return r.repo.UpdateInBatch(docs)
}

// CreateInBatch updates multiple songs in Firestore.
func (r *SongRepository) CreateInBatch(songs []*entities.Song) error {
	var docs []Doc
	for _, song := range songs {
		docs = append(docs, song)
	}
	return r.repo.UpdateInBatch(docs)
}
