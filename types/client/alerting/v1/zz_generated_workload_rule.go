package client

const (
	WorkloadRuleType                       = "workloadRule"
	WorkloadRuleFieldUnavailablePercentage = "unavailablePercentage"
)

type WorkloadRule struct {
	UnavailablePercentage *int64 `json:"unavailablePercentage,omitempty"`
}
