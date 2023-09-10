package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sugar-cat7/vspo-common-api/domain/ports"
	"github.com/sugar-cat7/vspo-common-api/infrastructure/http/mappers"
	usecases "github.com/sugar-cat7/vspo-common-api/usecases/discord"
)

// DiscordSendMessageHandler is a handler for getting all liveStreams.
type DiscordSendMessageHandler struct {
	discordSendMessage *usecases.DiscordSendMessage
	slackService       ports.SlackService
}

// NewDiscordSendMessageHandler creates a new DiscordSendMessageHandler.
func NewDiscordSendMessageHandler(u *usecases.DiscordSendMessage, s ports.SlackService) *DiscordSendMessageHandler {
	return &DiscordSendMessageHandler{
		discordSendMessage: u,
		slackService:       s,
	}
}

type VideosResponse mappers.VideosResponse

func (h *DiscordSendMessageHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// query param
	start := r.URL.Query().Get("start_date")
	end := r.URL.Query().Get("end_date")
	countryCode := r.URL.Query().Get("country_code")

	liveStreams, err := h.discordSendMessage.Execute(start, end, countryCode)

	if err != nil {
		errMsg := fmt.Sprintf("Discord Error: %s", err.Error())
		_ = h.slackService.SendMessage(errMsg)
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
