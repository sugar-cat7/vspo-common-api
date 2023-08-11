package mappers

import (
	entities "github.com/sugar-cat7/vspo-common-api/domain/entities"
	entities2 "github.com/sugar-cat7/vspo-common-api/domain/entities/legacy"
)

// ClipMapper maps a Clip to a domain Video.
type ClipMapper struct{}

// Map maps a Clip to a domain Video.
func (cm *ClipMapper) Map(clip *entities2.Clip) (*entities.Video, error) {
	return &entities.Video{
		ID:          clip.ID,
		Title:       clip.Title,
		Description: clip.Description,
		ViewCount: entities.Views{
			Daily:   clip.NewViewCount.Daily,
			Weekly:  clip.NewViewCount.Weekly,
			Monthly: clip.NewViewCount.Monthly,
			Total:   clip.NewViewCount.Total,
		},
		PublishedAt: clip.CreatedAt,
		Thumbnails: entities.Thumbnails{
			Default: entities.Thumbnail{
				URL: clip.ThumbnailURL,
			},
		},
		ChannelTitle: clip.ChannelTitle,
		ChannelID:    clip.ChannelID,
		ChannelIcon:  clip.IconURL,
		Platform:     clip.Platform,
	}, nil
}

// MapMultiple maps multiple Clips to domain Videos.
func (cm *ClipMapper) MapMultiple(clips []*entities2.Clip) ([]*entities.Video, error) {
	videos := make([]*entities.Video, len(clips))
	for i, clip := range clips {
		video, err := cm.Map(clip)
		if err != nil {
			return nil, err
		}
		videos[i] = video
	}

	return videos, nil
}
