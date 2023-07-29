package util

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/mocks/factories"
	"google.golang.org/api/youtube/v3"
)

func TestUpdateViewCounts(t *testing.T) {
	videoID := "testVideoID"
	song := factories.NewSong(videoID)

	testCases := []struct {
		name        string
		cronType    entities.CronType
		viewCount   uint64
		songs       []*entities.Song
		expectError bool
	}{
		{
			name:        "Success_Daily",
			cronType:    entities.Daily,
			viewCount:   1000,
			songs:       []*entities.Song{&song},
			expectError: false,
		},
		{
			name:        "Success_Weekly",
			cronType:    entities.Weekly,
			viewCount:   1000,
			songs:       []*entities.Song{&song},
			expectError: false,
		},
		{
			name:        "Success_Monthly",
			cronType:    entities.Monthly,
			viewCount:   1000,
			songs:       []*entities.Song{&song},
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a video instance with the given viewCount
			video := factories.NewYoutubeVideo(videoID)
			video.Statistics.ViewCount = tc.viewCount
			videos := []*youtube.Video{video}
			updatedSongs, err := UpdateViewCounts(tc.cronType, videos, tc.songs)
			if (err != nil) != tc.expectError {
				t.Errorf("UpdateViewCounts() error = %v, expectError %v", err, tc.expectError)
				return
			}

			// Check the updated song view count
			for _, updatedSong := range updatedSongs {
				switch tc.cronType {
				case entities.Daily:
					if updatedSong.ViewCount.Daily != strconv.FormatUint(tc.viewCount, 10) {
						t.Errorf("Daily view count doesn't match: got = %v, want %v", updatedSong.ViewCount.Daily, strconv.FormatUint(tc.viewCount, 10))
					}
				case entities.Weekly:
					if updatedSong.ViewCount.Weekly != strconv.FormatUint(tc.viewCount, 10) {
						t.Errorf("Weekly view count doesn't match: got = %v, want %v", updatedSong.ViewCount.Weekly, strconv.FormatUint(tc.viewCount, 10))
					}
				case entities.Monthly:
					if updatedSong.ViewCount.Monthly != strconv.FormatUint(tc.viewCount, 10) {
						t.Errorf("Monthly view count doesn't match: got = %v, want %v", updatedSong.ViewCount.Monthly, strconv.FormatUint(tc.viewCount, 10))
					}
				}
				if updatedSong.ViewCount.Total != strconv.FormatUint(tc.viewCount, 10) {
					t.Errorf("Total view count doesn't match: got = %v, want %v", updatedSong.ViewCount.Total, strconv.FormatUint(tc.viewCount, 10))
				}
			}
		})
	}
}

func TestGetSongIDs(t *testing.T) {
	videoID := "testVideoID"
	song := factories.NewSong(videoID)

	testCases := []struct {
		name     string
		songs    []*entities.Song
		expected []string
	}{
		{
			name:     "Success",
			songs:    []*entities.Song{&song},
			expected: []string{videoID},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := GetSongIDs(tc.songs)
			assert.Equal(t, tc.expected, result)
		})
	}
}
