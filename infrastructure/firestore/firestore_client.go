package firestore

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/sugar-cat7/vspo-common-api/domain/repositories"
	"google.golang.org/api/iterator"
)

type FirestoreClientImpl struct {
	Client *firestore.Client
}

func (c *FirestoreClientImpl) Collection(collectionPath string) *firestore.CollectionRef {
	return c.Client.Collection(collectionPath)
}

func (c *FirestoreClientImpl) BulkWriter(ctx context.Context) repositories.BulkWriter {
	return &BulkWriterImpl{BulkWriter: c.Client.BulkWriter(ctx)}
}

func (c *FirestoreClientImpl) GetAll(collectionName string) (docs []*firestore.DocumentSnapshot, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	iter := c.Client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		docs = append(docs, doc)
	}

	return docs, nil
}

func (c *FirestoreClientImpl) GetByID(collectionName, id string) (*firestore.DocumentSnapshot, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return c.Client.Collection(collectionName).Doc(id).Get(ctx)
}

func (c *FirestoreClientImpl) Create(collectionName, id string, data interface{}) (*firestore.WriteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return c.Client.Collection(collectionName).Doc(id).Set(ctx, data)
}

func (c *FirestoreClientImpl) Update(collectionName, id string, data interface{}) (*firestore.WriteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return c.Client.Collection(collectionName).Doc(id).Set(ctx, data, firestore.MergeAll)
}

func (c *FirestoreClientImpl) Delete(collectionName, id string) (*firestore.WriteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return c.Client.Collection(collectionName).Doc(id).Delete(ctx)
}

func (c *FirestoreClientImpl) GetInBatch(collectionName string, ids []string) (docs []*firestore.DocumentSnapshot, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for _, id := range ids {
		doc, err := c.Client.Collection(collectionName).Doc(id).Get(ctx)
		if err != nil {
			return nil, err
		}
		docs = append(docs, doc)
	}

	return docs, nil
}

type BulkWriterImpl struct {
	BulkWriter *firestore.BulkWriter
}

func (b *BulkWriterImpl) Create(doc *firestore.DocumentRef, data interface{}) error {
	_, err := b.BulkWriter.Create(doc, data)
	return err
}

func (b *BulkWriterImpl) Update(doc *firestore.DocumentRef, updates []firestore.Update) error {
	_, err := b.BulkWriter.Update(doc, updates)
	return err
}

func (b *BulkWriterImpl) Flush() {
	b.BulkWriter.Flush()
}

const (
	channelsCollectionName    = "channels"
	songsCollectionName       = "songs"
	clipsCollectionName       = "clips"
	liveStreamsCollectionName = "livestreams"
)

var maxBatchSize = 500
