package client

import (
	"github.com/rancher/norman/types"
)

const (
	AlertType                       = "alert"
	AlertFieldAlertState            = "alertState"
	AlertFieldAnnotations           = "annotations"
	AlertFieldCreated               = "created"
	AlertFieldCreatorID             = "creatorId"
	AlertFieldDescription           = "description"
	AlertFieldInitialWaitSeconds    = "initialWaitSeconds"
	AlertFieldLabels                = "labels"
	AlertFieldName                  = "name"
	AlertFieldNamespaceId           = "namespaceId"
	AlertFieldNotifierList          = "notifier"
	AlertFieldOwnerReferences       = "ownerReferences"
	AlertFieldRemoved               = "removed"
	AlertFieldRepeatIntervalSeconds = "repeatIntervalSeconds"
	AlertFieldSeverity              = "severity"
	AlertFieldStartedAt             = "startedAt"
	AlertFieldTargetNode            = "targetNode"
	AlertFieldTargetPod             = "targetPod"
	AlertFieldTargetSystemService   = "targetSystemService"
	AlertFieldTargetWorkload        = "targetWorkload"
	AlertFieldUuid                  = "uuid"
)

type Alert struct {
	types.Resource
	AlertState            string               `json:"alertState,omitempty"`
	Annotations           map[string]string    `json:"annotations,omitempty"`
	Created               string               `json:"created,omitempty"`
	CreatorID             string               `json:"creatorId,omitempty"`
	Description           string               `json:"description,omitempty"`
	InitialWaitSeconds    *int64               `json:"initialWaitSeconds,omitempty"`
	Labels                map[string]string    `json:"labels,omitempty"`
	Name                  string               `json:"name,omitempty"`
	NamespaceId           string               `json:"namespaceId,omitempty"`
	NotifierList          *Recipient           `json:"notifier,omitempty"`
	OwnerReferences       []OwnerReference     `json:"ownerReferences,omitempty"`
	Removed               string               `json:"removed,omitempty"`
	RepeatIntervalSeconds *int64               `json:"repeatIntervalSeconds,omitempty"`
	Severity              string               `json:"severity,omitempty"`
	StartedAt             string               `json:"startedAt,omitempty"`
	TargetNode            *TargetNode          `json:"targetNode,omitempty"`
	TargetPod             *TargetPod           `json:"targetPod,omitempty"`
	TargetSystemService   *TargetSystemService `json:"targetSystemService,omitempty"`
	TargetWorkload        *TargetWorkload      `json:"targetWorkload,omitempty"`
	Uuid                  string               `json:"uuid,omitempty"`
}
type AlertCollection struct {
	types.Collection
	Data   []Alert `json:"data,omitempty"`
	client *AlertClient
}

type AlertClient struct {
	apiClient *Client
}

type AlertOperations interface {
	List(opts *types.ListOpts) (*AlertCollection, error)
	Create(opts *Alert) (*Alert, error)
	Update(existing *Alert, updates interface{}) (*Alert, error)
	ByID(id string) (*Alert, error)
	Delete(container *Alert) error
}

func newAlertClient(apiClient *Client) *AlertClient {
	return &AlertClient{
		apiClient: apiClient,
	}
}

func (c *AlertClient) Create(container *Alert) (*Alert, error) {
	resp := &Alert{}
	err := c.apiClient.Ops.DoCreate(AlertType, container, resp)
	return resp, err
}

func (c *AlertClient) Update(existing *Alert, updates interface{}) (*Alert, error) {
	resp := &Alert{}
	err := c.apiClient.Ops.DoUpdate(AlertType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *AlertClient) List(opts *types.ListOpts) (*AlertCollection, error) {
	resp := &AlertCollection{}
	err := c.apiClient.Ops.DoList(AlertType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *AlertCollection) Next() (*AlertCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &AlertCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *AlertClient) ByID(id string) (*Alert, error) {
	resp := &Alert{}
	err := c.apiClient.Ops.DoByID(AlertType, id, resp)
	return resp, err
}

func (c *AlertClient) Delete(container *Alert) error {
	return c.apiClient.Ops.DoResourceDelete(AlertType, &container.Resource)
}
