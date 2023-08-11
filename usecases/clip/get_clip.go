package usecases

import (
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/domain/services"
	"github.com/sugar-cat7/vspo-common-api/usecases/mappers"
)

// GetClipsByPeriod is a use case for getting all clips from Firestore.
type GetClipsByPeriod struct {
	clipService services.ClipService
	clipMapper  *mappers.ClipMapper
}

// NewGetClipsByPeriod creates a new GetClipsByPeriod.
func NewGetClipsByPeriod(clipService services.ClipService, clipMapper *mappers.ClipMapper) *GetClipsByPeriod {
	return &GetClipsByPeriod{
		clipService: clipService,
		clipMapper:  clipMapper,
	}
}

// Execute gets all clips from Firestore.
func (g *GetClipsByPeriod) Execute(start, end string) ([]*entities.Video, error) {
	// Get all clips from Firestore
	clips, err := g.clipService.FindAllByPeriod(start, end)
	if err != nil {
		return nil, err
	}

	videos, err := g.clipMapper.MapMultiple(clips)
	if err != nil {
		return nil, err
	}

	return videos, nil
}
