package firestore

import (
	"cloud.google.com/go/firestore"
	"github.com/google/wire"
	"github.com/sugar-cat7/vspo-common-api/config"
	"github.com/sugar-cat7/vspo-common-api/domain/repositories"
)

var clientProvider = func() (*firestore.Client, error) {
	//TODO: Add Firestore client provider
	return config.Config()
}

// SetClientProvider is used in testing to replace the client provider function.
func SetClientProvider(provider func() (*firestore.Client, error)) {
	clientProvider = provider
}

// ProvideFirestoreClient provides a Firestore client.
func ProvideFirestoreClient() (*firestore.Client, error) {
	return clientProvider()
}

// ProvideSongRepository provides a song repository.
func ProvideSongRepository(client *firestore.Client, repo *SongRepository) repositories.SongRepository {
	return repo
}

// Set is a Wire provider set that provides a Firestore client.
var Set = wire.NewSet(NewRepository, NewSongRepository, ProvideFirestoreClient, ProvideSongRepository, NewFirestoreClientImpl)
