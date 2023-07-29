//go:generate go run github.com/google/wire/cmd/wire@v0.5.0
//go:build wireinject

package di

import (
	"github.com/google/wire"
	channel_handlers "github.com/sugar-cat7/vspo-common-api/app/http/handlers/channels"
	song_handlers "github.com/sugar-cat7/vspo-common-api/app/http/handlers/songs"
	"github.com/sugar-cat7/vspo-common-api/domain/services"
	"github.com/sugar-cat7/vspo-common-api/infrastructure/firestore"
	"github.com/sugar-cat7/vspo-common-api/infrastructure/youtube"
	channel_usecases "github.com/sugar-cat7/vspo-common-api/usecases/channel"
	song_usecases "github.com/sugar-cat7/vspo-common-api/usecases/song"
)

// Application is the main application struct which holds all the dependencies together.
type Application struct {
	GetAllSongsHandler               *song_handlers.GetAllSongsHandler
	CreateSongHandler                *song_handlers.CreateSongHandler
	UpdateSongsHandler               *song_handlers.UpdateSongsHandler
	GetChannelsHandler               *channel_handlers.GetChannelsHandler
	CreateChannelHandler             *channel_handlers.CreateChannelHandler
	UpdateChannelsFromYoutubeHandler *channel_handlers.UpdateChannelsFromYoutubeHandler
}

// NewApplication creates a new Application.
func NewApplication(getAllSongsHandler *song_handlers.GetAllSongsHandler, createSongHandler *song_handlers.CreateSongHandler, updateSongsHandler *song_handlers.UpdateSongsHandler,
	getChannelsHandler *channel_handlers.GetChannelsHandler, createChannelHandler *channel_handlers.CreateChannelHandler, updateChannelsFromYoutubeHandler *channel_handlers.UpdateChannelsFromYoutubeHandler,
) *Application {
	return &Application{
		GetAllSongsHandler:               getAllSongsHandler,
		CreateSongHandler:                createSongHandler,
		UpdateSongsHandler:               updateSongsHandler,
		GetChannelsHandler:               getChannelsHandler,
		CreateChannelHandler:             createChannelHandler,
		UpdateChannelsFromYoutubeHandler: updateChannelsFromYoutubeHandler,
	}
}

// InitializeApplication initializes the entire application with all its dependencies using wire.
func InitializeApplication() (*Application, func(), error) {
	panic(wire.Build(services.Set, firestore.Set, youtube.Set, song_usecases.Set, channel_usecases.Set, song_handlers.Set, channel_handlers.Set, NewApplication))
}
