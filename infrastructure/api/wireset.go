package ports

import (
	"github.com/google/wire"
	discord "github.com/sugar-cat7/vspo-common-api/infrastructure/api/discord"
	youtube "github.com/sugar-cat7/vspo-common-api/infrastructure/api/youtube"
)

var Set = wire.NewSet(
	youtube.Set,
	discord.Set,
)
