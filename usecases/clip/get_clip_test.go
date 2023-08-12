package usecases

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	entities2 "github.com/sugar-cat7/vspo-common-api/domain/entities/legacy"
	"github.com/sugar-cat7/vspo-common-api/mocks/factories"
	mocks "github.com/sugar-cat7/vspo-common-api/mocks/services"
	"github.com/sugar-cat7/vspo-common-api/usecases/mappers"
)

func TestNewGetClipsByPeriod(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClipService := mocks.NewMockClipService(ctrl)
	got := NewGetClipsByPeriod(mockClipService, &mappers.ClipMapper{})

	assert.NotNil(t, got, "NewGetClipsByPeriod() should not return nil")
}

func TestGetClipsByPeriod_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	c := &mappers.ClipMapper{}
	testVideo, _ := c.Map(factories.NewClip("testID"))

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
			mockClipService := mocks.NewMockClipService(ctrl)

			start, end := "2022-01-01", "2022-01-31"

			if tt.wantErr {
				mockClipService.EXPECT().FindAllByPeriod(start, end).Return(nil, errors.New("some error"))
			} else {
				mockClipService.EXPECT().FindAllByPeriod(start, end).Return(
					[]*entities2.Clip{factories.NewClip("testID")}, nil)

			}

			g := NewGetClipsByPeriod(mockClipService, &mappers.ClipMapper{})

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
