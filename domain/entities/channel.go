package entities

import (
	"time"

	"cloud.google.com/go/firestore"
)

// Channel represents a YouTube channel.
type Channel struct {
	ID         string            `json:"id" firestore:"id"`
	Snippet    ChannelSnippet    `json:"snippet" firestore:"snippet"`
	Statistics ChannelStatistics `json:"statistics" firestore:"statistics"`
}

// ChannelSnippet represents a YouTube channel snippet.
type ChannelSnippet struct {
	Title       string     `json:"title" firestore:"title"`
	Description string     `json:"description" firestore:"description"`
	CustomURL   string     `json:"customUrl" firestore:"customUrl"`
	PublishedAt time.Time  `json:"publishedAt" firestore:"publishedAt"`
	Thumbnails  Thumbnails `json:"thumbnails" firestore:"thumbnails"`
}

// ChannelStatistics represents a YouTube channel statistics.
type ChannelStatistics struct {
	ViewCount             string `json:"viewCount" firestore:"viewCount"`
	SubscriberCount       string `json:"subscriberCount" firestore:"subscriberCount"`
	HiddenSubscriberCount bool   `json:"hiddenSubscriberCount" firestore:"hiddenSubscriberCount"`
	VideoCount            string `json:"videoCount" firestore:"videoCount"`
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
