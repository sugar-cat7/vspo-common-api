package entities

import (
	"fmt"
	"strings"
	"time"
)

// CronType represents a cron job type.
type CronType string

// Cron represents a cron job.
const (
	Daily   CronType = "daily"
	Weekly  CronType = "weekly"
	Monthly CronType = "monthly"
	None    CronType = "none"
)

// ParseCronType parses a string into a CronType.
func ParseCronType(s string) (CronType, error) {
	switch strings.ToLower(s) {
	case string(Daily):
		return Daily, nil
	case string(Weekly):
		return Weekly, nil
	case string(Monthly):
		return Monthly, nil
	case string(None):
		return None, nil
	default:
		return "", fmt.Errorf("invalid CronType: %s", s)
	}
}

func GetStartTime(cronType CronType) (string, error) {
	now := time.Now()
	startOfToday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	var startTime time.Time
	switch cronType {
	case Daily:
		startTime = startOfToday
	case Weekly:
		startTime = startOfToday.AddDate(0, 0, -7)
	case Monthly:
		startTime = startOfToday.AddDate(0, -1, 0)
	default:
		return "", fmt.Errorf("Unsupported CronType: %s", cronType)
	}

	return startTime.Format(time.RFC3339), nil
}
