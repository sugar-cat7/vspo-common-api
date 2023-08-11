package usecases

import (
	"github.com/google/wire"
	channel "github.com/sugar-cat7/vspo-common-api/usecases/channel"
	clip "github.com/sugar-cat7/vspo-common-api/usecases/clip"
	mappers "github.com/sugar-cat7/vspo-common-api/usecases/mappers"
	song "github.com/sugar-cat7/vspo-common-api/usecases/song"
)

var Set = wire.NewSet(
	song.Set,
	clip.Set,
	channel.Set,
	mappers.Set,
)
