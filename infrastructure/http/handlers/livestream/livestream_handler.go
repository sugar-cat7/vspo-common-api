package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sugar-cat7/vspo-common-api/infrastructure/http/mappers"
	usecases "github.com/sugar-cat7/vspo-common-api/usecases/livestream"
)

// GetLiveStreamsByPeriodHandler is a handler for getting all liveStreams.
type GetLiveStreamsByPeriodHandler struct {
	getLiveStreamsByPeriodUsecase *usecases.GetLiveStreamsByPeriod
}

// NewGetLiveStreamsByPeriodHandler creates a new GetLiveStreamsByPeriodHandler.
func NewGetLiveStreamsByPeriodHandler(u *usecases.GetLiveStreamsByPeriod) *GetLiveStreamsByPeriodHandler {
	return &GetLiveStreamsByPeriodHandler{
		getLiveStreamsByPeriodUsecase: u,
	}
}

type VideosResponse mappers.VideosResponse

// @Summary Get all liveStreams
// @Description Retrieve all liveStreams
// @Param start_date query string true "Start Date" example="2023-08-06"
// @Param end_date query string false "End Date" example="2023-08-12"
// @Accept  json
// @Produce  json
// @Success 200 {object} VideosResponse
// @Router /liveStreams [get]
func (h *GetLiveStreamsByPeriodHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// query param
	start := r.URL.Query().Get("start_date")
	end := r.URL.Query().Get("end_date")
	countryCode := r.URL.Query().Get("country_code")
	liveStreams, err := h.getLiveStreamsByPeriodUsecase.Execute(start, end, countryCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(mappers.MapVideosToResponse(liveStreams))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
