package entities

import (
	"time"

	"cloud.google.com/go/firestore"
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
)

// FIXME: OldVideo is a legacy entity.
type OldVideo struct {
	ID           string            `firestore:"id"`
	Title        string            `firestore:"title"`
	Description  string            `firestore:"description"`
	ChannelID    string            `firestore:"channelId"`
	ChannelTitle string            `firestore:"channelTitle"`
	ThumbnailURL string            `firestore:"thumbnailUrl"`
	IconURL      string            `firestore:"iconUrl"`
	Platform     entities.Platform `firestore:"platform"`
	ViewCount    string            `firestore:"viewCount"`
	// LikeCount    string            `firestore:"likeCount"`
	CommentCount       string         `firestore:"commentCount"`
	NewViewCount       entities.Views `firestore:"newViewCount"`
	CreatedAt          time.Time      `firestore:"createdAt"`
	ScheduledStartTime time.Time      `firestore:"scheduledStartTime"`
	ActualEndTime      time.Time      `firestore:"actualEndTime"`
	TwitchName         string         `firestore:"twitchName"`
	TwitchPastVideoId  string         `firestore:"twitchPastVideoId"`
}

// GetUpdate returns the update of the OldVideo.
func (s OldVideo) GetUpdate() []firestore.Update {
	return []firestore.Update{
		{Path: "title", Value: s.Title},
		{Path: "description", Value: s.Description},
		{Path: "channelTitle", Value: s.ChannelTitle},
		{Path: "thumbnailUrl", Value: s.ThumbnailURL},
		{Path: "iconUrl", Value: s.IconURL},
		{Path: "viewCount", Value: s.ViewCount},
		// {Path: "likeCount", Value: s.LikeCount},
		{Path: "commentCount", Value: s.CommentCount},
		{Path: "newViewCount", Value: s.NewViewCount},
	}
}
