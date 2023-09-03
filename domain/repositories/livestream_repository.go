//go:generate mockgen -destination=../../mocks/repositories/mock_livestream_repository.go -package=mocks github.com/sugar-cat7/vspo-common-api/domain/repositories LiveStreamRepository
package repositories

import entities "github.com/sugar-cat7/vspo-common-api/domain/entities/legacy"

type LiveStreamRepository interface {
	FindAllByPeriod(start, end string) ([]*entities.OldVideo, error)
	UpdateInBatch(clips []*entities.OldVideo) error
	CreateInBatch(clips []*entities.OldVideo) error
}
