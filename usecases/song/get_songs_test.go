package usecases

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/mocks/factories"
	mock_repo "github.com/sugar-cat7/vspo-common-api/mocks/repositories"
)

func TestNewGetAllSongs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSongRepository := mock_repo.NewMockSongRepository(ctrl)
	got := NewGetAllSongs(mockSongRepository)

	assert.NotNil(t, got, "NewGetAllSongs() should not return nil")
}

func TestGetAllSongs_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testSong := factories.NewVideo("testID")

	tests := []struct {
		name    string
		songs   entities.Videos
		wantErr bool
	}{
		{
			name:    "Success",
			songs:   entities.Videos{&testSong},
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
			mockSongRepository := mock_repo.NewMockSongRepository(ctrl)

			if tt.wantErr {
				mockSongRepository.EXPECT().GetAll().Return(tt.songs, errors.New("some error"))
			} else {
				mockSongRepository.EXPECT().GetAll().Return(tt.songs, nil)
			}

			g := NewGetAllSongs(mockSongRepository)

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
