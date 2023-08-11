package firestore

import (
	"context"
	"fmt"
	"time"

	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/domain/repositories"
	"github.com/sugar-cat7/vspo-common-api/util"
)

// SongRepository is a Firestore implementation of a song repository.
type SongRepository struct {
	client         repositories.FirestoreClient
	collectionName string
}

// NewSongRepository creates a new SongRepository.
func NewSongRepository(client repositories.FirestoreClient) *SongRepository {
	return &SongRepository{
		client:         client,
		collectionName: songsCollectionName,
	}
}

// Create creates a new song in Firestore.
func (r *SongRepository) Create(song *entities.Video) error {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.client.Create(r.collectionName, song.GetID(), song)

	if err != nil {
		return fmt.Errorf("failed to create document in Firestore: %w", err)
	}

	return nil
}

// GetAll retrieves all songs from Firestore.
func (r *SongRepository) GetAll() ([]*entities.Video, error) {
	docs, err := r.client.GetAll(r.collectionName)
	if err != nil {
		return nil, fmt.Errorf("failed to get all documents from Firestore: %w", err)
	}

	var songs []*entities.Video
	for _, doc := range docs {
		var song entities.Video
		err = doc.DataTo(&song)
		if err != nil {
			return nil, fmt.Errorf("failed to map document data to the provided struct: %w", err)
		}
		songs = append(songs, &song)
	}

	return songs, nil
}

// GetByID retrieves a song by its ID from Firestore.
func (r *SongRepository) GetByID(id string) (*entities.Video, error) {
	doc, err := r.client.GetByID(r.collectionName, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get document by ID from Firestore: %w", err)
	}

	var song entities.Video
	err = doc.DataTo(&song)
	if err != nil {
		return nil, fmt.Errorf("failed to map document data to the provided struct: %w", err)
	}

	return &song, nil
}

// Update updates a song in Firestore.
func (r *SongRepository) Update(song *entities.Video) error {
	_, err := r.client.Update(r.collectionName, song.GetID(), song.GetUpdate())

	if err != nil {
		return fmt.Errorf("failed to update document in Firestore: %w", err)
	}

	return nil
}

// Delete deletes a song from Firestore.
func (r *SongRepository) Delete(id string) error {
	_, err := r.client.Delete(r.collectionName, id)

	if err != nil {
		return fmt.Errorf("failed to delete document from Firestore: %w", err)
	}

	return nil
}

// UpdateInBatch updates multiple songs in Firestore using batch operation.
func (r *SongRepository) UpdateInBatch(songs []*entities.Video) error {
	// split songs into chunks of 500 (maxBatchSize)
	songChunks, err := util.Chunk(songs, maxBatchSize)
	if err != nil {
		return fmt.Errorf("failed to chunk songs: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	bulkWriter := r.client.BulkWriter(ctx)

	for _, songChunk := range songChunks {
		// Iterate over songs and add them to the bulkWriter.
		for _, song := range songChunk {
			docRef := r.client.Collection(r.collectionName).Doc(song.GetID())
			err := bulkWriter.Update(docRef, song.GetUpdate())
			if err != nil {
				return fmt.Errorf("failed to add update to bulkWriter in Firestore: %w", err)
			}
		}
	}

	// Commit the bulk writes.
	bulkWriter.Flush()

	return nil
}

// CreateInBatch creates multiple songs in Firestore using batch operation.
func (r *SongRepository) CreateInBatch(songs []*entities.Video) error {
	// split songs into chunks of 500 (maxBatchSize)
	songChunks, err := util.Chunk(songs, maxBatchSize)
	if err != nil {
		return fmt.Errorf("failed to chunk songs: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	bulkWriter := r.client.BulkWriter(ctx)

	for _, songChunk := range songChunks {
		// Iterate over songs and add them to the bulkWriter.
		for _, song := range songChunk {
			docRef := r.client.Collection(r.collectionName).Doc(song.GetID())
			err := bulkWriter.Create(docRef, song)
			if err != nil {
				return fmt.Errorf("failed to add creation to bulkWriter in Firestore: %w", err)
			}
		}
	}

	// Commit the bulk writes.
	bulkWriter.Flush()

	return nil
}
