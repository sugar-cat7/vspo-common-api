package repositories

import (
	"context"

	"cloud.google.com/go/firestore"
)

// FirestoreClient is the interface that wraps firestore client's method set.
type FirestoreClient interface {
	Collection(string) *firestore.CollectionRef
	BulkWriter(context.Context) BulkWriter
	GetAll(string) ([]*firestore.DocumentSnapshot, error)
	GetByID(string, string) (*firestore.DocumentSnapshot, error)
	Create(string, string, interface{}) (*firestore.WriteResult, error)
	Update(string, string, interface{}) (*firestore.WriteResult, error)
	Delete(string, string) (*firestore.WriteResult, error)
	GetInBatch(string, []string) ([]*firestore.DocumentSnapshot, error)
}

// BulkWriter is the interface that wraps firestore's BulkWriter method set.
type BulkWriter interface {
	Create(*firestore.DocumentRef, interface{}) error
	Update(*firestore.DocumentRef, []firestore.Update) error
	Flush()
}
