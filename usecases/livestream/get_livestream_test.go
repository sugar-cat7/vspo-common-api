package usecases

import (
	"errors"
	"testing"
	"time"

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

	// Update testVideo with a predefined time zone for the purpose of the test
	// This assumes `ScheduledStartTime` and `ActualEndTime` are of type `time.Time`
	predefinedTimeZone, _ := time.LoadLocation("Asia/Tokyo")
	testVideo.ScheduledStartTime = testVideo.ScheduledStartTime.In(predefinedTimeZone)
	testVideo.ActualEndTime = testVideo.ActualEndTime.In(predefinedTimeZone)

	tests := []struct {
		name    string
		videos  entities.Videos
		wantErr bool
	}{
		{
			name:    "Success",
			videos:  entities.Videos{testVideo},
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

			start, end, countryCode := "2022-01-01", "2022-01-31", "JP"

			if tt.wantErr {
				mockLiveStreamRepository.EXPECT().FindAllByPeriod(start, end).Return(nil, errors.New("some error"))
			} else {
				mockLiveStreamRepository.EXPECT().FindAllByPeriod(start, end).Return(
					entities2.OldVideos{factories.NewLiveStream("testID")}, nil)
			}

			g := NewGetLiveStreamsByPeriod(mockLiveStreamRepository)

			got, err := g.Execute(start, end, countryCode)

			if tt.wantErr {
				assert.Error(t, err, "Expected error")
			} else {
				assert.NoError(t, err, "Unexpected error")
			}

			for i := range got {
				if got[i].ScheduledStartTime.Location() != predefinedTimeZone {
					t.Logf("Expected time zone: %v, but got: %v", predefinedTimeZone, got[i].ScheduledStartTime.Location())
				}
				if got[i].ActualEndTime.Location() != predefinedTimeZone {
					t.Logf("Expected time zone: %v, but got: %v", predefinedTimeZone, got[i].ActualEndTime.Location())
				}
				assert.True(t, got[i].ScheduledStartTime.Location().String() == predefinedTimeZone.String(), "Expected the time zones to match for ScheduledStartTime")
				assert.True(t, got[i].ActualEndTime.Location().String() == predefinedTimeZone.String(), "Expected the time zones to match for ActualEndTime")
			}
		})
	}
}
