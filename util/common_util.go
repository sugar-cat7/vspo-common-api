package util

import (
	"errors"
	"fmt"
	"time"

	"github.com/sugar-cat7/vspo-common-api/domain/entities"
)

// Chunk splits a slice into chunks of the given size.
func Chunk[T any](data []T, chunkSize int) ([][]T, error) {
	if chunkSize <= 0 {
		return nil, errors.New("chunkSize must be greater than 0")
	}

	var chunks [][]T
	length := len(data)

	for i := 0; i < length; i += chunkSize {
		end := i + chunkSize

		// If end is more than the length of the data,
		// adjust it to the length of the data.
		if end > length {
			end = length
		}

		chunk := data[i:end]
		chunks = append(chunks, chunk)
	}

	return chunks, nil
}

// ConvertToPtrSlice converts a slice of values to a slice of pointers.
func ConvertToPtrSlice[T any](data []T) []*T {
	var ptrs []*T
	for _, d := range data {
		d := d // Create a new 'd' variable on each loop iteration
		ptrs = append(ptrs, &d)
	}
	return ptrs
}

func GetStartTime(cronType entities.CronType) (string, error) {
	now := time.Now()
	startOfToday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	var startTime time.Time
	switch cronType {
	case entities.Daily:
		startTime = startOfToday
	case entities.Weekly:
		startTime = startOfToday.AddDate(0, 0, -7)
	case entities.Monthly:
		startTime = startOfToday.AddDate(0, -1, 0)
	default:
		return "", fmt.Errorf("Unsupported CronType: %s", cronType)
	}

	return startTime.Format(time.RFC3339), nil
}
