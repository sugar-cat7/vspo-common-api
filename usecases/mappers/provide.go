package mappers

import "github.com/google/wire"

func ProvideSongMapper() *SongMapper {
	return &SongMapper{}
}

func ProvideChannelMapper() *ChannelMapper {
	return &ChannelMapper{}
}

func ProvideClipMapper() *ClipMapper {
	return &ClipMapper{}
}

var Set = wire.NewSet(ProvideChannelMapper, ProvideClipMapper, ProvideSongMapper)
