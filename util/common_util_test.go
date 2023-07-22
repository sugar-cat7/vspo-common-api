package util

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunk(t *testing.T) {
	type testCase struct {
		name      string
		data      []int
		chunkSize int
		expected  [][]int
		wantErr   bool
	}

	testCases := []testCase{
		{
			name:      "chunkSize 3",
			data:      []int{1, 2, 3, 4, 5, 6, 7},
			chunkSize: 3,
			expected:  [][]int{{1, 2, 3}, {4, 5, 6}, {7}},
			wantErr:   false,
		},
		{
			name:      "chunkSize 1",
			data:      []int{1, 2, 3, 4, 5},
			chunkSize: 1,
			expected:  [][]int{{1}, {2}, {3}, {4}, {5}},
			wantErr:   false,
		},
		{
			name:      "chunkSize greater than len(data)",
			data:      []int{1, 2, 3, 4, 5},
			chunkSize: 10,
			expected:  [][]int{{1, 2, 3, 4, 5}},
			wantErr:   false,
		},
		{
			name:      "chunkSize 0",
			data:      []int{1, 2, 3, 4, 5},
			chunkSize: 0,
			expected:  nil,
			wantErr:   true,
		},
		// Add more cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Chunk(tc.data, tc.chunkSize)

			if (err != nil) != tc.wantErr {
				t.Errorf("Chunk() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Chunk() = %v, want %v", result, tc.expected)
			}
		})
	}
}

func TestConvertToPtrSlice(t *testing.T) {
	type testCase struct {
		name     string
		data     []int
		expected []int // Use an int slice to store the expected values.
	}

	testCases := []testCase{
		{
			name:     "case 1",
			data:     []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "case 2",
			data:     []int{10, 20, 30},
			expected: []int{10, 20, 30},
		},
		// Add more cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ConvertToPtrSlice(tc.data)

			// Check the length of the result slice
			assert.Equal(t, len(tc.expected), len(result), "they should have the same length")

			// Check each element of the result slice
			for i := 0; i < len(result); i++ {
				assert.Equal(t, tc.expected[i], *result[i], "the elements should be equal")
			}
		})
	}
}
