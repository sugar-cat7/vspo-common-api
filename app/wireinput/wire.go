//go:generate go run github.com/google/wire/cmd/wire@v0.5.0 -o ../di/wire_gen.go
//go:build wireinject

package wireinput

import (
	"github.com/google/wire"
	handlers "github.com/sugar-cat7/vspo-common-api/app/http/handlers/songs"
	"github.com/sugar-cat7/vspo-common-api/domain/services"
	"github.com/sugar-cat7/vspo-common-api/infrastructure/firestore"
	"github.com/sugar-cat7/vspo-common-api/infrastructure/youtube"
	usecases "github.com/sugar-cat7/vspo-common-api/usecases/song"
)

// Application is the main application struct which holds all the dependencies together.
type Application struct {
	GetAllSongsHandler            *handlers.GetAllSongsHandler
	CreateSongHandler             *handlers.CreateSongHandler
	UpdateSongsFromYoutubeHandler *handlers.UpdateSongsFromYoutubeHandler
}

// NewApplication creates a new Application.
func NewApplication(getAllSongsHandler *handlers.GetAllSongsHandler, createSongHandler *handlers.CreateSongHandler, updateSongsFromYoutubeHandler *handlers.UpdateSongsFromYoutubeHandler) *Application {
	return &Application{
		GetAllSongsHandler:            getAllSongsHandler,
		CreateSongHandler:             createSongHandler,
		UpdateSongsFromYoutubeHandler: updateSongsFromYoutubeHandler,
	}
}

// InitializeApplication initializes the entire application with all its dependencies using wire.
func InitializeApplication() (*Application, func(), error) {
	panic(wire.Build(services.Set, firestore.Set, youtube.Set, usecases.Set, handlers.Set, NewApplication))
}
