package usecases

import (
	"fmt"

	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/domain/ports"
	"github.com/sugar-cat7/vspo-common-api/domain/repositories"
	"github.com/sugar-cat7/vspo-common-api/usecases/mappers"
)

// UpdateClipsByPeriod is a use case for updateting all clips from Firestore.
type UpdateClipsByPeriod struct {
	clipRepository repositories.ClipRepository
	youtubeService ports.YouTubeService
}

// NewUpdateClipsByPeriod creates a new UpdateClipsByPeriod.
func NewUpdateClipsByPeriod(clipRepository repositories.ClipRepository, youtubeService ports.YouTubeService) *UpdateClipsByPeriod {
	return &UpdateClipsByPeriod{
		clipRepository: clipRepository,
		youtubeService: youtubeService,
	}
}

// Execute updates all clips from Firestore.
func (g *UpdateClipsByPeriod) Execute(cronType entities.CronType) (entities.Videos, error) {
	start, err := cronType.GetStartTime()
	if err != nil {
		return nil, err
	}

	// Update all clips from Firestore
	clips, err := g.clipRepository.FindAllByPeriod(start, "")
	if err != nil {
		return nil, err
	}

	videos, err := mappers.ClipMapMultiple(clips)
	if err != nil {
		return nil, err
	}

	videoIDs := videos.GetIDs()
	// Fetch video data from YouTube API
	ytVideos, err := g.youtubeService.GetVideos(videoIDs)
	if err != nil {
		return nil, err
	}

	if len(ytVideos) == 0 {
		return nil, fmt.Errorf("Fail Fetching Video")
	}

	err = mappers.BindAndUpdateMultiple(cronType, clips, mappers.MapToVideos(cronType, ytVideos))
	if err != nil {
		return nil, err
	}
	err = g.clipRepository.UpdateInBatch(clips)
	if err != nil {
		return nil, err
	}

	videos, err = mappers.ClipMapMultiple(clips)
	if err != nil {
		return nil, err
	}

	return videos, nil
}
