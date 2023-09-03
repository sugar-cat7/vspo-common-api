package handlers

import (
	"github.com/google/wire"
	channel "github.com/sugar-cat7/vspo-common-api/infrastructure/http/handlers/channel"
	clip "github.com/sugar-cat7/vspo-common-api/infrastructure/http/handlers/clip"
	cron "github.com/sugar-cat7/vspo-common-api/infrastructure/http/handlers/cron"
	discord "github.com/sugar-cat7/vspo-common-api/infrastructure/http/handlers/discord"
	livestream "github.com/sugar-cat7/vspo-common-api/infrastructure/http/handlers/livestream"
	song "github.com/sugar-cat7/vspo-common-api/infrastructure/http/handlers/song"
)

var Set = wire.NewSet(
	song.Set,
	channel.Set,
	clip.Set,
	cron.Set,
	livestream.Set,
	discord.Set,
)
