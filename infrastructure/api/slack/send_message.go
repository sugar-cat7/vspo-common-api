package ports

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/sugar-cat7/vspo-common-api/domain/ports"
)

type SlackPayload struct {
	Text string `json:"text"`
}

type slackServiceImpl struct {
	webhookURL string
}

func NewSlackService() (ports.SlackService, error) {
	url, ok := os.LookupEnv("SLACK_WEBHOOK_URL")
	if !ok {
		return nil, fmt.Errorf("SLACK_WEBHOOK_URL not set")
	}
	return &slackServiceImpl{webhookURL: url}, nil
}

func (s *slackServiceImpl) SendMessage(message string) error {
	data := SlackPayload{
		Text: "```" + message + "```",
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("Error marshalling payload: %v", err)
	}

	resp, err := http.Post(s.webhookURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Slack responded with non-OK status: %v", resp.Status)
	}

	return nil
}
