package client

const (
	TargetPodType             = "targetPod"
	TargetPodFieldIsRunning   = "isRunning"
	TargetPodFieldIsScheduled = "isScheduled"
	TargetPodFieldRestartTime = "restartTime"
)

type TargetPod struct {
	IsRunning   *bool  `json:"isRunning,omitempty"`
	IsScheduled *bool  `json:"isScheduled,omitempty"`
	RestartTime *int64 `json:"restartTime,omitempty"`
}
