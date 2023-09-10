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

func (r *ClipRepository) FindAllByPeriod(start, end string) ([]*entities.OldVideo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	const layout = "2006-01-02"
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, fmt.Errorf("failed to load JST location: %w", err)
	}

	if start == "" {
		startTimeJST := time.Now().In(jst).AddDate(0, 0, -7) // 1週間前の日付
		start = startTimeJST.Format(layout)
	}
	startTimeJST, err := time.ParseInLocation(layout, start, jst)
	if err != nil {
		return nil, fmt.Errorf("failed to parse start time: %w", err)
	}
	startTimeUTC := startTimeJST.UTC()

	if end == "" {
		endTimeJST := time.Now().In(jst).AddDate(0, 0, 7) // 7日後の日付
		end = endTimeJST.Format(layout)
	}
	endTimeJST, err := time.ParseInLocation(layout, end, jst)
	if err != nil {
		return nil, fmt.Errorf("failed to parse end time: %w", err)
	}
	endTimeUTC := endTimeJST.UTC()

	docs, err := r.client.Collection(r.collectionName).Where("createdAt", ">=", startTimeUTC).Where("createdAt", "<", endTimeUTC).Documents(ctx).GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get documents from Firestore: %w", err)
	}

	var clips []*entities.OldVideo
	for _, doc := range docs {
		var clip entities.OldVideo
		err = doc.DataTo(&clip)
		if err != nil {
			return nil, fmt.Errorf("failed to map document data to the provided struct: %w", err)
		}
		clips = append(clips, &clip)
	}

	return clips, nil
}

func (r *ClipRepository) UpdateInBatch(clips []*entities.OldVideo) error {
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

func (r *ClipRepository) CreateInBatch(clips []*entities.OldVideo) error {
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
