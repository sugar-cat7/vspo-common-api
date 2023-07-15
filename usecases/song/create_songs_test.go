package usecases

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/mocks/factories"
	mocks "github.com/sugar-cat7/vspo-common-api/mocks/services"
)

func TestCreateSong_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	existingIDs := []string{"videoID2", "videoID3"}
	newVideoIDs := []string{"videoID1"}

	playlistResponse := []entities.YTYouTubePlaylistResponse{
		factories.NewYTPlayListListResponse(newVideoIDs[0]),
	}

	newVideoData := []entities.YTVideoListResponse{
		factories.NewYTVideoListResponse(newVideoIDs[0]),
	}

	tests := []struct {
		name              string
		playlistResponses []entities.YTYouTubePlaylistResponse
		existingIDs       []string
		newVideoData      []entities.YTVideoListResponse
		newVideoIDs       []string
		expectErr         bool
	}{
		{
			name:              "Success",
			playlistResponses: playlistResponse,
			existingIDs:       existingIDs,
			newVideoData:      newVideoData,
			newVideoIDs:       newVideoIDs,
			expectErr:         false,
		},

		// ... more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockYoutubeService := mocks.NewMockYouTubeService(ctrl)
			mockSongService := mocks.NewMockSongService(ctrl)

			mockYoutubeService.EXPECT().GetPlaylists().Return(tt.playlistResponses, nil).Times(1)
			mockSongService.EXPECT().GetSongIDs().Return(tt.existingIDs, nil).Times(1)

			if tt.newVideoData != nil && len(tt.newVideoData) != 0 {
				mockYoutubeService.EXPECT().GetSongs(tt.newVideoIDs).Return(tt.newVideoData, nil).Times(1)
				mockSongService.EXPECT().CreateSongsInBatch(gomock.Not(gomock.Len(0))).Return(nil).Times(1)
			}

			cs := &CreateSong{
				youtubeService: mockYoutubeService,
				songService:    mockSongService,
			}

			err := cs.Execute()
			if tt.expectErr {
				assert.Error(t, err, "Expected error")
			} else {
				assert.NoError(t, err, "Expected no error")
			}
		})
	}
}
