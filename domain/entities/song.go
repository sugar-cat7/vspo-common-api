package entities

import (
	"time"

	"cloud.google.com/go/firestore"
)

// Song represents a YouTube video.
type Song struct {
	ID           string     `firestore:"id" json:"id"`
	Title        string     `firestore:"title" json:"title"`
	Description  string     `firestore:"description" json:"description"`
	ViewCount    Views      `firestore:"viewCount" json:"viewCount"`
	PublishedAt  time.Time  `firestore:"publishedAt" json:"publishedAt"`
	Thumbnails   Thumbnails `firestore:"thumbnails" json:"thumbnails"`
	ChannelTitle string     `firestore:"channelTitle" json:"channelTitle"`
	ChannelID    string     `firestore:"channelId" json:"channelId"`
	Tags         []string   `firestore:"tags" json:"tags"`
}

// GetID returns the ID of the song.
func (s Song) GetID() string {
	return s.ID
}

// GetUpdate returns the update of the song.
func (s Song) GetUpdate() []firestore.Update {
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
