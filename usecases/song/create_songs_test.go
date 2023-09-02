package usecases

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/mocks/factories"
	mock_port "github.com/sugar-cat7/vspo-common-api/mocks/ports"
	mock_repo "github.com/sugar-cat7/vspo-common-api/mocks/repositories"
	"google.golang.org/api/youtube/v3"
)

func TestCreateSong_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	videoIDs := []string{"videoID1", "videoID2"}
	newVideoData := []*youtube.Video{
		factories.NewYoutubeVideo(videoIDs[0]),
		factories.NewYoutubeVideo(videoIDs[1]),
	}

	newSongData := []*entities.Video{
		factories.NewVideoPtr(videoIDs[0]),
		factories.NewVideoPtr(videoIDs[1]),
	}

	tests := []struct {
		name         string
		videoIDs     []string
		newVideoData []*youtube.Video
		newSongData  []*entities.Video
		expectErr    bool
	}{
		{
			name:         "Success",
			videoIDs:     videoIDs,
			newVideoData: newVideoData,
			newSongData:  newSongData,
			expectErr:    false,
		},

		// ... more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockYoutubeService := mock_port.NewMockYouTubeService(ctrl)
			mockSongRepository := mock_repo.NewMockSongRepository(ctrl)

			mockYoutubeService.EXPECT().GetVideos(tt.videoIDs).Return(tt.newVideoData, nil).Times(1)
			mockSongRepository.EXPECT().CreateInBatch(gomock.Not(gomock.Len(0))).Return(nil).Times(1)

			u := &CreateSong{
				youtubeService: mockYoutubeService,
				songRepository: mockSongRepository,
			}

			err := u.Execute(tt.videoIDs)
			if tt.expectErr {
				assert.Error(t, err, "Expected error")
			} else {
				assert.NoError(t, err, "Expected no error")
			}
		})
	}
}
