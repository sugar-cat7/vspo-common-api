// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	ports2 "github.com/sugar-cat7/vspo-common-api/infrastructure/api/discord"
	"github.com/sugar-cat7/vspo-common-api/infrastructure/api/youtube"
	"github.com/sugar-cat7/vspo-common-api/infrastructure/firestore"
	handlers2 "github.com/sugar-cat7/vspo-common-api/infrastructure/http/handlers/channel"
	handlers3 "github.com/sugar-cat7/vspo-common-api/infrastructure/http/handlers/clip"
	handlers4 "github.com/sugar-cat7/vspo-common-api/infrastructure/http/handlers/cron"
	handlers6 "github.com/sugar-cat7/vspo-common-api/infrastructure/http/handlers/discord"
	handlers5 "github.com/sugar-cat7/vspo-common-api/infrastructure/http/handlers/livestream"
	"github.com/sugar-cat7/vspo-common-api/infrastructure/http/handlers/song"
	usecases2 "github.com/sugar-cat7/vspo-common-api/usecases/channel"
	usecases3 "github.com/sugar-cat7/vspo-common-api/usecases/clip"
	usecases5 "github.com/sugar-cat7/vspo-common-api/usecases/discord"
	usecases4 "github.com/sugar-cat7/vspo-common-api/usecases/livestream"
	"github.com/sugar-cat7/vspo-common-api/usecases/song"
)

// Injectors from wire.go:

// InitializeApplication initializes the entire application with all its dependencies using wire.
func InitializeApplication() (*Application, func(), error) {
	firestoreClient, err := firestore.ProvideFirestoreClient()
	if err != nil {
		return nil, nil, err
	}
	songRepository := firestore.NewSongRepository(firestoreClient)
	repositoriesSongRepository := firestore.ProvideSongRepository(firestoreClient, songRepository)
	getAllSongs := usecases.NewGetAllSongs(repositoriesSongRepository)
	getAllSongsHandler := handlers.NewGetAllSongsHandler(getAllSongs)
	youTubeService, err := ports.NewYouTubeService()
	if err != nil {
		return nil, nil, err
	}
	createSong := usecases.NewCreateSong(youTubeService, repositoriesSongRepository)
	createSongHandler := handlers.NewCreateSongHandler(createSong)
	updateSongs := usecases.NewUpdateSongs(youTubeService, repositoriesSongRepository)
	updateSongsHandler := handlers.NewUpdateSongsHandler(updateSongs)
	addNewSong := usecases.NewAddNewSong(youTubeService, repositoriesSongRepository)
	addNewSongHandler := handlers.NewAddNewSongHandler(addNewSong)
	channelRepository := firestore.NewChannelRepository(firestoreClient)
	repositoriesChannelRepository := firestore.ProvideChannelRepository(firestoreClient, channelRepository)
	getChannels := usecases2.NewGetChannels(repositoriesChannelRepository)
	getChannelsHandler := handlers2.NewGetChannelsHandler(getChannels)
	createChannel := usecases2.NewCreateChannel(youTubeService, repositoriesChannelRepository)
	createChannelHandler := handlers2.NewCreateChannelHandler(createChannel)
	updateChannelsFromYoutube := usecases2.NewUpdateChannelsFromYoutube(youTubeService, repositoriesChannelRepository)
	updateChannelsFromYoutubeHandler := handlers2.NewUpdateChannelsFromYoutubeHandler(updateChannelsFromYoutube)
	clipRepository := firestore.NewClipRepository(firestoreClient)
	repositoriesClipRepository := firestore.ProvideClipRepository(firestoreClient, clipRepository)
	getClipsByPeriod := usecases3.NewGetClipsByPeriod(repositoriesClipRepository)
	getClipsByPeriodHandler := handlers3.NewGetClipsByPeriodHandler(getClipsByPeriod)
	updateClipsByPeriod := usecases3.NewUpdateClipsByPeriod(repositoriesClipRepository, youTubeService)
	updateClipsHandler := handlers3.NewUpdateClipsHandler(updateClipsByPeriod)
	cronHandler := handlers4.NewCronHandler(updateClipsByPeriod, updateSongs)
	liveStreamRepository := firestore.NewLiveStreamRepository(firestoreClient)
	repositoriesLiveStreamRepository := firestore.ProvideLiveStreamRepository(firestoreClient, liveStreamRepository)
	getLiveStreamsByPeriod := usecases4.NewGetLiveStreamsByPeriod(repositoriesLiveStreamRepository)
	getLiveStreamsByPeriodHandler := handlers5.NewGetLiveStreamsByPeriodHandler(getLiveStreamsByPeriod)
	discordService, err := ports2.NewDiscordService()
	if err != nil {
		return nil, nil, err
	}
	discordSendMessage := usecases5.NewDiscordSendMessage(discordService, repositoriesLiveStreamRepository)
	discordSendMessageHandler := handlers6.NewDiscordSendMessageHandler(discordSendMessage)
	application := NewApplication(getAllSongsHandler, createSongHandler, updateSongsHandler, addNewSongHandler, getChannelsHandler, createChannelHandler, updateChannelsFromYoutubeHandler, getClipsByPeriodHandler, updateClipsHandler, cronHandler, getLiveStreamsByPeriodHandler, discordSendMessageHandler)
	return application, func() {
	}, nil
}

// wire.go:

// Application is the main application struct which holds all the dependencies together.
type Application struct {
	GetAllSongsHandler               *handlers.GetAllSongsHandler
	CreateSongHandler                *handlers.CreateSongHandler
	UpdateSongsHandler               *handlers.UpdateSongsHandler
	AddNewSongHandler                *handlers.AddNewSongHandler
	GetChannelsHandler               *handlers2.GetChannelsHandler
	CreateChannelHandler             *handlers2.CreateChannelHandler
	UpdateChannelsFromYoutubeHandler *handlers2.UpdateChannelsFromYoutubeHandler
	GetClipsByPeriodHandler          *handlers3.GetClipsByPeriodHandler
	UpdateClipsHandler               *handlers3.UpdateClipsHandler
	CronHandler                      *handlers4.CronHandler
	GetLiveStreamsByPeriodHandler    *handlers5.GetLiveStreamsByPeriodHandler
	DiscordSendMessageHandler        *handlers6.DiscordSendMessageHandler
}

// NewApplication creates a new Application.
func NewApplication(getAllSongsHandler *handlers.GetAllSongsHandler, createSongHandler *handlers.CreateSongHandler, updateSongsHandler *handlers.UpdateSongsHandler, addNewSongHandler *handlers.AddNewSongHandler,
	getChannelsHandler *handlers2.GetChannelsHandler, createChannelHandler *handlers2.CreateChannelHandler, updateChannelsFromYoutubeHandler *handlers2.UpdateChannelsFromYoutubeHandler,
	getClipsByPeriodHandler *handlers3.GetClipsByPeriodHandler, updateClipsHandler *handlers3.UpdateClipsHandler,
	cronHandler *handlers4.CronHandler,
	getLiveStreamsByPeriodHandler *handlers5.GetLiveStreamsByPeriodHandler,
	discordGetLiveStreamsHandler *handlers6.DiscordSendMessageHandler,
) *Application {
	return &Application{
		GetAllSongsHandler:               getAllSongsHandler,
		CreateSongHandler:                createSongHandler,
		UpdateSongsHandler:               updateSongsHandler,
		AddNewSongHandler:                addNewSongHandler,
		GetChannelsHandler:               getChannelsHandler,
		CreateChannelHandler:             createChannelHandler,
		UpdateChannelsFromYoutubeHandler: updateChannelsFromYoutubeHandler,
		GetClipsByPeriodHandler:          getClipsByPeriodHandler,
		UpdateClipsHandler:               updateClipsHandler,
		CronHandler:                      cronHandler,
		GetLiveStreamsByPeriodHandler:    getLiveStreamsByPeriodHandler,
		DiscordSendMessageHandler:        discordGetLiveStreamsHandler,
	}
}
