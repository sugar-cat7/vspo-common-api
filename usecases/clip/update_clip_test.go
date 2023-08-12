package usecases

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	entities2 "github.com/sugar-cat7/vspo-common-api/domain/entities/legacy"
	"github.com/sugar-cat7/vspo-common-api/mocks/factories"
	mocks "github.com/sugar-cat7/vspo-common-api/mocks/services"
	"github.com/sugar-cat7/vspo-common-api/usecases/mappers"
	"google.golang.org/api/youtube/v3"
)

func TestUpdateClipsByPeriod_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	videoIDs := []string{"videoID1", "videoID2"}
	newVideoData := []*youtube.Video{
		factories.NewYoutubeVideo(videoIDs[0]),
		factories.NewYoutubeVideo(videoIDs[1]),
	}
	allClipsData := []*entities2.Clip{
		factories.NewClip(videoIDs[0]), // この関数はClipのポインタを返すように作成するか、既存の関数を使用します。
		factories.NewClip(videoIDs[1]),
	}

	tests := []struct {
		name         string
		cronType     entities.CronType
		videoIDs     []string
		newVideoData []*youtube.Video
		allClipsData []*entities2.Clip
		expectErr    bool
	}{
		{
			name:         "Success",
			cronType:     entities.Daily,
			videoIDs:     videoIDs,
			newVideoData: newVideoData,
			allClipsData: allClipsData,
			expectErr:    false,
		},
		// ... more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockYoutubeService := mocks.NewMockYouTubeService(ctrl)
			mockClipService := mocks.NewMockClipService(ctrl)
			mockClipMapper := &mappers.ClipMapper{}

			mockClipService.EXPECT().FindAllByPeriod(gomock.Any(), "").Return(allClipsData, nil).Times(1)
			mockYoutubeService.EXPECT().GetVideos(tt.videoIDs).Return(tt.newVideoData, nil).Times(1)
			mockClipService.EXPECT().UpdateClipsInBatch(gomock.Not(gomock.Len(0))).Return(nil).Times(1)

			us := NewUpdateClipsByPeriod(
				mockClipService,
				mockClipMapper,
				mockYoutubeService,
			)

			videos, err := us.Execute(tt.cronType)
			if tt.expectErr {
				assert.Error(t, err, "Expected error")
			} else {
				assert.NoError(t, err, "Expected no error")
				assert.Len(t, videos, len(tt.allClipsData), "Expected matching video lengths")
			}
		})
	}
}
