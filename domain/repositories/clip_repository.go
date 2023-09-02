//go:generate mockgen -destination=../../mocks/repositories/mock_clip_repository.go -package=mocks github.com/sugar-cat7/vspo-common-api/domain/repositories ClipRepository
package repositories

import entities "github.com/sugar-cat7/vspo-common-api/domain/entities/legacy"

type ClipRepository interface {
	FindAllByPeriod(start, end string) ([]*entities.Clip, error)
	UpdateInBatch(clips []*entities.Clip) error
	CreateInBatch(clips []*entities.Clip) error
}
