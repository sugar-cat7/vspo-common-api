package entities

import (
	"time"

	"cloud.google.com/go/firestore"
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
func (v Video) GetID() string {
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
		return LiveLinkTwitcasting.String() + v.ID
	case Niconico:
		return LiveLinkNiconico.String() + v.ID
	default:
		return ""
	}
}

// 定数として色を定義
const (
	ColorUpcoming = 0x00FF00 // Green color for upcoming
	ColorLive     = 0xFF0000 // Red color for live
	ColorArchived = 0x0000FF // Blue color for archived
	ColorDefault  = 0x808080 // Default grey color if none of the conditions match
)

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
