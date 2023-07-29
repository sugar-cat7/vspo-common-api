package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/mocks/factories"
)

func TestConvertToSongs(t *testing.T) {
	videoID := "testVideoID"
	videoListResponse := factories.NewYTVideoListResponse(videoID)
	song := factories.NewSong(videoID)

	testCases := []struct {
		name        string
		input       []entities.YTVideoListResponse
		expected    []*entities.Song
		expectError bool
	}{
		{
			name:        "Success",
			input:       []entities.YTVideoListResponse{videoListResponse},
			expected:    []*entities.Song{&song},
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ConvertToSongs(tc.input)
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

func TestUpdateViewCounts(t *testing.T) {
	videoID := "testVideoID"
	song := factories.NewSong(videoID)

	testCases := []struct {
		name        string
		cronType    entities.CronType
		viewCount   string
		songs       []*entities.Song
		expectError bool
	}{
		{
			name:        "Success_Daily",
			cronType:    entities.Daily,
			viewCount:   "1000",
			songs:       []*entities.Song{&song},
			expectError: false,
		},
		{
			name:        "Success_Weekly",
			cronType:    entities.Weekly,
			viewCount:   "1000",
			songs:       []*entities.Song{&song},
			expectError: false,
		},
		{
			name:        "Success_Monthly",
			cronType:    entities.Monthly,
			viewCount:   "1000",
			songs:       []*entities.Song{&song},
			expectError: false,
		},
		{
			name:        "Success_None",
			cronType:    entities.None,
			viewCount:   "1000",
			songs:       []*entities.Song{&song},
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Set the view count in the video list response
			videoListResponse := factories.NewYTVideoListResponse(videoID)
			videoListResponse.Items[0].Statistics.ViewCount = tc.viewCount
			videoLists := []entities.YTVideoListResponse{videoListResponse}

			err := UpdateViewCounts(tc.cronType, videoLists, tc.songs)
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				// Check the updated song view count
				for _, song := range tc.songs {
					switch tc.cronType {
					case entities.Daily:
						assert.Equal(t, tc.viewCount, song.ViewCount.Daily)
					case entities.Weekly:
						assert.Equal(t, tc.viewCount, song.ViewCount.Weekly)
					case entities.Monthly:
						assert.Equal(t, tc.viewCount, song.ViewCount.Monthly)
					}
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
