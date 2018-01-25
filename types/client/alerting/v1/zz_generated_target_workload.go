package client

const (
	TargetWorkloadType                       = "targetWorkload"
	TargetWorkloadFieldSelector              = "seletor"
	TargetWorkloadFieldUnavailablePercentage = "unavailablePercentage"
)

type TargetWorkload struct {
	Selector              map[string]string `json:"seletor,omitempty"`
	UnavailablePercentage int64             `json:"unavailablePercentage,omitempty"`
}
