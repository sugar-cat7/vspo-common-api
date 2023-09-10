package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	clip_usecases "github.com/sugar-cat7/vspo-common-api/usecases/clip"
	song_usecases "github.com/sugar-cat7/vspo-common-api/usecases/song"
)

// CronHandler is a handler for updating clips from YouTube.
type CronHandler struct {
	updateClipsUsecase *clip_usecases.UpdateClipsByPeriod
	updateSongsUsecase *song_usecases.UpdateSongs
}

// NewCronHandler creates a new CronHandler.
func NewCronHandler(updateClipsUsecase *clip_usecases.UpdateClipsByPeriod, updateSongsUsecase *song_usecases.UpdateSongs) *CronHandler {
	return &CronHandler{
		updateClipsUsecase: updateClipsUsecase,
		updateSongsUsecase: updateSongsUsecase,
	}
}

func (h *CronHandler) Handle(w http.ResponseWriter, r *http.Request) {
	type requestBody struct {
		CronType string `json:"cronType"`
	}

	rb := &requestBody{}
	if err := json.NewDecoder(r.Body).Decode(rb); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cronType, err := entities.ParseCronType(rb.CronType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var videos entities.Videos
	errCh := make(chan error, 2)

	// Define a helper function to process tasks concurrently and capture errors
	processConcurrently := func(fn func() error) {
		go func() {
			errCh <- fn()
		}()
	}

	processConcurrently(func() error {
		var err error
		videos, err = h.updateClipsUsecase.Execute(cronType)
		return err
	})

	processConcurrently(func() error {
		return h.updateSongsUsecase.Execute(cronType)
	})

	for i := 0; i < 2; i++ {
		if err := <-errCh; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(videos); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
