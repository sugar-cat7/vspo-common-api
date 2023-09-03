//go:generate mockgen -destination=../../mocks/ports/mock_youtube_port.go -package=mocks github.com/sugar-cat7/vspo-common-api/domain/ports DiscordService
package ports

import (
	"github.com/sugar-cat7/vspo-common-api/domain/entities"
)

// DiscordService is an interface for a Discord implementation of a song service.
type DiscordService interface {
	SendMessages(liveStreams []*entities.Video, countryCode string) error
}
