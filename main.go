package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sugar-cat7/vspo-common-api/app/di"
	_ "github.com/sugar-cat7/vspo-common-api/infrastructure/http/handlers/channel"
	_ "github.com/sugar-cat7/vspo-common-api/infrastructure/http/handlers/clip"
	_ "github.com/sugar-cat7/vspo-common-api/infrastructure/http/handlers/song"
)

// @title VSPO Common API
// @version 1.0
// @description This is the API documentation for VSPO Common services.
// @BasePath /api/v1

func main() {
	r := mux.NewRouter()
	app, cleanup, err := di.InitializeApplication()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}
	defer cleanup()

	apiRouter := r.PathPrefix("/api/v1").Subrouter()
	apiRouter.HandleFunc("/songs", app.GetAllSongsHandler.Handle).Methods("GET")
	apiRouter.HandleFunc("/clips", app.GetClipsByPeriodHandler.Handle).Methods("GET")
	apiRouter.HandleFunc("/channels", app.GetChannelsHandler.Handle).Methods("GET")
	apiRouter.HandleFunc("/livestreams", app.GetLiveStreamsByPeriodHandler.Handle).Methods("GET")
	apiRouter.HandleFunc("/discord-livestreams", app.DiscordSendMessageHandler.Handle).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting server on 127.0.0.1:8000")
	log.Fatal(srv.ListenAndServe())
}
