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

func TestUpdateSongs_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	videoIDs := []string{"videoID1", "videoID2"}
	newVideoData := []*youtube.Video{
		factories.NewYoutubeVideo(videoIDs[0]),
		factories.NewYoutubeVideo(videoIDs[1]),
	}
	allSongsData := entities.Videos{
		factories.NewVideoPtr(videoIDs[0]),
		factories.NewVideoPtr(videoIDs[1]),
	}

	tests := []struct {
		name         string
		cronType     entities.CronType
		videoIDs     []string
		newVideoData []*youtube.Video
		allSongsData entities.Videos
		expectErr    bool
	}{
		{
			name:         "Success",
			cronType:     entities.None,
			videoIDs:     videoIDs,
			newVideoData: newVideoData,
			allSongsData: allSongsData,
			expectErr:    false,
		},
		// ... more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockYoutubeService := mock_port.NewMockYouTubeService(ctrl)
			mockSongRepository := mock_repo.NewMockSongRepository(ctrl)

			mockSongRepository.EXPECT().GetAll().Return(tt.allSongsData, nil).Times(1)
			mockYoutubeService.EXPECT().GetVideos(tt.videoIDs).Return(tt.newVideoData, nil).Times(1)
			mockSongRepository.EXPECT().UpdateInBatch(gomock.Not(gomock.Len(0))).Return(nil).Times(1)

			us := &UpdateSongs{
				youtubeService: mockYoutubeService,
				songRepository: mockSongRepository,
			}

			err := us.Execute(tt.cronType)
			if tt.expectErr {
				assert.Error(t, err, "Expected error")
			} else {
				assert.NoError(t, err, "Expected no error")
			}
		})
	}
}
