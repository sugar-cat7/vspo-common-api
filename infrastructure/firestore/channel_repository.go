package firestore

import (
	"context"
	"fmt"
	"time"

	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/domain/repositories"
	"github.com/sugar-cat7/vspo-common-api/util"
)

// ChannelRepository is a Firestore implementation of a channel repository.
type ChannelRepository struct {
	client         repositories.FirestoreClient
	collectionName string
}

// NewChannelRepository creates a new ChannelRepository.
func NewChannelRepository(client repositories.FirestoreClient) *ChannelRepository {
	return &ChannelRepository{
		client:         client,
		collectionName: channelsCollectionName,
	}
}

// Create creates a new channel in Firestore.
func (r *ChannelRepository) Create(channel *entities.Channel) error {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.client.Create(r.collectionName, channel.GetID(), channel)

	if err != nil {
		return fmt.Errorf("failed to create document in Firestore: %w", err)
	}

	return nil
}

// GetAll retrieves all channels from Firestore.
func (r *ChannelRepository) GetAll() ([]*entities.Channel, error) {
	docs, err := r.client.GetAll(r.collectionName)
	if err != nil {
		return nil, fmt.Errorf("failed to get all documents from Firestore: %w", err)
	}

	var channels []*entities.Channel
	for _, doc := range docs {
		var channel entities.Channel
		err = doc.DataTo(&channel)
		if err != nil {
			return nil, fmt.Errorf("failed to map document data to the provided struct: %w", err)
		}
		channels = append(channels, &channel)
	}

	return channels, nil
}

// GetByID retrieves a channel by its ID from Firestore.
func (r *ChannelRepository) GetByID(id string) (*entities.Channel, error) {
	doc, err := r.client.GetByID(r.collectionName, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get document by ID from Firestore: %w", err)
	}

	var channel entities.Channel
	err = doc.DataTo(&channel)
	if err != nil {
		return nil, fmt.Errorf("failed to map document data to the provided struct: %w", err)
	}

	return &channel, nil
}

// Update updates a channel in Firestore.
func (r *ChannelRepository) Update(channel *entities.Channel) error {
	_, err := r.client.Update(r.collectionName, channel.GetID(), channel.GetUpdate())

	if err != nil {
		return fmt.Errorf("failed to update document in Firestore: %w", err)
	}

	return nil
}

// Delete deletes a channel from Firestore.
func (r *ChannelRepository) Delete(id string) error {
	_, err := r.client.Delete(r.collectionName, id)

	if err != nil {
		return fmt.Errorf("failed to delete document from Firestore: %w", err)
	}

	return nil
}

// UpdateInBatch updates multiple channels in Firestore using batch operation.
func (r *ChannelRepository) UpdateInBatch(channels []*entities.Channel) error {
	// split channels into chunks of 500 (maxBatchSize)
	channelChunks, err := util.Chunk(channels, maxBatchSize)
	if err != nil {
		return fmt.Errorf("failed to chunk channels: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	bulkWriter := r.client.BulkWriter(ctx)

	for _, channelChunk := range channelChunks {
		// Iterate over channels and add them to the bulkWriter.
		for _, channel := range channelChunk {
			docRef := r.client.Collection(r.collectionName).Doc(channel.GetID())
			err := bulkWriter.Update(docRef, channel.GetUpdate())
			if err != nil {
				return fmt.Errorf("failed to add update to bulkWriter in Firestore: %w", err)
			}
		}
	}

	// Commit the bulk writes.
	bulkWriter.Flush()

	return nil
}

// CreateInBatch creates multiple channels in Firestore using batch operation.
func (r *ChannelRepository) CreateInBatch(channels []*entities.Channel) error {
	// split channels into chunks of 500 (maxBatchSize)
	channelChunks, err := util.Chunk(channels, maxBatchSize)
	if err != nil {
		return fmt.Errorf("failed to chunk channels: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	bulkWriter := r.client.BulkWriter(ctx)

	for _, channelChunk := range channelChunks {
		// Iterate over channels and add them to the bulkWriter.
		for _, channel := range channelChunk {
			docRef := r.client.Collection(r.collectionName).Doc(channel.GetID())
			err := bulkWriter.Create(docRef, channel)
			if err != nil {
				return fmt.Errorf("failed to add creation to bulkWriter in Firestore: %w", err)
			}
		}
	}

	// Commit the bulk writes.
	bulkWriter.Flush()

	return nil
}

// GetInBatch retrieves channels from Firestore.
func (r *ChannelRepository) GetInBatch(ids []string) ([]*entities.Channel, error) {
	docs, err := r.client.GetInBatch(r.collectionName, ids)
	if err != nil {
		return nil, fmt.Errorf("failed to get all documents from Firestore: %w", err)
	}

	var channels []*entities.Channel
	for _, doc := range docs {
		var channel entities.Channel
		err = doc.DataTo(&channel)
		if err != nil {
			return nil, fmt.Errorf("failed to map document data to the provided struct: %w", err)
		}
		channels = append(channels, &channel)
	}

	return channels, nil
}
