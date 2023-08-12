package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/infrastructure/http/mappers"
	usecases "github.com/sugar-cat7/vspo-common-api/usecases/clip"
)

// GetClipsByPeriodHandler is a handler for getting all clips.
type GetClipsByPeriodHandler struct {
	getClipsByPeriodUsecase *usecases.GetClipsByPeriod
}

// NewGetClipsByPeriodHandler creates a new GetClipsByPeriodHandler.
func NewGetClipsByPeriodHandler(u *usecases.GetClipsByPeriod) *GetClipsByPeriodHandler {
	return &GetClipsByPeriodHandler{
		getClipsByPeriodUsecase: u,
	}
}

type VideosResponse mappers.VideosResponse

// @Summary Get all clips
// @Description Retrieve all clips
// @Accept  json
// @Produce  json
// @Success 200 {object} VideosResponse
// @Router /clips [get]
func (h *GetClipsByPeriodHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// query param
	start := r.URL.Query().Get("start_date")
	end := r.URL.Query().Get("end_date")
	clips, err := h.getClipsByPeriodUsecase.Execute(start, end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(mappers.MapVideosToResponse(clips))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// UpdateClipsHandler is a handler for updating clips from YouTube.
type UpdateClipsHandler struct {
	updateClipsUsecase *usecases.UpdateClipsByPeriod
}

// NewUpdateClipsHandler creates a new UpdateClipsHandler.
func NewUpdateClipsHandler(u *usecases.UpdateClipsByPeriod) *UpdateClipsHandler {
	return &UpdateClipsHandler{
		updateClipsUsecase: u,
	}
}

// @Summary Update clips from YouTube
// @Description Update clips based on provided cronType
// @Accept  json
// @Produce  json
// @Param cronType body string true "Type of the cron"
// @Success 200 {string} string "Clips updated successfully"
// @Router /clips [put]
func (h *UpdateClipsHandler) Handle(w http.ResponseWriter, r *http.Request) {
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
	videos, err := h.updateClipsUsecase.Execute(cronType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(videos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
