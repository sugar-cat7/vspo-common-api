package usecases

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	mocks "github.com/sugar-cat7/vspo-common-api/mocks/services"
)

func TestUpdateSongsFromYoutube_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name     string
		cronType entities.CronType
		wantErr  bool
	}{
		{
			name:    "Success",
			wantErr: false,
		},
		{
			name:    "Failure",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockYoutubeService := mocks.NewMockYouTubeService(ctrl)
			mockSongService := mocks.NewMockSongService(ctrl)

			if tt.wantErr {
				mockSongService.EXPECT().GetAllSongs().Return([]*entities.Song{}, nil)
				mockYoutubeService.EXPECT().GetSongs(gomock.Any()).Return(nil, errors.New("error"))
			} else {
				mockSongService.EXPECT().GetAllSongs().Return([]*entities.Song{}, nil)
				mockYoutubeService.EXPECT().GetSongs(gomock.Any()).Return([]entities.YTVideoListResponse{}, nil)
				mockSongService.EXPECT().UpdateSongsInBatch(gomock.Any()).Return(nil)
			}

			u := NewUpdateSongsFromYoutube(mockYoutubeService, mockSongService)

			err := u.Execute(tt.cronType)

			if tt.wantErr {
				assert.Error(t, err, "Expected error")
			} else {
				assert.NoError(t, err, "Unexpected error")
			}
		})
	}
}
