package usecases

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/mocks/factories"
	mock_repo "github.com/sugar-cat7/vspo-common-api/mocks/repositories"
	"github.com/sugar-cat7/vspo-common-api/util"
)

func TestGetChannels_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	channelIDs := []string{"channelID1", "channelID2"}
	fetchedChannels := []entities.Channel{
		factories.NewChannel(channelIDs[0]),
		factories.NewChannel(channelIDs[1]),
	}

	tests := []struct {
		name            string
		channelIDs      []string
		fetchedChannels []*entities.Channel
		expectErr       bool
	}{
		{
			name:            "Success",
			channelIDs:      channelIDs,
			fetchedChannels: util.ConvertToPtrSlice(fetchedChannels),
			expectErr:       false,
		},

		// ... more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockChannelRepository := mock_repo.NewMockChannelRepository(ctrl)
			mockChannelRepository.EXPECT().GetInBatch(tt.channelIDs).Return(tt.fetchedChannels, nil).Times(1)
			g := &GetChannels{
				channelRepository: mockChannelRepository,
			}

			channels, err := g.Execute(tt.channelIDs)
			if tt.expectErr {
				assert.Error(t, err, "Expected error")
			} else {
				assert.NoError(t, err, "Expected no error")
				assert.Equal(t, tt.fetchedChannels, channels)
			}
		})
	}
}
