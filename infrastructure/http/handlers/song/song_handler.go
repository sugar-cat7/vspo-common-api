package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/infrastructure/http/mappers"
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

type VideosResponse mappers.VideosResponse

// @Summary Get all songs
// @Description Retrieve all songs
// @Accept  json
// @Produce  json
// @Success 200 {object} VideosResponse
// @Router /songs [get]
func (h *GetAllSongsHandler) Handle(w http.ResponseWriter, r *http.Request) {
	songs, err := h.getAllSongsUsecase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(mappers.MapVideosToResponse(songs))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// UpdateSongsHandler is a handler for updating songs from Youtube.
type UpdateSongsHandler struct {
	updateSongsUsecase *usecases.UpdateSongs
}

// NewUpdateSongsHandler creates a new UpdateSongsHandler.
func NewUpdateSongsHandler(u *usecases.UpdateSongs) *UpdateSongsHandler {
	return &UpdateSongsHandler{
		updateSongsUsecase: u,
	}
}

// @Summary Update songs from Youtube
// @Description Update songs based on provided cronType
// @Accept  json
// @Produce  json
// @Param cronType body string true "Type of the cron"
// @Success 200 {string} string "Songs updated successfully"
// @Router /songs [put]
func (h *UpdateSongsHandler) Handle(w http.ResponseWriter, r *http.Request) {
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

// @Summary Create Song from Youtube
// @Description Updates songs by fetching from Youtube using provided Video IDs.
// @Accept  json
// @Produce  json
// @Param videoIds body []string true "Array of Video IDs"
// @Success 200 {string} string "Songs updated successfully"
// @Router /songs [post]
func (h *CreateSongHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// Define a new struct type to hold the request body parameters
	type requestBody struct {
		VideoIds []string `json:"videoIds"`
	}

	// Create a new instance of requestBody
	rb := &requestBody{}

	// Decode the request body into the rb instance
	err := json.NewDecoder(r.Body).Decode(rb)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Execute the use case with the cronType from the request body
	err = h.createSongsUsecase.Execute(rb.VideoIds)
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

// AddNewSongHandler is a handler for updating songs from Youtube.
type AddNewSongHandler struct {
	addNewSongsUsecase *usecases.AddNewSong
}

// NewAddNewSongHandler creates a new AddNewSongHandler.
func NewAddNewSongHandler(u *usecases.AddNewSong) *AddNewSongHandler {
	return &AddNewSongHandler{
		addNewSongsUsecase: u,
	}
}

// @Summary Create Song from Youtube
// @Description Updates songs by fetching from Youtube using provided Video IDs.
// @Accept  json
// @Produce  json
// @Param videoIds body []string true "Array of Video IDs"
// @Success 200 {string} string "Songs updated successfully"
// @Router /songs [post]
func (h *AddNewSongHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// Define a new struct type to hold the request body parameters
	type requestBody struct {
		PlayListIds []string `json:"playListIds"`
	}

	// Create a new instance of requestBody
	rb := &requestBody{}

	// Decode the request body into the rb instance
	err := json.NewDecoder(r.Body).Decode(rb)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Execute the use case with the cronType from the request body
	songs, err := h.addNewSongsUsecase.Execute(rb.PlayListIds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(songs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
