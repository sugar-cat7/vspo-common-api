package entities

import "strings"

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

func (p Platform) String() string {
	return string(p)
}

func (p Platform) Upper() string {
	return strings.ToUpper(string(p))
}

func (p Platform) GetPlatformIconURL() string {
	basePath := "https://raw.githubusercontent.com/sugar-cat7/vspo-common-api/main/assets/icon/"
	switch p {
	case YouTube:
		return basePath + "youtube.png"
	case Twitch:
		return basePath + "twitch.png"
	case Twitcasting:
		return basePath + "twitcasting.png"
	case Niconico:
		return basePath + "niconico.png"
	default:
		return ""
	}
}

const (
	YouTube     Platform = "youtube"
	Twitch      Platform = "twitch"
	Twitcasting Platform = "twitcasting"
	Niconico    Platform = "niconico"
	Discord     Platform = "discord"
)

type LiveStatus string

func (l LiveStatus) String() string {
	return string(l)
}

const (
	LiveStatusUpcoming LiveStatus = "upcoming"
	LiveStatusLive     LiveStatus = "live"
	LiveStatusArchived LiveStatus = "archived"
)

type LiveLink string

func (l LiveLink) String() string {
	return string(l)
}

const (
	LiveLinkYouTube     LiveLink = "https://www.youtube.com/watch?v="
	LiveLinkTwitch      LiveLink = "https://www.twitch.tv/"
	LiveLinkTwitcasting LiveLink = "https://twitcasting.tv/"
	LiveLinkNiconico    LiveLink = "https://live.nicovideo.jp/watch/"
)
