package entities

// YTVideoListResponse represents a YouTube video list response.
type YTVideoListResponse struct {
	Kind  string    `json:"kind"`
	Etag  string    `json:"etag"`
	Items []YTVideo `json:"items"`
}

// YTVideo represents a YouTube video.
type YTVideo struct {
	Kind       string         `json:"kind"`
	Etag       string         `json:"etag"`
	ID         string         `json:"id"`
	Snippet    YTVideoSnippet `json:"snippet"`
	Statistics YTStatistics   `json:"statistics"`
}

// YTVideoSnippet represents a YouTube video snippet.
type YTVideoSnippet struct {
	PublishedAt  string       `json:"publishedAt"`
	ChannelID    string       `json:"channelId"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	Thumbnails   YTThumbnails `json:"thumbnails"`
	ChannelTitle string       `json:"channelTitle"`
	Tags         []string     `json:"tags"`
	CategoryID   string       `json:"categoryId"`
}

// YTThumbnails represents a YouTube video thumbnails.
type YTThumbnails struct {
	Default  YTThumbnail `json:"default"`
	Medium   YTThumbnail `json:"medium"`
	High     YTThumbnail `json:"high"`
	Standard YTThumbnail `json:"standard"`
	Maxres   YTThumbnail `json:"maxres"`
}

// YTThumbnail represents a YouTube video thumbnail.
type YTThumbnail struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// YTStatistics represents a YouTube video statistics.
type YTStatistics struct {
	ViewCount     string `json:"viewCount"`
	LikeCount     string `json:"likeCount"`
	FavoriteCount string `json:"favoriteCount"`
	CommentCount  string `json:"commentCount"`
}

// YTYouTubePlaylistResponse represents a YouTube playlist response.
type YTYouTubePlaylistResponse struct {
	Kind  string `json:"kind"`
	Items []struct {
		YTPlayListSnippet `json:"snippet"`
	} `json:"items"`
	PageInfo YTPageInfo `json:"pageInfo"`
}

// YTPlayListSnippet represents a YouTube playlist snippet.
type YTPlayListSnippet struct {
	PublishedAt            string       `json:"publishedAt"`
	ChannelID              string       `json:"channelId"`
	Title                  string       `json:"title"`
	Description            string       `json:"description"`
	Thumbnails             YTThumbnails `json:"thumbnails"`
	ChannelTitle           string       `json:"channelTitle"`
	PlaylistID             string       `json:"playlistId"`
	Position               int          `json:"position"`
	ResourceID             YTResourceID `json:"resourceId"`
	VideoOwnerChannelTitle string       `json:"videoOwnerChannelTitle"`
	VideoOwnerChannelID    string       `json:"videoOwnerChannelId"`
}

// YTResourceID represents a YouTube resource ID.
type YTResourceID struct {
	Kind    string `json:"kind"`
	VideoID string `json:"videoId"`
}

// YTPageInfo represents a YouTube page info.
type YTPageInfo struct {
	TotalResults   int `json:"totalResults"`
	ResultsPerPage int `json:"resultsPerPage"`
}

// YTChannel represents a YouTube channel.
type YTChannel struct {
	ID         string              `json:"id"`
	Snippet    YTChannelSnippet    `json:"snippet"`
	Statistics YTChannelStatistics `json:"statistics"`
}

// YTChannelSnippet represents a YouTube channel snippet.
type YTChannelSnippet struct {
	PublishedAt  string       `json:"publishedAt"`
	ChannelID    string       `json:"channelId"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	Thumbnails   YTThumbnails `json:"thumbnails"`
	ChannelTitle string       `json:"channelTitle"`
	CustomURL    string       `json:"customUrl"`
}

// YTChannelStatistics represents a YouTube channel statistics.
type YTChannelStatistics struct {
	ViewCount             string `json:"viewCount"`
	SubscriberCount       string `json:"subscriberCount"`
	HiddenSubscriberCount bool   `json:"hiddenSubscriberCount"`
	VideoCount            string `json:"videoCount"`
}
