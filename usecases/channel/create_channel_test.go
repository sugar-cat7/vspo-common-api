package usecases

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/mocks/factories"
	mocks "github.com/sugar-cat7/vspo-common-api/mocks/services"
)

func TestCreateChannel_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	channelIDs := []string{"channelID1", "channelID2"}
	newChannelData := []entities.Channel{
		factories.NewChannel(channelIDs[0]),
		factories.NewChannel(channelIDs[1]),
	}

	tests := []struct {
		name           string
		channelIDs     []string
		newChannelData []entities.Channel
		expectErr      bool
	}{
		{
			name:           "Success",
			channelIDs:     channelIDs,
			newChannelData: newChannelData,
			expectErr:      false,
		},

		// ... more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockYoutubeService := mocks.NewMockYouTubeService(ctrl)
			mockChannelService := mocks.NewMockChannelService(ctrl)

			mockYoutubeService.EXPECT().GetChannels(tt.channelIDs).Return(tt.newChannelData, nil).Times(1)
			mockChannelService.EXPECT().CreateChannelsInBatch(gomock.Not(gomock.Len(0))).Return(nil).Times(1)

			cc := &CreateChannel{
				youtubeService: mockYoutubeService,
				channelService: mockChannelService,
			}

			err := cc.Execute(tt.channelIDs)
			if tt.expectErr {
				assert.Error(t, err, "Expected error")
			} else {
				assert.NoError(t, err, "Expected no error")
			}
		})
	}
}
