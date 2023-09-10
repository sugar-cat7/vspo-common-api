package ports

import (
	"github.com/google/wire"
	discord "github.com/sugar-cat7/vspo-common-api/infrastructure/api/discord"
	slack "github.com/sugar-cat7/vspo-common-api/infrastructure/api/slack"
	youtube "github.com/sugar-cat7/vspo-common-api/infrastructure/api/youtube"
)

var Set = wire.NewSet(
	youtube.Set,
	discord.Set,
	slack.Set,
)
