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
func ProvideFirestoreClient() (repositories.FirestoreClient, error) {
	client, err := clientProvider()
	if err != nil {
		return nil, err
	}

	return &FirestoreClientImpl{Client: client}, nil
}

// ProvideSongRepository provides a song repository.
func ProvideSongRepository(client repositories.FirestoreClient, repo *SongRepository) repositories.SongRepository {
	return repo
}

// ProvideChannelRepository provides a channel repository.
func ProvideChannelRepository(client repositories.FirestoreClient, repo *ChannelRepository) repositories.ChannelRepository {
	return repo
}

// Set is a Wire provider set that provides a Firestore client and a song repository.
var Set = wire.NewSet(ProvideFirestoreClient, NewSongRepository, ProvideSongRepository, ProvideChannelRepository, NewChannelRepository)
