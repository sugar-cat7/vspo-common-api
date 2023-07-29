package usecases

import (
	"github.com/google/wire"
	"github.com/sugar-cat7/vspo-common-api/usecases/mappers"
)

func ProvideSongMapper() *mappers.SongMapper {
	return &mappers.SongMapper{}
}

// Set is a Wire provider set that provides a song usecases.
var Set = wire.NewSet(NewGetAllSongs, NewCreateSong, NewUpdateSongs, ProvideSongMapper)
