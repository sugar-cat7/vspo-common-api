package firestore

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/util"
	"google.golang.org/api/iterator"
)

// SongRepository is a Firestore implementation of a song repository.
type SongRepository struct {
	client         FirestoreClient
	collectionName string
}

var collectionName = "songs"

// NewSongRepository creates a new SongRepository.
func NewSongRepository(client FirestoreClient) *SongRepository {
	return &SongRepository{
		client:         client,
		collectionName: collectionName,
	}
}

// Create creates a new song in Firestore.
func (r *SongRepository) Create(song *entities.Song) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.client.Collection(r.collectionName).Doc(song.GetID()).Set(ctx, song)

	if err != nil {
		return fmt.Errorf("failed to create document in Firestore: %w", err)
	}

	return nil
}

// GetAll retrieves all songs from Firestore.
func (r *SongRepository) GetAll() ([]*entities.Song, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	iter, err := r.client.Collection(r.collectionName).Documents(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all documents from Firestore: %w", err)
	}

	var songs []*entities.Song
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to iterate over documents in Firestore: %w", err)
		}

		var song entities.Song
		err = doc.DataTo(&song)
		if err != nil {
			return nil, fmt.Errorf("failed to map document data to the provided struct: %w", err)
		}
		songs = append(songs, &song)
	}

	return songs, nil
}

// GetByID retrieves a song by its ID from Firestore.
func (r *SongRepository) GetByID(id string) (*entities.Song, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	doc, err := r.client.Collection(r.collectionName).Doc(id).Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get document by ID from Firestore: %w", err)
	}

	var song entities.Song
	err = doc.DataTo(&song)
	if err != nil {
		return nil, fmt.Errorf("failed to map document data to the provided struct: %w", err)
	}

	return &song, nil
}

// Update updates a song in Firestore.
func (r *SongRepository) Update(song *entities.Song) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.client.Collection(r.collectionName).Doc(song.GetID()).Update(ctx, song.GetUpdate())

	if err != nil {
		return fmt.Errorf("failed to update document in Firestore: %w", err)
	}

	return nil
}

// Delete deletes a song from Firestore.
func (r *SongRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.client.Collection(r.collectionName).Doc(id).Delete(ctx)

	if err != nil {
		return fmt.Errorf("failed to delete document from Firestore: %w", err)
	}

	return nil
}

var maxBatchSize = 500

// UpdateInBatch updates multiple songs in Firestore using batch operation.
func (r *SongRepository) UpdateInBatch(songs []*entities.Song) error {
	// split songs into chunks of 500 (maxBatchSize)
	songChunks, err := util.Chunk(songs, maxBatchSize)
	if err != nil {
		return fmt.Errorf("failed to chunk songs: %w", err)
	}

	for _, songChunk := range songChunks {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Start a new batch.
		batch := r.client.Batch()

		// Iterate over songs and add them to the batch.
		for _, song := range songChunk {
			docRef := r.client.Collection(r.collectionName).Doc(song.GetID())
			batch.Update(docRef, song.GetUpdate())
		}

		// Commit the batch.
		_, err := batch.Commit(ctx)
		if err != nil {
			return fmt.Errorf("failed to update songs in Firestore: %w", err)
		}
	}

	return nil
}

// CreateInBatch creates multiple songs in Firestore.
func (r *SongRepository) CreateInBatch(songs []*entities.Song) error {
	// split songs into chunks of 500 (maxBatchSize)
	songChunks, err := util.Chunk(songs, maxBatchSize)
	if err != nil {
		return fmt.Errorf("failed to chunk songs: %w", err)
	}

	for _, songChunk := range songChunks {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Start a new batch.
		batch := r.client.Batch()

		// Iterate over songs and add them to the batch.
		for _, song := range songChunk {
			docRef := r.client.Collection(r.collectionName).Doc(song.GetID())
			batch.Set(docRef, song, firestore.MergeAll)
		}

		// Commit the batch.
		_, err := batch.Commit(ctx)
		if err != nil {
			return fmt.Errorf("failed to create songs in Firestore: %w", err)
		}
	}

	return nil
}
