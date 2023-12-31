package mappers

import (
	entities "github.com/sugar-cat7/vspo-common-api/domain/entities"
	entities2 "github.com/sugar-cat7/vspo-common-api/domain/entities/legacy"
)

// LiveStreamMap maps a LiveStream to a domain Video.
func LiveStreamMap(liveStream *entities2.OldVideo) (*entities.Video, error) {
	v := &entities.Video{
		ID:          liveStream.ID,
		Title:       liveStream.Title,
		Description: liveStream.Description,
		ViewCount: entities.Views{
			Daily:   liveStream.NewViewCount.Daily,
			Weekly:  liveStream.NewViewCount.Weekly,
			Monthly: liveStream.NewViewCount.Monthly,
			Total:   liveStream.NewViewCount.Total,
		},
		PublishedAt:        liveStream.CreatedAt,
		ChannelTitle:       liveStream.ChannelTitle,
		ChannelID:          liveStream.ChannelID,
		ChannelIcon:        liveStream.IconURL,
		Platform:           liveStream.Platform,
		ScheduledStartTime: liveStream.ScheduledStartTime,
		ActualEndTime:      liveStream.ActualEndTime,
	}

	if liveStream.Platform == entities.Twitch {
		v.ChannelID = liveStream.TwitchName
		v.ID = liveStream.TwitchPastVideoId
	}
	if liveStream.Platform == entities.Twitcasting {
		v.Link = liveStream.TwitcastingLink
	}
	v.LiveStatus = v.GetLiveStatus()
	v.Link = v.GetLink()
	v.Thumbnails = entities.Thumbnails{
		Default: entities.Thumbnail{
			URL: v.FormatThumbnailURL(liveStream.ThumbnailURL),
		},
		Medium: entities.Thumbnail{
			URL: v.FormatThumbnailURL(liveStream.ThumbnailURL),
		},
	}

	return v, nil
}

// LiveStreamMapMultiple maps multiple LiveStreams to domain Videos.
func LiveStreamMapMultiple(liveStreams entities2.OldVideos) (entities.Videos, error) {
	videos := make(entities.Videos, len(liveStreams))
	for i, liveStream := range liveStreams {
		video, err := LiveStreamMap(liveStream)
		if err != nil {
			return nil, err
		}
		videos[i] = video
	}

	return videos, nil
}
