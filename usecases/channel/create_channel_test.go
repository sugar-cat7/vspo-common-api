package usecases

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/sugar-cat7/vspo-common-api/mocks/factories"
	mock_port "github.com/sugar-cat7/vspo-common-api/mocks/ports"
	mock_repo "github.com/sugar-cat7/vspo-common-api/mocks/repositories"
	"google.golang.org/api/youtube/v3"
)

func TestCreateChannel_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name               string
		channelIDs         []string
		newChannelData     []*youtube.Channel
		expectCreateError  bool
		expectExecuteError bool
	}{
		{
			name:               "Success",
			channelIDs:         []string{"channelID1", "channelID2"},
			newChannelData:     []*youtube.Channel{factories.NewYoutubeChannel("channelID1"), factories.NewYoutubeChannel("channelID2")},
			expectCreateError:  false,
			expectExecuteError: false,
		},
		// ... more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockYoutubeService := mock_port.NewMockYouTubeService(ctrl)
			mockYoutubeService.EXPECT().GetChannels(tt.channelIDs).Return(tt.newChannelData, nil).Times(1)
			mockChannelRepository := mock_repo.NewMockChannelRepository(ctrl)
			mockChannelRepository.EXPECT().CreateInBatch(gomock.Any()).Return(nil).Times(1)
			cc := &CreateChannel{
				youtubeService:    mockYoutubeService,
				channelRepository: mockChannelRepository,
			}

			err := cc.Execute(tt.channelIDs)
			if tt.expectExecuteError {
				assert.Error(t, err, "Expected error")
			} else {
				assert.NoError(t, err, "Expected no error")
			}
		})
	}
}
