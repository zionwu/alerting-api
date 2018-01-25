package client

const (
	PodRuleType           = "podRule"
	PodRuleFieldUnhealthy = "unhealthy"
)

type PodRule struct {
	Unhealthy *bool `json:"unhealthy,omitempty"`
}
