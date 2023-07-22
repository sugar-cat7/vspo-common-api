package entities

import "time"

// Thumbnail represents a YouTube video thumbnail.
type Thumbnail struct {
	URL    string `json:"url" firestore:"url"`
	Width  int    `json:"width" firestore:"width"`
	Height int    `json:"height" firestore:"height"`
}

// Thumbnails represents a YouTube video thumbnails.
type Thumbnails struct {
	Default  Thumbnail `json:"default" firestore:"default"`
	Medium   Thumbnail `json:"medium" firestore:"medium"`
	High     Thumbnail `json:"high" firestore:"high"`
	Standard Thumbnail `json:"standard" firestore:"standard"`
	Maxres   Thumbnail `json:"maxres" firestore:"maxres"`
}

// Views represents a YouTube video view count.
type Views struct {
	Daily   string `firestore:"daily" json:"daily"`
	Weekly  string `firestore:"weekly" json:"weekly"`
	Monthly string `firestore:"monthly" json:"monthly"`
	Total   string `firestore:"total" json:"total"`
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
