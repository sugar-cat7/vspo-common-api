//go:generate mockgen -destination=../../mocks/ports/mock_slack_port.go -package=mocks github.com/sugar-cat7/vspo-common-api/domain/ports SlackService
package ports

// SlackService is an interface for a Slack implementation of a song service.
type SlackService interface {
	SendMessage(message string) error
}
