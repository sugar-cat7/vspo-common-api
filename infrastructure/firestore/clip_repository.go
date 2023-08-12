package firestore

import (
	"context"
	"fmt"
	"time"

	entities "github.com/sugar-cat7/vspo-common-api/domain/entities/legacy"
	"github.com/sugar-cat7/vspo-common-api/domain/repositories"
	"github.com/sugar-cat7/vspo-common-api/util"
)

type ClipRepository struct {
	client         repositories.FirestoreClient
	collectionName string
}

func NewClipRepository(client repositories.FirestoreClient) *ClipRepository {
	return &ClipRepository{
		client:         client,
		collectionName: clipsCollectionName,
	}
}

func (r *ClipRepository) FindAllByPeriod(start, end string) ([]*entities.Clip, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	startTime, err := time.Parse(time.RFC3339, start)
	if err != nil {
		return nil, fmt.Errorf("failed to parse start time: %w", err)
	}

	// endTime, err := time.Parse(time.RFC3339, end)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to parse end time: %w", err)
	// }

	docs, err := r.client.Collection(r.collectionName).Where("CreatedAt", ">=", startTime).Documents(ctx).GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get documents from Firestore: %w", err)
	}

	var clips []*entities.Clip
	for _, doc := range docs {
		var clip entities.Clip
		err = doc.DataTo(&clip)
		if err != nil {
			return nil, fmt.Errorf("failed to map document data to the provided struct: %w", err)
		}
		clips = append(clips, &clip)
	}

	return clips, nil
}

func (r *ClipRepository) UpdateInBatch(clips []*entities.Clip) error {
	clipChunks, err := util.Chunk(clips, maxBatchSize)
	if err != nil {
		return fmt.Errorf("failed to chunk clips: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	bulkWriter := r.client.BulkWriter(ctx)

	for _, clipChunk := range clipChunks {
		for _, clip := range clipChunk {
			docRef := r.client.Collection(r.collectionName).Doc(clip.ID)
			err := bulkWriter.Update(docRef, clip.GetUpdate())
			if err != nil {
				return fmt.Errorf("failed to add update to bulkWriter in Firestore: %w", err)
			}
		}
	}

	bulkWriter.Flush()
	return nil
}

func (r *ClipRepository) CreateInBatch(clips []*entities.Clip) error {
	clipChunks, err := util.Chunk(clips, maxBatchSize)
	if err != nil {
		return fmt.Errorf("failed to chunk clips: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	bulkWriter := r.client.BulkWriter(ctx)

	for _, clipChunk := range clipChunks {
		for _, clip := range clipChunk {
			docRef := r.client.Collection(r.collectionName).Doc(clip.ID)
			err := bulkWriter.Create(docRef, clip)
			if err != nil {
				return fmt.Errorf("failed to add creation to bulkWriter in Firestore: %w", err)
			}
		}
	}

	bulkWriter.Flush()
	return nil
}
