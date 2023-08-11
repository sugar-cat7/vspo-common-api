package handlers

import (
	"github.com/google/wire"
	channel "github.com/sugar-cat7/vspo-common-api/infrastructure/http/handlers/channel"
	clip "github.com/sugar-cat7/vspo-common-api/infrastructure/http/handlers/clip"
	song "github.com/sugar-cat7/vspo-common-api/infrastructure/http/handlers/song"
)

var Set = wire.NewSet(
	song.Set,
	channel.Set,
	clip.Set,
)
