package mappers

import (
	"fmt"

	entities "github.com/sugar-cat7/vspo-common-api/domain/entities"
	entities2 "github.com/sugar-cat7/vspo-common-api/domain/entities/legacy"
)

// ClipMap maps a Clip to a domain Video.
func ClipMap(clip *entities2.OldVideo) (*entities.Video, error) {
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

// ClipMapMultiple maps multiple Clips to domain Videos.
func ClipMapMultiple(clips []*entities2.OldVideo) ([]*entities.Video, error) {
	videos := make([]*entities.Video, len(clips))
	for i, clip := range clips {
		video, err := ClipMap(clip)
		if err != nil {
			return nil, err
		}
		videos[i] = video
	}

	return videos, nil
}

// BindAndUpdate binds a Video to a Clip and updates the Clip.
func BindAndUpdate(cronType entities.CronType, clip *entities2.OldVideo, video *entities.Video,
) error {
	clip.ThumbnailURL = video.Thumbnails.Medium.URL
	clip.ViewCount = video.ViewCount.Total
	clip.NewViewCount.Total = video.ViewCount.Total
	switch cronType {
	case entities.Daily:
		clip.NewViewCount.Daily = video.ViewCount.Daily
	case entities.Weekly:
		clip.NewViewCount.Weekly = video.ViewCount.Weekly
	case entities.Monthly:
		clip.NewViewCount.Monthly = video.ViewCount.Monthly
	}
	return nil
}

// BindAndUpdateMultiple binds multiple Videos to multiple Clips and updates the Clips.
func BindAndUpdateMultiple(cronType entities.CronType, clips []*entities2.OldVideo, videos []*entities.Video) error {
	if len(clips) != len(videos) {
		return fmt.Errorf("Length of clips and videos must be the same")
	}
	for i, clip := range clips {
		video := videos[i]
		err := BindAndUpdate(cronType, clip, video)
		if err != nil {
			return err
		}
	}
	return nil
}
