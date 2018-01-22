package config

import (
	"context"

	alertingv1 "github.com/rancher/alerting-api/types/apis/alerting.cattle.io/v1"
	alertingSchema "github.com/rancher/alerting-api/types/apis/alerting.cattle.io/v1/schema"
	"github.com/rancher/norman/controller"
	"github.com/rancher/norman/event"
	"github.com/rancher/norman/signal"
	"github.com/rancher/norman/types"
	"github.com/sirupsen/logrus"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
)

type AlertingContext struct {
	eventBroadcaster record.EventBroadcaster

	RESTConfig        rest.Config
	UnversionedClient rest.Interface
	K8sClient         kubernetes.Interface
	Events            record.EventRecorder
	EventLogger       event.Logger
	Schemas           *types.Schemas
	Scheme            *runtime.Scheme

	Alert alertingv1.Interface
}

func (c *AlertingContext) controllers() []controller.Starter {
	return []controller.Starter{
		c.Alert,
	}
}

func NewAlertContext(config rest.Config) (*AlertingContext, error) {
	var err error

	context := &AlertingContext{
		RESTConfig: config,
	}

	context.Alert, err = alertingv1.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	context.K8sClient, err = kubernetes.NewForConfig(&config)
	if err != nil {
		return nil, err
	}

	dynamicConfig := config
	if dynamicConfig.NegotiatedSerializer == nil {
		configConfig := dynamic.ContentConfig()
		dynamicConfig.NegotiatedSerializer = configConfig.NegotiatedSerializer
	}

	context.UnversionedClient, err = rest.UnversionedRESTClientFor(&dynamicConfig)
	if err != nil {
		return nil, err
	}

	context.Schemas = types.NewSchemas().
		AddSchemas(alertingSchema.Schemas)

	context.eventBroadcaster = record.NewBroadcaster()
	context.Events = context.eventBroadcaster.NewRecorder(context.Scheme, v1.EventSource{
		Component: "CattleManagementServer",
	})
	context.EventLogger = event.NewLogger(context.Events)

	return context, err
}

func (c *AlertingContext) Start(ctx context.Context) error {
	logrus.Info("Starting alert controllers")

	watcher := c.eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{
		Interface: c.K8sClient.CoreV1().Events(""),
	})

	go func() {
		<-ctx.Done()
		watcher.Stop()
	}()

	return controller.SyncThenStart(ctx, 5, c.controllers()...)
}

func (c *AlertingContext) StartAndWait() error {
	ctx := signal.SigTermCancelContext(context.Background())
	c.Start(ctx)
	<-ctx.Done()
	return ctx.Err()
}
