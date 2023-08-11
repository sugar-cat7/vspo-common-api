package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/sugar-cat7/vspo-common-api/infrastructure/http/mappers"
	usecases "github.com/sugar-cat7/vspo-common-api/usecases/channel"
)

// GetChannelsHandler is a handler for getting all channels.
type GetChannelsHandler struct {
	getChannelsUsecase *usecases.GetChannels
}

// NewGetChannelsHandler creates a new GetChannelsHandler.
func NewGetChannelsHandler(u *usecases.GetChannels) *GetChannelsHandler {
	return &GetChannelsHandler{
		getChannelsUsecase: u,
	}
}

type ChannelsResponse mappers.ChannelsResponse

// @Summary Get Channels
// @Description Retrieves all channels based on provided IDs.
// @Accept  json
// @Produce  json
// @Param ids query []string false "Comma-separated list of channel IDs"
// @Success 200 {object} ChannelsResponse
// @Router /channels [get]
func (h *GetChannelsHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// Parse the query parameters for "ids", which should be a comma-separated list of channel IDs.
	idsParam := r.URL.Query().Get("ids")
	ids := strings.Split(idsParam, ",")
	channels, err := h.getChannelsUsecase.Execute(ids)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(mappers.MapChannelsToResponse(channels))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// UpdateChannelsFromYoutubeHandler is a handler for updating channels from Youtube.
type UpdateChannelsFromYoutubeHandler struct {
	updateChannelsUsecase *usecases.UpdateChannelsFromYoutube
}

// NewUpdateChannelsFromYoutubeHandler creates a new UpdateChannelsFromYoutubeHandler.
func NewUpdateChannelsFromYoutubeHandler(u *usecases.UpdateChannelsFromYoutube) *UpdateChannelsFromYoutubeHandler {
	return &UpdateChannelsFromYoutubeHandler{
		updateChannelsUsecase: u,
	}
}

// @Summary Update Channels from Youtube
// @Description Updates channels by fetching from Youtube using provided Channel IDs.
// @Accept  json
// @Produce  json
// @Param channelIds body []string true "Array of Channel IDs"
// @Success 200 {string} string "Channels updated successfully"
// @Router /channels [put]
func (h *UpdateChannelsFromYoutubeHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// Define a new struct type to hold the request body parameters
	type requestBody struct {
		ChannelIds []string `json:"channelIds"`
	}

	// Create a new instance of requestBody
	rb := &requestBody{}

	// Decode the request body into the rb instance
	err := json.NewDecoder(r.Body).Decode(rb)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Execute the use case with the channelIds from the request body
	err = h.updateChannelsUsecase.Execute(rb.ChannelIds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Channels updated successfully"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// CreateChannelHandler is a handler for updating channels from Youtube.
type CreateChannelHandler struct {
	createChannelsUsecase *usecases.CreateChannel
}

// NewCreateChannelHandler creates a new CreateChannelHandler.
func NewCreateChannelHandler(u *usecases.CreateChannel) *CreateChannelHandler {
	return &CreateChannelHandler{
		createChannelsUsecase: u,
	}
}

// @Summary Create Channels from Youtube
// @Description Creates channels by fetching from Youtube using provided Channel IDs.
// @Accept  json
// @Produce  json
// @Param channelIds body []string true "Array of Channel IDs"
// @Success 200 {string} string "Channels created successfully"
// @Router /channels [post]
func (h *CreateChannelHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// Define a new struct type to hold the request body parameters
	type requestBody struct {
		ChannelIds []string `json:"channelIds"`
	}

	// Create a new instance of requestBody
	rb := &requestBody{}

	// Decode the request body into the rb instance
	err := json.NewDecoder(r.Body).Decode(rb)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Execute the use case with the channelIds from the request body
	err = h.createChannelsUsecase.Execute(rb.ChannelIds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Channels created successfully"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
