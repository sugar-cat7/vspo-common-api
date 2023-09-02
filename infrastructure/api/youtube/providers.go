package ports

import (
	"github.com/google/wire"
)

// Set is a Wire provider set that provides a clip usecases.
var Set = wire.NewSet(NewYouTubeService)
