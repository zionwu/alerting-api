package client

const (
	SlackConfigType            = "slackConfig"
	SlackConfigFieldChannel    = "channel"
	SlackConfigFieldWebhookURL = "webhookURL"
)

type SlackConfig struct {
	Channel    string `json:"channel,omitempty"`
	WebhookURL string `json:"webhookURL,omitempty"`
}
