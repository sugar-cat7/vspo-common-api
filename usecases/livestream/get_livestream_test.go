package usecases

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	entities2 "github.com/sugar-cat7/vspo-common-api/domain/entities/legacy"
	"github.com/sugar-cat7/vspo-common-api/mocks/factories"
	mock_repo "github.com/sugar-cat7/vspo-common-api/mocks/repositories"
	"github.com/sugar-cat7/vspo-common-api/usecases/mappers"
)

func TestGetLiveStreamsByPeriod_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testVideo, _ := mappers.LiveStreamMap(factories.NewLiveStream("testID"))

	tests := []struct {
		name    string
		videos  []*entities.Video
		wantErr bool
	}{
		{
			name:    "Success",
			videos:  []*entities.Video{testVideo},
			wantErr: false,
		},
		{
			name:    "Failure",
			videos:  nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLiveStreamRepository := mock_repo.NewMockLiveStreamRepository(ctrl)

			start, end := "2022-01-01", "2022-01-31"

			if tt.wantErr {
				mockLiveStreamRepository.EXPECT().FindAllByPeriod(start, end).Return(nil, errors.New("some error"))
			} else {
				mockLiveStreamRepository.EXPECT().FindAllByPeriod(start, end).Return(
					[]*entities2.OldVideo{factories.NewLiveStream("testID")}, nil)

			}

			g := NewGetLiveStreamsByPeriod(mockLiveStreamRepository)

			got, err := g.Execute(start, end)

			if tt.wantErr {
				assert.Error(t, err, "Expected error")
			} else {
				assert.NoError(t, err, "Unexpected error")
			}

			assert.Equal(t, tt.videos, got, "Expected and returned videos should be the same")
		})
	}
}
