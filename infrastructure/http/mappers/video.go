package mappers

import (
	"time"

	"github.com/sugar-cat7/vspo-common-api/domain/entities"
)

type ThumbnailResponse struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type ThumbnailsResponse struct {
	Default  ThumbnailResponse `json:"default"`
	Medium   ThumbnailResponse `json:"medium"`
	High     ThumbnailResponse `json:"high"`
	Standard ThumbnailResponse `json:"standard"`
	Maxres   ThumbnailResponse `json:"maxres"`
}

type ViewsResponse struct {
	Daily   string `json:"daily"`
	Weekly  string `json:"weekly"`
	Monthly string `json:"monthly"`
	Total   string `json:"total"`
}

type VideoResponse struct {
	ID           string             `json:"id"`
	Title        string             `json:"title"`
	Description  string             `json:"description"`
	ViewCount    ViewsResponse      `json:"viewCount"`
	PublishedAt  time.Time          `json:"publishedAt"`
	Thumbnails   ThumbnailsResponse `json:"thumbnails"`
	ChannelTitle string             `json:"channelTitle"`
	ChannelID    string             `json:"channelId"`
	Tags         []string           `json:"tags"`
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
