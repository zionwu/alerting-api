package v1

import (
	"context"
	"sync"

	"github.com/rancher/norman/clientbase"
	"github.com/rancher/norman/controller"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

type Interface interface {
	RESTClient() rest.Interface
	controller.Starter

	NotifiersGetter
	AlertsGetter
}

type Client struct {
	sync.Mutex
	restClient rest.Interface
	starters   []controller.Starter

	notifierControllers map[string]NotifierController
	alertControllers    map[string]AlertController
}

func NewForConfig(config rest.Config) (Interface, error) {
	if config.NegotiatedSerializer == nil {
		configConfig := dynamic.ContentConfig()
		config.NegotiatedSerializer = configConfig.NegotiatedSerializer
	}

	restClient, err := rest.UnversionedRESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &Client{
		restClient: restClient,

		notifierControllers: map[string]NotifierController{},
		alertControllers:    map[string]AlertController{},
	}, nil
}

func (c *Client) RESTClient() rest.Interface {
	return c.restClient
}

func (c *Client) Sync(ctx context.Context) error {
	return controller.Sync(ctx, c.starters...)
}

func (c *Client) Start(ctx context.Context, threadiness int) error {
	return controller.Start(ctx, threadiness, c.starters...)
}

type NotifiersGetter interface {
	Notifiers(namespace string) NotifierInterface
}

func (c *Client) Notifiers(namespace string) NotifierInterface {
	objectClient := clientbase.NewObjectClient(namespace, c.restClient, &NotifierResource, NotifierGroupVersionKind, notifierFactory{})
	return &notifierClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}

type AlertsGetter interface {
	Alerts(namespace string) AlertInterface
}

func (c *Client) Alerts(namespace string) AlertInterface {
	objectClient := clientbase.NewObjectClient(namespace, c.restClient, &AlertResource, AlertGroupVersionKind, alertFactory{})
	return &alertClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}
