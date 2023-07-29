package usecases

import (
	"github.com/google/wire"
	"github.com/sugar-cat7/vspo-common-api/usecases/mappers"
)

func ProvideChannelMapper() *mappers.ChannelMapper {
	return &mappers.ChannelMapper{}
}

// Set is a Wire provider set that provides a song usecases.
var Set = wire.NewSet(NewCreateChannel, NewGetChannels, NewUpdateChannelsFromYoutube, ProvideChannelMapper)
