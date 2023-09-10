package usecases

import (
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/domain/repositories"
	"github.com/sugar-cat7/vspo-common-api/usecases/mappers"
)

// GetLiveStreamsByPeriod is a use case for getting all liveStreams from Firestore.
type GetLiveStreamsByPeriod struct {
	liveStreamRepository repositories.LiveStreamRepository
}

// NewGetLiveStreamsByPeriod creates a new GetLiveStreamsByPeriod.
func NewGetLiveStreamsByPeriod(liveStreamRepository repositories.LiveStreamRepository) *GetLiveStreamsByPeriod {
	return &GetLiveStreamsByPeriod{
		liveStreamRepository: liveStreamRepository,
	}
}

// Execute gets all liveStreams from Firestore.
func (g *GetLiveStreamsByPeriod) Execute(start, end, countryCode string) (entities.Videos, error) {
	// Get all liveStreams from Firestore
	liveStreams, err := g.liveStreamRepository.FindAllByPeriod(start, end)
	if err != nil {
		return nil, err
	}

	videos, err := mappers.LiveStreamMapMultiple(liveStreams)
	if err != nil {
		return nil, err
	}

	if videos.SetLocalTime(countryCode) != nil {
		return nil, err
	}

	return videos, nil
}
