/*
package alert

import (
	"context"

	"github.com/rancher/alerting-api/types/apis/alerting.cattle.io/v1"
	"github.com/rancher/alerting-api/types/config"
)

func Register(ctx context.Context, alert *config.AlertingContext) {

	clusterClient := management.Management.Clusters("")

	alert.Alert.Alerts("").AddHandler("", )
	clusterClient.AddLifecycle("cluster-agent-controller", lifecycle)
}

type ClusterLifecycle struct {
	Manager *Manager
	ctx     context.Context
}

func (c *ClusterLifecycle) Create(obj *v3.Cluster) (*v3.Cluster, error) {
	return nil, nil
}

func (c *ClusterLifecycle) Remove(obj *v3.Cluster) (*v3.Cluster, error) {
	c.Manager.Stop(c.ctx, obj)
	return nil, nil
}

func (c *ClusterLifecycle) Updated(obj *v3.Cluster) (*v3.Cluster, error) {
	return nil, c.Manager.Start(c.ctx, obj)
}

*/