package ports

import (
	"github.com/google/wire"
	youtube "github.com/sugar-cat7/vspo-common-api/infrastructure/api/youtube"
)

var Set = wire.NewSet(
	youtube.Set,
)
