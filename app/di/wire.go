//go:build wireinject

//go:generate go run github.com/google/wire/cmd/wire@v0.5.0
package di

import (
	"github.com/google/wire"
	"github.com/sugar-cat7/vspo-common-api/domain/services"
	"github.com/sugar-cat7/vspo-common-api/infrastructure/firestore"
	"github.com/sugar-cat7/vspo-common-api/infrastructure/http/handlers"
	channel_handlers "github.com/sugar-cat7/vspo-common-api/infrastructure/http/handlers/channel"
	clip_handlers "github.com/sugar-cat7/vspo-common-api/infrastructure/http/handlers/clip"
	song_handlers "github.com/sugar-cat7/vspo-common-api/infrastructure/http/handlers/song"
	"github.com/sugar-cat7/vspo-common-api/usecases"
)

// Application is the main application struct which holds all the dependencies together.
type Application struct {
	GetAllSongsHandler               *song_handlers.GetAllSongsHandler
	CreateSongHandler                *song_handlers.CreateSongHandler
	UpdateSongsHandler               *song_handlers.UpdateSongsHandler
	GetChannelsHandler               *channel_handlers.GetChannelsHandler
	CreateChannelHandler             *channel_handlers.CreateChannelHandler
	UpdateChannelsFromYoutubeHandler *channel_handlers.UpdateChannelsFromYoutubeHandler
	GetClipsByPeriodHandler          *clip_handlers.GetClipsByPeriodHandler
	UpdateClipsHandler               *clip_handlers.UpdateClipsHandler
}

// NewApplication creates a new Application.
func NewApplication(getAllSongsHandler *song_handlers.GetAllSongsHandler, createSongHandler *song_handlers.CreateSongHandler, updateSongsHandler *song_handlers.UpdateSongsHandler,
	getChannelsHandler *channel_handlers.GetChannelsHandler, createChannelHandler *channel_handlers.CreateChannelHandler, updateChannelsFromYoutubeHandler *channel_handlers.UpdateChannelsFromYoutubeHandler,
	getClipsByPeriodHandler *clip_handlers.GetClipsByPeriodHandler, UpdateClipsHandler *clip_handlers.UpdateClipsHandler,
) *Application {
	return &Application{
		GetAllSongsHandler:               getAllSongsHandler,
		CreateSongHandler:                createSongHandler,
		UpdateSongsHandler:               updateSongsHandler,
		GetChannelsHandler:               getChannelsHandler,
		CreateChannelHandler:             createChannelHandler,
		UpdateChannelsFromYoutubeHandler: updateChannelsFromYoutubeHandler,
		GetClipsByPeriodHandler:          getClipsByPeriodHandler,
		UpdateClipsHandler:               UpdateClipsHandler,
	}
}

// InitializeApplication initializes the entire application with all its dependencies using wire.
func InitializeApplication() (*Application, func(), error) {
	panic(wire.Build(services.Set, firestore.Set, handlers.Set, NewApplication, usecases.Set))
}
