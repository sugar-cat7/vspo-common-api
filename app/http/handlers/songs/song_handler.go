package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	usecases "github.com/sugar-cat7/vspo-common-api/usecases/song"
)

// GetAllSongsHandler is a handler for getting all songs.
type GetAllSongsHandler struct {
	getAllSongsUsecase *usecases.GetAllSongs
}

// NewGetAllSongsHandler creates a new GetAllSongsHandler.
func NewGetAllSongsHandler(u *usecases.GetAllSongs) *GetAllSongsHandler {
	return &GetAllSongsHandler{
		getAllSongsUsecase: u,
	}
}

// Handle returns all songs.
func (h *GetAllSongsHandler) Handle(w http.ResponseWriter, r *http.Request) {
	songs, err := h.getAllSongsUsecase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(songs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// UpdateSongsFromYoutubeHandler is a handler for updating songs from Youtube.
type UpdateSongsFromYoutubeHandler struct {
	updateSongsUsecase *usecases.UpdateSongsFromYoutube
}

// NewUpdateSongsFromYoutubeHandler creates a new UpdateSongsFromYoutubeHandler.
func NewUpdateSongsFromYoutubeHandler(u *usecases.UpdateSongsFromYoutube) *UpdateSongsFromYoutubeHandler {
	return &UpdateSongsFromYoutubeHandler{
		updateSongsUsecase: u,
	}
}

// Handle updates songs from Youtube.
func (h *UpdateSongsFromYoutubeHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// Define a new struct type to hold the request body parameters
	type requestBody struct {
		CronType string `json:"cronType"`
	}

	// Create a new instance of requestBody
	rb := &requestBody{}

	// Decode the request body into the rb instance
	err := json.NewDecoder(r.Body).Decode(rb)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the cronType
	cronType, err := entities.ParseCronType(rb.CronType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Execute the use case with the cronType from the request body
	err = h.updateSongsUsecase.Execute(cronType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Songs updated successfully"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// CreateSongHandler is a handler for updating songs from Youtube.
type CreateSongHandler struct {
	createSongsUsecase *usecases.CreateSong
}

// NewCreateSongHandler creates a new CreateSongHandler.
func NewCreateSongHandler(u *usecases.CreateSong) *CreateSongHandler {
	return &CreateSongHandler{
		createSongsUsecase: u,
	}
}

// Handle updates songs from Youtube.
func (h *CreateSongHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// Execute the use case with the cronType from the request body
	err := h.createSongsUsecase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Songs updated successfully"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
