package util

import (
	"strings"
	"time"
)

// CountryCodeToTimeZone provides a basic mapping from country codes to time zones.
var CountryCodeToTimeZone = map[string]string{
	"US": "America/New_York", // You might adjust to more specific ones like "America/Los_Angeles" for PT
	"GB": "Europe/London",
	"JP": "Asia/Tokyo",
	// Add more countries and their main time zones as needed.
}

func FormatTimeForCountry(videoTime time.Time, countryCode string) (string, error) {
	// Convert the country code to uppercase for consistency in lookup
	upperCountryCode := strings.ToUpper(countryCode)

	// Get timezone for country code
	tz, exists := CountryCodeToTimeZone[upperCountryCode]
	if !exists {
		// Default timezone if country code is not recognized
		tz = "UTC"
	}

	// Convert ScheduledStartTime to the timezone
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return "", err
	}

	localTime := videoTime.In(loc)

	// Choose format based on language
	var formattedTime string
	switch upperCountryCode {
	case "JP":
		formattedTime = localTime.Format("2006年01月02日 15時04分")
	case "US", "GB":
		formattedTime = localTime.Format("January 2, 2006, 3:04 PM")
	default:
		formattedTime = localTime.Format("2006-01-02 15:04:05")
	}

	return formattedTime, nil
}
