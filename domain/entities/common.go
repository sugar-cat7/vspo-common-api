package entities

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

type Platform string

const (
	YouTube     Platform = "youtube"
	Twitch      Platform = "twitch"
	Twitcasting Platform = "twitcasting"
)
