package usecases

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/mocks/factories"
	mocks "github.com/sugar-cat7/vspo-common-api/mocks/services"
)

func TestNewGetAllSongs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSongService := mocks.NewMockSongService(ctrl)
	got := NewGetAllSongs(mockSongService)

	assert.NotNil(t, got, "NewGetAllSongs() should not return nil")
}

func TestGetAllSongs_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testSong := factories.NewSong("testID")

	tests := []struct {
		name    string
		songs   []*entities.Song
		wantErr bool
	}{
		{
			name:    "Success",
			songs:   []*entities.Song{&testSong},
			wantErr: false,
		},
		{
			name:    "Failure",
			songs:   nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSongService := mocks.NewMockSongService(ctrl)

			if tt.wantErr {
				mockSongService.EXPECT().GetAllSongs().Return(tt.songs, errors.New("some error"))
			} else {
				mockSongService.EXPECT().GetAllSongs().Return(tt.songs, nil)
			}

			g := NewGetAllSongs(mockSongService)

			got, err := g.Execute()

			if tt.wantErr {
				assert.Error(t, err, "Expected error")
			} else {
				assert.NoError(t, err, "Unexpected error")
			}

			assert.Equal(t, tt.songs, got, "Expected and returned songs should be the same")
		})
	}
}
