package usecases

import (
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/domain/repositories"
	"github.com/sugar-cat7/vspo-common-api/usecases/mappers"
)

// GetClipsByPeriod is a use case for getting all clips from Firestore.
type GetClipsByPeriod struct {
	clipRepository repositories.ClipRepository
}

// NewGetClipsByPeriod creates a new GetClipsByPeriod.
func NewGetClipsByPeriod(clipRepository repositories.ClipRepository) *GetClipsByPeriod {
	return &GetClipsByPeriod{
		clipRepository: clipRepository,
	}
}

// Execute gets all clips from Firestore.
func (g *GetClipsByPeriod) Execute(start, end string) ([]*entities.Video, error) {
	// Get all clips from Firestore
	clips, err := g.clipRepository.FindAllByPeriod(start, end)
	if err != nil {
		return nil, err
	}

	videos, err := mappers.ClipMapMultiple(clips)
	if err != nil {
		return nil, err
	}

	return videos, nil
}
