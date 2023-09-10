package util

import (
	"errors"
	"strings"
	"time"
)

// CountryCodeToTimeZone provides a basic mapping from country codes to time zones.
var CountryCodeToTimeZone = map[string]string{
	"US": "America/New_York",
	"GB": "Europe/London",
	"JP": "Asia/Tokyo",
}

// ConvertTimeToCountryTimeZone converts the given time to the specified country's timezone.
func ConvertTimeToCountryTimeZone(t time.Time, countryCode string) (time.Time, error) {
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

	// Convert the given time to the specified time zone and return
	return t.In(loc), nil
}

func FormatTimeForCountry(videoTime time.Time, countryCode string) (string, error) {
	localTime, err := ConvertTimeToCountryTimeZone(videoTime, countryCode)
	if err != nil {
		return "", err
	}

	// Choose format based on language
	upperCountryCode := strings.ToUpper(countryCode)
	switch upperCountryCode {
	case "JP":
		return localTime.Format("2006年01月02日 15時04分"), nil
	case "US", "GB":
		return localTime.Format("January 2, 2006, 3:04 PM"), nil
	default:
		return localTime.Format("2006-01-02 15:04:05"), nil
	}
}

func ParseTimeForCountry(formattedTime, countryCode string) (time.Time, error) {
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
	localNow, err := ConvertTimeToCountryTimeZone(time.Now(), countryCode)
	if err != nil {
		return false, err
	}

	return t.After(localNow), nil
}
