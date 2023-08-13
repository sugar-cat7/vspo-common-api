package usecases

import (
	"fmt"

	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/domain/services"
	"github.com/sugar-cat7/vspo-common-api/usecases/mappers"
	"github.com/sugar-cat7/vspo-common-api/util"
)

// UpdateClipsByPeriod is a use case for updateting all clips from Firestore.
type UpdateClipsByPeriod struct {
	clipService    services.ClipService
	clipMapper     *mappers.ClipMapper
	youtubeService services.YouTubeService
}

// NewUpdateClipsByPeriod creates a new UpdateClipsByPeriod.
func NewUpdateClipsByPeriod(clipService services.ClipService, clipMapper *mappers.ClipMapper, youtubeService services.YouTubeService) *UpdateClipsByPeriod {
	return &UpdateClipsByPeriod{
		clipService:    clipService,
		clipMapper:     clipMapper,
		youtubeService: youtubeService,
	}
}

// Execute updates all clips from Firestore.
func (g *UpdateClipsByPeriod) Execute(cronType entities.CronType) ([]*entities.Video, error) {
	start, err := util.GetStartTime(cronType)
	if err != nil {
		return nil, err
	}

	// Update all clips from Firestore
	clips, err := g.clipService.FindAllByPeriod(start, "")
	if err != nil {
		return nil, err
	}

	videos, err := g.clipMapper.MapMultiple(clips)
	if err != nil {
		return nil, err
	}

	videoIDs := util.GetVideoIDs(videos)
	// Fetch video data from YouTube API
	ytVideos, err := g.youtubeService.GetVideos(videoIDs)
	if err != nil {
		return nil, err
	}

	if len(ytVideos) == 0 {
		return nil, fmt.Errorf("Fail Fetching Video")
	}

	updatedVideos, err := util.UpdateViewCounts(cronType, ytVideos, videos)
	if err != nil {
		return nil, err
	}

	err = g.clipMapper.BindAndUpdateMultiple(cronType, clips, updatedVideos)
	if err != nil {
		return nil, err
	}
	err = g.clipService.UpdateClipsInBatch(clips)
	if err != nil {
		return nil, err
	}

	videos, err = g.clipMapper.MapMultiple(clips)
	if err != nil {
		return nil, err
	}

	return videos, nil
}
