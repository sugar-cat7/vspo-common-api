package usecases

import (
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/domain/ports"
	"github.com/sugar-cat7/vspo-common-api/domain/repositories"
	"github.com/sugar-cat7/vspo-common-api/usecases/mappers"
)

type DiscordSendMessage struct {
	liveStreamRepository repositories.LiveStreamRepository
	discordService       ports.DiscordService
}

func NewDiscordSendMessage(discordService ports.DiscordService, liveStreamRepository repositories.LiveStreamRepository) *DiscordSendMessage {
	return &DiscordSendMessage{
		liveStreamRepository: liveStreamRepository,
		discordService:       discordService,
	}
}

func (c *DiscordSendMessage) Execute(start, end, countryCode string) (entities.Videos, error) {

	// Get all liveStreams from Firestore
	liveStreams, err := c.liveStreamRepository.FindAllByPeriod(start, end)
	if err != nil {
		return nil, err
	}

	videos, err := mappers.LiveStreamMapMultiple(liveStreams)
	if err != nil {
		return nil, err
	}
	sendVideos := make(entities.Videos, 0, len(videos))
	for _, v := range videos {
		if v.GetLiveStatus() != entities.LiveStatusArchived {
			sendVideos = append(sendVideos, v)
		}
	}

	if len(sendVideos) == 0 {
		err = c.discordService.DeleteMessages()
		if err != nil {
			return nil, err
		}
		return nil, nil
	}

	if sendVideos.SetLocalTime(countryCode) != nil {
		return nil, err
	}

	err = c.discordService.SendMessages(sendVideos, countryCode)
	if err != nil {
		return nil, err
	}

	return sendVideos, nil
}
