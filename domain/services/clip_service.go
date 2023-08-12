//go:generate mockgen -destination=../../mocks/services/mock_clip_service.go -package=mocks github.com/sugar-cat7/vspo-common-api/domain/services ClipService
package services

import (
	entities "github.com/sugar-cat7/vspo-common-api/domain/entities/legacy"
	"github.com/sugar-cat7/vspo-common-api/domain/repositories"
)

// ClipService is an interface for a clip service.
type ClipService interface {
	FindAllByPeriod(start, end string) ([]*entities.Clip, error)
	CreateClipsInBatch(clips []*entities.Clip) error
	UpdateClipsInBatch(clips []*entities.Clip) error
}

type clipService struct {
	repo repositories.ClipRepository
}

// NewClipService creates a new ClipService.
func NewClipService(repo repositories.ClipRepository) ClipService {
	return &clipService{repo: repo}
}

// FindAllByPeriod finds all clips by period.
func (c *clipService) FindAllByPeriod(start, end string) ([]*entities.Clip, error) {
	return c.repo.FindAllByPeriod(start, end)
}

// UpdateClipsInBatch updates multiple clips.
func (c *clipService) UpdateClipsInBatch(clips []*entities.Clip) error {
	return c.repo.UpdateInBatch(clips)
}

// CreateClipsInBatch creates multiple clips.
func (c *clipService) CreateClipsInBatch(clips []*entities.Clip) error {
	return c.repo.CreateInBatch(clips)
}
