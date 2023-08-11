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
	video := factories.NewVideo(videoID)

	testCases := []struct {
		name        string
		cronType    entities.CronType
		viewCount   uint64
		songs       []*entities.Video
		expectError bool
	}{
		{
			name:        "Success_Daily",
			cronType:    entities.Daily,
			viewCount:   1000,
			songs:       []*entities.Video{&video},
			expectError: false,
		},
		{
			name:        "Success_Weekly",
			cronType:    entities.Weekly,
			viewCount:   1000,
			songs:       []*entities.Video{&video},
			expectError: false,
		},
		{
			name:        "Success_Monthly",
			cronType:    entities.Monthly,
			viewCount:   1000,
			songs:       []*entities.Video{&video},
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a video instance with the given viewCount
			video := factories.NewYoutubeVideo(videoID)
			video.Statistics.ViewCount = tc.viewCount
			videos := []*youtube.Video{video}
			updatedVideos, err := UpdateViewCounts(tc.cronType, videos, tc.songs)
			if (err != nil) != tc.expectError {
				t.Errorf("UpdateViewCounts() error = %v, expectError %v", err, tc.expectError)
				return
			}

			// Check the updated video view count
			for _, updatedVideo := range updatedVideos {
				switch tc.cronType {
				case entities.Daily:
					if updatedVideo.ViewCount.Daily != strconv.FormatUint(tc.viewCount, 10) {
						t.Errorf("Daily view count doesn't match: got = %v, want %v", updatedVideo.ViewCount.Daily, strconv.FormatUint(tc.viewCount, 10))
					}
				case entities.Weekly:
					if updatedVideo.ViewCount.Weekly != strconv.FormatUint(tc.viewCount, 10) {
						t.Errorf("Weekly view count doesn't match: got = %v, want %v", updatedVideo.ViewCount.Weekly, strconv.FormatUint(tc.viewCount, 10))
					}
				case entities.Monthly:
					if updatedVideo.ViewCount.Monthly != strconv.FormatUint(tc.viewCount, 10) {
						t.Errorf("Monthly view count doesn't match: got = %v, want %v", updatedVideo.ViewCount.Monthly, strconv.FormatUint(tc.viewCount, 10))
					}
				}
				if updatedVideo.ViewCount.Total != strconv.FormatUint(tc.viewCount, 10) {
					t.Errorf("Total view count doesn't match: got = %v, want %v", updatedVideo.ViewCount.Total, strconv.FormatUint(tc.viewCount, 10))
				}
			}
		})
	}
}

func TestGetVideoIDs(t *testing.T) {
	videoID := "testVideoID"
	video := factories.NewVideo(videoID)

	testCases := []struct {
		name     string
		songs    []*entities.Video
		expected []string
	}{
		{
			name:     "Success",
			songs:    []*entities.Video{&video},
			expected: []string{videoID},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := GetVideoIDs(tc.songs)
			assert.Equal(t, tc.expected, result)
		})
	}
}
