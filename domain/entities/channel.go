package entities

import (
	"cloud.google.com/go/firestore"
)

// Channel represents a YouTube channel.
type Channel struct {
	ID         string            `json:"id" firestore:"id"`
	Snippet    ChannelSnippet    `json:"snippet" firestore:"snippet"`
	Statistics ChannelStatistics `json:"statistics" firestore:"statistics"`
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
