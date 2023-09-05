package entities

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type SlackPayload struct {
	Text string `json:"text"`
}

type SlackNotifier struct {
	webhookURL string
}

func NewSlackNotifier() *SlackNotifier {
	url, ok := os.LookupEnv("SLACK_WEBHOOK_URL")
	if !ok {
		panic("SLACK_WEBHOOK_URL not set")
	}
	return &SlackNotifier{webhookURL: url}
}

func (s *SlackNotifier) SendMessage(message string) error {
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
