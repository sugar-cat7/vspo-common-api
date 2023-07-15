//go:generate mockgen -destination=../../mocks/repositories/mock_firestore.go -package=mocks github.com/sugar-cat7/vspo-common-api/infrastructure/firestore FirestoreClient,CollectionRef,DocumentRef,Batch
package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
)

// FirestoreClient is an interface for a Firestore client.
type FirestoreClient interface {
	Collection(string) CollectionRef
	Doc(string) DocumentRef
	Batch() Batch
}

// CollectionRef is an interface for a Firestore collection reference.
type CollectionRef interface {
	Doc(string) DocumentRef
	Documents(ctx context.Context) (*firestore.DocumentIterator, error)
}

// DocumentRef is an interface for a Firestore document reference.
type DocumentRef interface {
	Set(context.Context, interface{}, ...firestore.SetOption) (*firestore.WriteResult, error)
	Delete(context.Context, ...firestore.Precondition) (*firestore.WriteResult, error)
	Update(context.Context, []firestore.Update, ...firestore.Precondition) (*firestore.WriteResult, error)
	Get(context.Context, ...firestore.DocumentSnapshot) (*firestore.DocumentSnapshot, error)
}

// Batch is an interface for a Firestore batch.
type Batch interface {
	Set(DocumentRef, interface{}, ...firestore.SetOption) Batch
	Delete(DocumentRef, ...firestore.Precondition) Batch
	Update(DocumentRef, []firestore.Update, ...firestore.Precondition) Batch
	Commit(context.Context) ([]*firestore.WriteResult, error)
}

// FirestoreClientImpl is a Firestore implementation of a client.
type FirestoreClientImpl struct {
	client *firestore.Client
}

// NewFirestoreClientImpl creates a new FirestoreClientImpl.
func NewFirestoreClientImpl(client *firestore.Client) FirestoreClient {
	return &FirestoreClientImpl{client: client}
}

// Collection returns a Firestore collection reference.
func (f *FirestoreClientImpl) Collection(id string) CollectionRef {
	return &CollectionRefImpl{ref: f.client.Collection(id)}
}

// Doc returns a Firestore document reference.
func (f *FirestoreClientImpl) Doc(path string) DocumentRef {
	return &DocumentRefImpl{ref: f.client.Doc(path)}
}

// Batch returns a Firestore batch.
func (f *FirestoreClientImpl) Batch() Batch {
	return &BatchImpl{batch: f.client.Batch()}
}

// CollectionRefImpl is a Firestore implementation of a collection reference.
type CollectionRefImpl struct {
	ref *firestore.CollectionRef
}

// Doc returns a Firestore document reference.
func (c *CollectionRefImpl) Doc(path string) DocumentRef {
	return &DocumentRefImpl{ref: c.ref.Doc(path)}
}

// Documents returns a Firestore document iterator.
func (c *CollectionRefImpl) Documents(ctx context.Context) (*firestore.DocumentIterator, error) {
	return c.ref.Documents(ctx), nil
}

// DocumentRefImpl is a Firestore implementation of a document reference.
type DocumentRefImpl struct {
	ref *firestore.DocumentRef
}

// Set sets a document in Firestore.
func (d *DocumentRefImpl) Set(ctx context.Context, doc interface{}, opts ...firestore.SetOption) (*firestore.WriteResult, error) {
	return d.ref.Set(ctx, doc, opts...)
}

// Delete deletes a document from Firestore.
func (d *DocumentRefImpl) Delete(ctx context.Context, preconds ...firestore.Precondition) (*firestore.WriteResult, error) {
	return d.ref.Delete(ctx, preconds...)
}

// Update updates a document in Firestore.
func (d *DocumentRefImpl) Update(ctx context.Context, updates []firestore.Update, preconds ...firestore.Precondition) (*firestore.WriteResult, error) {
	return d.ref.Update(ctx, updates, preconds...)
}

// Get retrieves a document from Firestore.
func (d *DocumentRefImpl) Get(ctx context.Context, snap ...firestore.DocumentSnapshot) (*firestore.DocumentSnapshot, error) {
	return d.ref.Get(ctx)
}

// BatchImpl is a Firestore implementation of a batch.
type BatchImpl struct {
	batch *firestore.WriteBatch
}

// Set sets a document in Firestore.
func (b *BatchImpl) Set(doc DocumentRef, data interface{}, opts ...firestore.SetOption) Batch {
	docRefImpl := doc.(*DocumentRefImpl)
	b.batch.Set(docRefImpl.ref, data, opts...)
	return b
}

// Delete deletes a document from Firestore.
func (b *BatchImpl) Delete(doc DocumentRef, preconds ...firestore.Precondition) Batch {
	docRefImpl := doc.(*DocumentRefImpl)
	b.batch.Delete(docRefImpl.ref, preconds...)
	return b
}

// Update updates a document in Firestore.
func (b *BatchImpl) Update(doc DocumentRef, updates []firestore.Update, preconds ...firestore.Precondition) Batch {
	docRefImpl := doc.(*DocumentRefImpl)
	b.batch.Update(docRefImpl.ref, updates, preconds...)
	return b
}

// Commit commits a batch to Firestore.
func (b *BatchImpl) Commit(ctx context.Context) ([]*firestore.WriteResult, error) {
	return b.batch.Commit(ctx)
}
