package entities

import (
	"fmt"
	"strings"
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
