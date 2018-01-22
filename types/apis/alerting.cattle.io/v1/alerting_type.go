package v1

import (
	"github.com/rancher/norman/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Notifier struct {
	types.Namespaced

	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	DisplayName     string           `json:"displayName,omitempty"`
	Description     string           `json:"description,omitempty"`
	EmailConfig     *EmailConfig     `json:"emailConfig,omitempty"`
	SlackConfig     *SlackConfig     `json:"slackConfig,omitempty"`
	PagerdutyConfig *PagerdutyConfig `json:"pagerdutyConfig,omitempty"`
	WebhookConfig   *WebhookConfig   `json:"webhookConfig,omitempty"`
}

type EmailConfig struct {
	SmtpHost     string `json:"smtpHost,omitempty"`
	SmtpPort     string `json:"smtpPort,omitempty"`
	SmtpUsername string `json:"smtpUsername,omitempty"`
	SmtpPassword string `json:"smtpPassword,omitempty"`
	Receiver     string `json:"receiver,omitempty"`
	RequireTLS   string `json:"requireTLS,omitempty"`
}

type SlackConfig struct {
	WebhookURL string `json:"webhookURL,omitempty"`
	Channel    string `json:"channel,omitempty"`
}

type PagerdutyConfig struct {
	ServiceKey string `json:"serviceKey,omitempty"`
}

type WebhookConfig struct {
	URL string `json:"url,omitempty"`
}
