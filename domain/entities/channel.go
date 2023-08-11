package entities

import (
	"time"

	"cloud.google.com/go/firestore"
)

// Channel represents a YouTube channel.
type Channel struct {
	ID         string            `firestore:"id"`
	Snippet    ChannelSnippet    `firestore:"snippet"`
	Statistics ChannelStatistics `firestore:"statistics"`
}

// ChannelSnippet represents a YouTube channel snippet.
type ChannelSnippet struct {
	Title       string     `firestore:"title"`
	Description string     `firestore:"description"`
	CustomURL   string     `firestore:"customUrl"`
	PublishedAt time.Time  `firestore:"publishedAt"`
	Thumbnails  Thumbnails `firestore:"thumbnails"`
}

// ChannelStatistics represents a YouTube channel statistics.
type ChannelStatistics struct {
	ViewCount             string `firestore:"viewCount"`
	SubscriberCount       string `firestore:"subscriberCount"`
	HiddenSubscriberCount bool   `firestore:"hiddenSubscriberCount"`
	VideoCount            string `firestore:"videoCount"`
}

// GetID returns the ID of the channel.
func (c Channel) GetID() string {
	return c.ID
}

// GetUpdate returns the update of the channel.
func (c Channel) GetUpdate() []firestore.Update {
	return []firestore.Update{
		{Path: "snippet", Value: c.Snippet},
		{Path: "statistics", Value: c.Statistics},
	}
}
