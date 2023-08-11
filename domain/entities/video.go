package entities

import (
	"time"

	"cloud.google.com/go/firestore"
)

// Video represents a YouTube video.
type Video struct {
	ID           string     `firestore:"id"`
	Title        string     `firestore:"title"`
	Description  string     `firestore:"description"`
	ViewCount    Views      `firestore:"viewCount"`
	PublishedAt  time.Time  `firestore:"publishedAt"`
	Thumbnails   Thumbnails `firestore:"thumbnails"`
	ChannelTitle string     `firestore:"channelTitle"`
	ChannelID    string     `firestore:"channelId"`
	Tags         []string   `firestore:"tags"`
}

// GetID returns the ID of the video.
func (s Video) GetID() string {
	return s.ID
}

// GetUpdate returns the update of the video.
func (s Video) GetUpdate() []firestore.Update {
	return []firestore.Update{
		{Path: "title", Value: s.Title},
		{Path: "description", Value: s.Description},
		{Path: "viewCount", Value: s.ViewCount},
		{Path: "publishedAt", Value: s.PublishedAt},
		{Path: "thumbnails", Value: s.Thumbnails},
		{Path: "channelTitle", Value: s.ChannelTitle},
		{Path: "channelId", Value: s.ChannelID},
		{Path: "tags", Value: s.Tags},
	}
}
