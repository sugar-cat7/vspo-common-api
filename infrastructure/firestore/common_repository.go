package firestore

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/sugar-cat7/vspo-common-api/util"

	"google.golang.org/api/iterator"
)

// Doc represents an entity that can be stored in Firestore.
type Doc interface {
	GetID() string
	GetUpdate() []firestore.Update
}

// Repository is a Firestore implementation of a repository.
type Repository struct {
	client         FirestoreClient
	collectionName string
}

// NewRepository creates a new Repository.
func NewRepository(client FirestoreClient, collectionName string) *Repository {
	return &Repository{
		client:         client,
		collectionName: collectionName,
	}
}

// Create creates a new document in Firestore.
func (r *Repository) Create(doc Doc) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.client.Collection(r.collectionName).Doc(doc.GetID()).Set(ctx, doc)

	if err != nil {
		return fmt.Errorf("failed to create document in Firestore: %w", err)
	}

	return nil
}

// CreateInBatch creates multiple new documents in Firestore using batch operation.
func (r *Repository) CreateInBatch(docs []Doc) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Split into chunks of 500 (Firestore batch limit)
	chunks, err := util.Chunk[Doc](docs, 500)
	if err != nil {
		return fmt.Errorf("failed to chunk docs: %w", err)
	}

	for _, chunk := range chunks {
		batch := r.client.Batch()

		for _, doc := range chunk {
			docRef := r.client.Collection(r.collectionName).Doc(doc.GetID())
			batch.Set(docRef, doc, firestore.MergeAll)
		}

		_, err := batch.Commit(ctx)
		if err != nil {
			return fmt.Errorf("failed to create documents in Firestore: %w", err)
		}
	}

	return nil
}

// GetAll retrieves all documents from Firestore.
func (r *Repository) GetAll(docs any) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	iter, err := r.client.Collection(r.collectionName).Documents(ctx)
	if err != nil {
		return fmt.Errorf("failed to get all documents from Firestore: %w", err)
	}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to iterate over documents in Firestore: %w", err)
		}

		err = doc.DataTo(docs)
		if err != nil {
			return fmt.Errorf("failed to map document data to the provided struct: %w", err)
		}
	}

	return nil
}

// GetByID retrieves a document by ID from Firestore.
func (r *Repository) GetByID(id string, docData any) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	doc, err := r.client.Collection(r.collectionName).Doc(id).Get(ctx)
	if err != nil {
		return fmt.Errorf("failed to get document by ID from Firestore: %w", err)
	}

	err = doc.DataTo(docData)
	if err != nil {
		return fmt.Errorf("failed to map document data to the provided struct: %w", err)
	}

	return nil
}

// Update updates a document in Firestore.
func (r *Repository) Update(doc Doc) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.client.Collection(r.collectionName).Doc(doc.GetID()).Update(ctx, doc.GetUpdate())

	if err != nil {
		return fmt.Errorf("failed to update document in Firestore: %w", err)
	}

	return nil
}

// Delete deletes a document from Firestore.
func (r *Repository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.client.Collection(r.collectionName).Doc(id).Delete(ctx)

	if err != nil {
		return fmt.Errorf("failed to delete document from Firestore: %w", err)
	}

	return nil
}

// ChunkDocs splits a slice of Doc into chunks.
func (r *Repository) ChunkDocs(docs []Doc, chunkSize int) [][]Doc {
	var chunks [][]Doc
	for i := 0; i < len(docs); i += chunkSize {
		end := i + chunkSize
		if end > len(docs) {
			end = len(docs)
		}
		chunks = append(chunks, docs[i:end])
	}
	return chunks
}

// UpdateInBatch updates multiple documents in Firestore using batch operation.
func (r *Repository) UpdateInBatch(docs []Doc) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Split into chunks of 500
	chunks, err := util.Chunk[Doc](docs, 500)
	if err != nil {
		return fmt.Errorf("failed to chunk docs: %w", err)
	}

	for _, chunk := range chunks {
		batch := r.client.Batch()

		for _, doc := range chunk {
			docRef := r.client.Collection(r.collectionName).Doc(doc.GetID())
			batch.Update(docRef, doc.GetUpdate())
		}

		_, err := batch.Commit(ctx)
		if err != nil {
			return fmt.Errorf("failed to update documents in Firestore: %w", err)
		}
	}

	return nil
}

// UpsertInBatch upserts multiple documents in Firestore using batch operation.
func (r *Repository) UpsertInBatch(docs []Doc) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Split into chunks of 500
	chunks, err := util.Chunk[Doc](docs, 500)
	if err != nil {
		return fmt.Errorf("failed to chunk docs: %w", err)
	}

	for _, chunk := range chunks {
		batch := r.client.Batch()

		for _, doc := range chunk {
			docRef := r.client.Collection(r.collectionName).Doc(doc.GetID())
			batch.Set(docRef, doc.GetUpdate(), firestore.MergeAll)
		}

		_, err := batch.Commit(ctx)
		if err != nil {
			return fmt.Errorf("failed to upsert documents in Firestore: %w", err)
		}
	}

	return nil
}
