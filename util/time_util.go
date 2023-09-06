package util

import (
	"errors"
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

func ParseTimeForCountry(formattedTime, countryCode string) (time.Time, error) {
	// Convert the country code to uppercase for consistency in lookup
	upperCountryCode := strings.ToUpper(countryCode)

	// Get timezone for country code
	tz, exists := CountryCodeToTimeZone[upperCountryCode]
	if !exists {
		tz = "UTC"
	}

	// Load the location for the time zone
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return time.Time{}, err
	}

	// Choose format based on language and parse
	var t time.Time
	switch upperCountryCode {
	case "JP":
		t, err = time.ParseInLocation("2006年01月02日 15時04分", formattedTime, loc)
	case "US", "GB":
		t, err = time.ParseInLocation("January 2, 2006, 3:04 PM", formattedTime, loc)
	default:
		t, err = time.ParseInLocation("2006-01-02 15:04:05", formattedTime, loc)
	}

	if err != nil {
		return time.Time{}, errors.New("error parsing time")
	}

	return t, nil
}

func IsFuture(t time.Time, countryCode string) (bool, error) {

	// Get timezone for country code
	tz, exists := CountryCodeToTimeZone[strings.ToUpper(countryCode)]
	if !exists {
		tz = "UTC"
	}

	// Load the location for the time zone
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return false, err
	}

	// Get current time in the given time zone
	currentTime := time.Now().In(loc)

	// Return true if the time is in the future, otherwise false
	return t.After(currentTime), nil
}
