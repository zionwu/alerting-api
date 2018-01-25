package client

const (
	TargetNodeType              = "targetNode"
	TargetNodeFieldCPUThreshold = "diskThreshold"
	TargetNodeFieldIsReady      = "isReady"
	TargetNodeFieldSelector     = "seletor"
)

type TargetNode struct {
	CPUThreshold *int64            `json:"diskThreshold,omitempty"`
	IsReady      *bool             `json:"isReady,omitempty"`
	Selector     map[string]string `json:"seletor,omitempty"`
}
