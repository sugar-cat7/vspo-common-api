package usecases

import (
	"github.com/google/wire"
)

// Set is a Wire provider set that provides a song usecases.
var Set = wire.NewSet(NewDiscordSendMessage)
