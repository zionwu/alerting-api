package v1

import (
	"github.com/rancher/norman/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Alert struct {
	types.Namespaced

	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	DisplayName           string    `json:"displayName,omitempty"`
	Description           string    `json:"description,omitempty"`
	Severity              string    `json:"severity,omitempty"`
	NotifierList          Recipient `json:"notifier,omitempty"`
	InitialWaitSeconds    int       `json:"initialWaitSeconds,omitempty"`
	RepeatIntervalSeconds int       `json:"repeatIntervalSeconds,omitempty"`

	StartedAt string `json:"startedAt,omitempty"`
	//TODO: status/state not working
	AlertState string `json:"alertState,omitempty"`

	TargetWorkload      TargetWorkload      `json:"targetWorkload,omitempty"`
	TargetPod           TargetPod           `json:"targetPod,omitempty"`
	TargetNode          TargetNode          `json:"targetNode,omitempty"`
	TargetSystemService TargetSystemService `json:"targetSystemService,omitempty"`
}

type Recipient struct {
	Recipient  string `json:"recipient,omitempty"`
	NotifierId string `json:"notifierId,omitempty" norman:"type=reference[notifier]"`
}

//TODO: what node rule should we support
type TargetNode struct {
	ID            string            `json:"id,omitempty"`
	Selector      map[string]string `json:"seletor,omitempty"`
	IsReady       bool              `json:"isReady,omitempty"`
	DiskThreshold int               `json:"diskThreshold,omitempty"`
	MemThreshold  int               `json:"diskThreshold,omitempty"`
	CPUThreshold  int               `json:"diskThreshold,omitempty"`
}

type TargetPod struct {
	ID          string `json:"id,omitempty"`
	IsRunning   bool   `json:"isRunning,omitempty"`
	IsScheduled bool   `json:"isScheduled,omitempty"`
	RestartTime int    `json:"restartTime,omitempty"`
}

type TargetWorkload struct {
	ID                    string            `json:"id,omitempty"`
	Selector              map[string]string `json:"seletor,omitempty"`
	UnavailablePercentage int               `json:"unavailablePercentage,omitempty"`
}

type TargetSystemService struct {
	Type string `json:"type,omitempty"`
}
