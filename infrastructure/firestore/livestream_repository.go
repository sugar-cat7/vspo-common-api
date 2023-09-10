package firestore

import (
	"context"
	"fmt"
	"time"

	entities "github.com/sugar-cat7/vspo-common-api/domain/entities/legacy"
	"github.com/sugar-cat7/vspo-common-api/domain/repositories"
	"github.com/sugar-cat7/vspo-common-api/util"
)

type LiveStreamRepository struct {
	client         repositories.FirestoreClient
	collectionName string
}

func NewLiveStreamRepository(client repositories.FirestoreClient) *LiveStreamRepository {
	return &LiveStreamRepository{
		client:         client,
		collectionName: liveStreamsCollectionName,
	}
}

func (r *LiveStreamRepository) FindAllByPeriod(start, end string) ([]*entities.OldVideo, error) {
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

	docs, err := r.client.Collection(r.collectionName).
		Where("scheduledStartTime", ">=", startTimeUTC).
		Where("scheduledStartTime", "<", endTimeUTC).
		Documents(ctx).GetAll()

	if err != nil {
		return nil, fmt.Errorf("failed to get documents from Firestore: %w", err)
	}

	var liveStreams []*entities.OldVideo
	for _, doc := range docs {
		var liveStream entities.OldVideo
		err = doc.DataTo(&liveStream)
		if err != nil {
			return nil, fmt.Errorf("failed to map document data to the provided struct: %w", err)
		}
		liveStreams = append(liveStreams, &liveStream)
	}

	return liveStreams, nil
}

func (r *LiveStreamRepository) UpdateInBatch(liveStreams []*entities.OldVideo) error {
	liveStreamChunks, err := util.Chunk(liveStreams, maxBatchSize)
	if err != nil {
		return fmt.Errorf("failed to chunk liveStreams: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	bulkWriter := r.client.BulkWriter(ctx)

	for _, liveStreamChunk := range liveStreamChunks {
		for _, liveStream := range liveStreamChunk {
			docRef := r.client.Collection(r.collectionName).Doc(liveStream.ID)
			err := bulkWriter.Update(docRef, liveStream.GetUpdate())
			if err != nil {
				return fmt.Errorf("failed to add update to bulkWriter in Firestore: %w", err)
			}
		}
	}

	bulkWriter.Flush()
	return nil
}

func (r *LiveStreamRepository) CreateInBatch(liveStreams []*entities.OldVideo) error {
	liveStreamChunks, err := util.Chunk(liveStreams, maxBatchSize)
	if err != nil {
		return fmt.Errorf("failed to chunk liveStreams: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	bulkWriter := r.client.BulkWriter(ctx)

	for _, liveStreamChunk := range liveStreamChunks {
		for _, liveStream := range liveStreamChunk {
			docRef := r.client.Collection(r.collectionName).Doc(liveStream.ID)
			err := bulkWriter.Create(docRef, liveStream)
			if err != nil {
				return fmt.Errorf("failed to add creation to bulkWriter in Firestore: %w", err)
			}
		}
	}

	bulkWriter.Flush()
	return nil
}
