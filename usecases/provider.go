package usecases

import (
	"github.com/google/wire"
	channel "github.com/sugar-cat7/vspo-common-api/usecases/channel"
	clip "github.com/sugar-cat7/vspo-common-api/usecases/clip"
	discord "github.com/sugar-cat7/vspo-common-api/usecases/discord"
	livestream "github.com/sugar-cat7/vspo-common-api/usecases/livestream"
	song "github.com/sugar-cat7/vspo-common-api/usecases/song"
)

var Set = wire.NewSet(
	song.Set,
	clip.Set,
	channel.Set,
	livestream.Set,
	discord.Set,
)
