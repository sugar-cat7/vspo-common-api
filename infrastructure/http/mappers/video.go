package mappers

import (
	"time"

	"github.com/sugar-cat7/vspo-common-api/domain/entities"
)

// ThumbnailResponse is the response structure for a thumbnail.
type ThumbnailResponse struct {
	URL    string `json:"url" example:"https://i.ytimg.com/vi/Qh6aSTTkmEs/default.jpg"`
	Width  int    `json:"width" example:"120"`
	Height int    `json:"height" example:"90"`
}

// ThumbnailsResponse contains multiple thumbnails.
type ThumbnailsResponse struct {
	Default  ThumbnailResponse `json:"default"`
	Medium   ThumbnailResponse `json:"medium"`
	High     ThumbnailResponse `json:"high"`
	Standard ThumbnailResponse `json:"standard"`
	Maxres   ThumbnailResponse `json:"maxres"`
}

type ViewsResponse struct {
	Daily   string `json:"daily" example:"1000"`
	Weekly  string `json:"weekly" example:"10000"`
	Monthly string `json:"monthly" example:"100000"`
	Total   string `json:"total" example:"1000000"`
}

type VideoResponse struct {
	ID           string             `json:"id" example:"Qh6aSTTkmEs"`
	Title        string             `json:"title" example:"【ぶいすぽっ！】Blessing ~12人で歌ってみた~"`
	Description  string             `json:"description" example:""`
	ViewCount    ViewsResponse      `json:"viewCount"`
	PublishedAt  time.Time          `json:"publishedAt" example:"2020-12-31T12:34:56+09:00"`
	Thumbnails   ThumbnailsResponse `json:"thumbnails"`
	ChannelTitle string             `json:"channelTitle" example:"花芽なずな / Nazuna Kaga"`
	ChannelID    string             `json:"channelId" example:"UCiMG6VdScBabPhJ1ZtaVmbw"`
	Tags         []string           `json:"tags" example:"[ぶいすぽっ！, 歌ってみた]"`
}

// VideosResponse Clip, Song, Live...結局形式は同じなのでresponseとしてはまとめる
type VideosResponse struct {
	Videos []VideoResponse `json:"videos"`
}

func MapThumbnailToResponse(thumbnail entities.Thumbnail) ThumbnailResponse {
	return ThumbnailResponse{
		URL:    thumbnail.URL,
		Width:  thumbnail.Width,
		Height: thumbnail.Height,
	}
}

func MapThumbnailsToResponse(thumbnails entities.Thumbnails) ThumbnailsResponse {
	return ThumbnailsResponse{
		Default:  MapThumbnailToResponse(thumbnails.Default),
		Medium:   MapThumbnailToResponse(thumbnails.Medium),
		High:     MapThumbnailToResponse(thumbnails.High),
		Standard: MapThumbnailToResponse(thumbnails.Standard),
		Maxres:   MapThumbnailToResponse(thumbnails.Maxres),
	}
}

func MapViewsToResponse(views entities.Views) ViewsResponse {
	return ViewsResponse{
		Daily:   views.Daily,
		Weekly:  views.Weekly,
		Monthly: views.Monthly,
		Total:   views.Total,
	}
}

func MapVideoToResponse(video *entities.Video) VideoResponse {
	return VideoResponse{
		ID:           video.ID,
		Title:        video.Title,
		Description:  video.Description,
		ViewCount:    MapViewsToResponse(video.ViewCount),
		PublishedAt:  video.PublishedAt,
		Thumbnails:   MapThumbnailsToResponse(video.Thumbnails),
		ChannelTitle: video.ChannelTitle,
		ChannelID:    video.ChannelID,
		Tags:         video.Tags,
	}
}

func MapVideosToResponse(videos []*entities.Video) VideosResponse {
	var videoResponses []VideoResponse
	for _, video := range videos {
		videoResponses = append(videoResponses, MapVideoToResponse(video))
	}
	return VideosResponse{Videos: videoResponses}
}
