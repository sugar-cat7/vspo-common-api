package entities

import (
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/sugar-cat7/vspo-common-api/util"
)

// Video represents a YouTube video.
type Video struct {
	ID                 string     `firestore:"id"`
	Title              string     `firestore:"title"`
	Description        string     `firestore:"description"`
	ViewCount          Views      `firestore:"viewCount"`
	PublishedAt        time.Time  `firestore:"publishedAt"`
	Thumbnails         Thumbnails `firestore:"thumbnails"`
	ChannelTitle       string     `firestore:"channelTitle"`
	ChannelID          string     `firestore:"channelId"`
	ChannelIcon        string     `firestore:"channelIcon"`
	Platform           Platform   `firestore:"platform"`
	Tags               []string   `firestore:"tags"`
	ScheduledStartTime time.Time  `firestore:"scheduledStartTime"`
	ActualEndTime      time.Time  `firestore:"actualEndTime"`
	LiveStatus         LiveStatus
	Link               string
}

// GetID returns the ID of the video.
func (v *Video) GetID() string {
	return v.ID
}

// GetUpdate returns the update of the video.
func (v Video) GetUpdate() []firestore.Update {
	return []firestore.Update{
		{Path: "title", Value: v.Title},
		{Path: "description", Value: v.Description},
		{Path: "viewCount", Value: v.ViewCount},
		{Path: "publishedAt", Value: v.PublishedAt},
		{Path: "thumbnails", Value: v.Thumbnails},
		{Path: "channelTitle", Value: v.ChannelTitle},
		{Path: "channelId", Value: v.ChannelID},
		{Path: "tags", Value: v.Tags},
	}
}

// GetLiveStatus returns the live status of the video.
func (v *Video) GetLiveStatus() LiveStatus {
	currentTime := time.Now().UTC()

	// FIXME: db内のゼロ値修正後に修正
	year2000 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	// ActualEndTime が 2000年以前であるか、またはゼロの場合を確認します。
	if v.ActualEndTime.Before(year2000) {
		if v.ScheduledStartTime.After(currentTime) {
			return LiveStatusUpcoming
		}
		return LiveStatusLive
	}

	return LiveStatusArchived
}

// GetLink returns the link of the video.
func (v *Video) GetLink() string {
	switch v.Platform {
	case YouTube:
		return LiveLinkYouTube.String() + v.ID
	case Twitch:
		if v.GetLiveStatus() != LiveStatusLive {
			return LiveLinkTwitch.String() + "videos/" + v.ID
		}
		return LiveLinkTwitch.String() + v.ChannelID
	case Twitcasting:
		return v.Link + "movie/" + v.ID
	case Niconico:
		return LiveLinkNiconico.String() + v.ID
	default:
		return ""
	}
}

// Color constants for the video.(used Discord embed color)
const (
	ColorUpcoming = 0x00FF00 // Green color for upcoming
	ColorLive     = 0xFF0000 // Red color for live
	ColorArchived = 0x0000FF // Blue color for archived
	ColorDefault  = 0x808080 // Default grey color if none of the conditions match
)

// GetStatusColor returns the color of the video.
func (v *Video) GetStatusColor() int {
	switch v.GetLiveStatus() {
	case LiveStatusUpcoming:
		return ColorUpcoming
	case LiveStatusLive:
		return ColorLive
	case LiveStatusArchived:
		return ColorArchived
	default:
		return ColorDefault
	}
}

// UpdateViewCount updates the view count of the video by CronType.
func (v *Video) UpdateViewCount(CronType CronType, viewCount string) {
	switch CronType {
	case Daily:
		v.ViewCount.Daily = viewCount
	case Weekly:
		v.ViewCount.Weekly = viewCount
	case Monthly:
		v.ViewCount.Monthly = viewCount
	}
	v.ViewCount.Total = viewCount
}

func (s *Video) FormatThumbnailURL(url string) string {
	url = strings.Replace(url, "http://", "https://", -1)
	switch s.Platform {
	case Twitch:
		url = strings.Replace(url, "%{width}", "400", -1)
		url = strings.Replace(url, "%{height}", "220", -1)
		url = strings.Replace(url, "-{width}x{height}", "-400x220", -1)
		return url
	default:
		return url
	}
}

// Videos represents a list of videos.
type Videos []*Video

// GetIDs returns the IDs of the videos.
func (v Videos) GetIDs() []string {
	ids := make([]string, len(v))
	for i, video := range v {
		ids[i] = video.ID
	}
	return ids
}

func (v Videos) SetLocalTime(countryCode string) error {
	for _, video := range v {
		scheduledTime, err := util.ConvertTimeToCountryTimeZone(video.ScheduledStartTime, countryCode)
		if err != nil {
			return err
		}
		video.ScheduledStartTime = scheduledTime
		actualTime, err := util.ConvertTimeToCountryTimeZone(video.ActualEndTime, countryCode)
		if err != nil {
			return err
		}
		video.ActualEndTime = actualTime
	}
	return nil
}
