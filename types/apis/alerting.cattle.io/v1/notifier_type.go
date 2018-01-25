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
	SmtpConfig      *SmtpConfig      `json:"smtpConfig,omitempty"`
	SlackConfig     *SlackConfig     `json:"slackConfig,omitempty"`
	PagerdutyConfig *PagerdutyConfig `json:"pagerdutyConfig,omitempty"`
	WebhookConfig   *WebhookConfig   `json:"webhookConfig,omitempty"`
}

type SmtpConfig struct {
	Host             string `json:"host,omitempty"`
	Port             int    `json:"port,omitempty"`
	Username         string `json:"username,omitempty"`
	Password         string `json:"password,omitempty"`
	DefaultRecipient string `json:"defaultRecipient,omitempty"`
	TLS              bool   `json:"tls,omitempty"`
}

type SlackConfig struct {
	URL     string `json:"url,omitempty"`
	Channel string `json:"channel,omitempty"`
}

type PagerdutyConfig struct {
	ServiceKey string `json:"serviceKey,omitempty"`
}

type WebhookConfig struct {
	URL string `json:"url,omitempty"`
}
