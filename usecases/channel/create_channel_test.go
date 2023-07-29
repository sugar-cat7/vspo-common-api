package usecases

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/sugar-cat7/vspo-common-api/mocks/factories"
	mocks "github.com/sugar-cat7/vspo-common-api/mocks/services"

	"github.com/sugar-cat7/vspo-common-api/usecases/mappers"
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
			mockYoutubeService := mocks.NewMockYouTubeService(ctrl)
			mockChannelService := mocks.NewMockChannelService(ctrl)

			mockYoutubeService.EXPECT().GetChannels(tt.channelIDs).Return(tt.newChannelData, nil).Times(1)
			if tt.expectCreateError {
				mockChannelService.EXPECT().CreateChannelsInBatch(gomock.Not(gomock.Len(0))).Return(errors.New("create error")).Times(1)
			} else {
				mockChannelService.EXPECT().CreateChannelsInBatch(gomock.Not(gomock.Len(0))).Return(nil).Times(1)
			}

			cc := &CreateChannel{
				youtubeService: mockYoutubeService,
				channelService: mockChannelService,
				channelMapper:  &mappers.ChannelMapper{},
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
