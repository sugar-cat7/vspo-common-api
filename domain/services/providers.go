package services

import (
	"github.com/google/wire"
)

// Set is a Wire provider set that provides a song service.
var Set = wire.NewSet(
	NewSongService,
	NewYouTubeService,
	NewChannelService,
	NewClipService,
)
