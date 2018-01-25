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

	DisplayName    string            `json:"displayName,omitempty"`
	Description    string            `json:"description,omitempty"`
	Severity       string            `json:"severity,omitempty"`
	NotifierId     string            `json:"notifierId,omitempty" norman:"type=reference[notifier]"`
	InitialWait    string            `json:"initialWait,omitempty"`
	RepeatInterval string            `json:"repeatInterval,omitempty"`
	TargetType     string            `json:"targetType,omitempty"`
	TargetID       string            `json:"targetID,omitempty"`
	TargetSelector map[string]string `json:"targetSelector,omitempty"`
	StartedAt      string            `json:"startedAt,omitempty"`
	State          string            `json:"state,omitempty"`
	NodeRule       NodeRule          `json:"nodeRule,omitempty"`
	PodRule        PodRule           `json:"podRule,omitempty"`
	WorkloadRule   WorkloadRule      `json:"workloadRule,omitempty"`
}

//TODO: what node rule should we support
type NodeRule struct {
}

//TODO: precam rule

type PodRule struct {
	Unhealthy bool `json:"unhealthy,omitempty"`
}

type WorkloadRule struct {
	UnavailablePercentage int `json:"unavailablePercentage,omitempty"`
}
