package handlers

import (
	"encoding/json"
	"net/http"

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
