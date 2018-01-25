package v1

import (
	"context"

	"github.com/rancher/norman/clientbase"
	"github.com/rancher/norman/controller"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

var (
	AlertGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "Alert",
	}
	AlertResource = metav1.APIResource{
		Name:         "alerts",
		SingularName: "alert",
		Namespaced:   true,

		Kind: AlertGroupVersionKind.Kind,
	}
)

type AlertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Alert
}

type AlertHandlerFunc func(key string, obj *Alert) error

type AlertLister interface {
	List(namespace string, selector labels.Selector) (ret []*Alert, err error)
	Get(namespace, name string) (*Alert, error)
}

type AlertController interface {
	Informer() cache.SharedIndexInformer
	Lister() AlertLister
	AddHandler(name string, handler AlertHandlerFunc)
	AddClusterScopedHandler(name, clusterName string, handler AlertHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type AlertInterface interface {
	ObjectClient() *clientbase.ObjectClient
	Create(*Alert) (*Alert, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*Alert, error)
	Get(name string, opts metav1.GetOptions) (*Alert, error)
	Update(*Alert) (*Alert, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*AlertList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() AlertController
	AddHandler(name string, sync AlertHandlerFunc)
	AddLifecycle(name string, lifecycle AlertLifecycle)
	AddClusterScopedHandler(name, clusterName string, sync AlertHandlerFunc)
	AddClusterScopedLifecycle(name, clusterName string, lifecycle AlertLifecycle)
}

type alertLister struct {
	controller *alertController
}

func (l *alertLister) List(namespace string, selector labels.Selector) (ret []*Alert, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*Alert))
	})
	return
}

func (l *alertLister) Get(namespace, name string) (*Alert, error) {
	var key string
	if namespace != "" {
		key = namespace + "/" + name
	} else {
		key = name
	}
	obj, exists, err := l.controller.Informer().GetIndexer().GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(schema.GroupResource{
			Group:    AlertGroupVersionKind.Group,
			Resource: "alert",
		}, name)
	}
	return obj.(*Alert), nil
}

type alertController struct {
	controller.GenericController
}

func (c *alertController) Lister() AlertLister {
	return &alertLister{
		controller: c,
	}
}

func (c *alertController) AddHandler(name string, handler AlertHandlerFunc) {
	c.GenericController.AddHandler(name, func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*Alert))
	})
}

func (c *alertController) AddClusterScopedHandler(name, cluster string, handler AlertHandlerFunc) {
	c.GenericController.AddHandler(name, func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}

		if !controller.ObjectInCluster(cluster, obj) {
			return nil
		}

		return handler(key, obj.(*Alert))
	})
}

type alertFactory struct {
}

func (c alertFactory) Object() runtime.Object {
	return &Alert{}
}

func (c alertFactory) List() runtime.Object {
	return &AlertList{}
}

func (s *alertClient) Controller() AlertController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.alertControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(AlertGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &alertController{
		GenericController: genericController,
	}

	s.client.alertControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type alertClient struct {
	client       *Client
	ns           string
	objectClient *clientbase.ObjectClient
	controller   AlertController
}

func (s *alertClient) ObjectClient() *clientbase.ObjectClient {
	return s.objectClient
}

func (s *alertClient) Create(o *Alert) (*Alert, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*Alert), err
}

func (s *alertClient) Get(name string, opts metav1.GetOptions) (*Alert, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*Alert), err
}

func (s *alertClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*Alert, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*Alert), err
}

func (s *alertClient) Update(o *Alert) (*Alert, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*Alert), err
}

func (s *alertClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *alertClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *alertClient) List(opts metav1.ListOptions) (*AlertList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*AlertList), err
}

func (s *alertClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *alertClient) Patch(o *Alert, data []byte, subresources ...string) (*Alert, error) {
	obj, err := s.objectClient.Patch(o.Name, o, data, subresources...)
	return obj.(*Alert), err
}

func (s *alertClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *alertClient) AddHandler(name string, sync AlertHandlerFunc) {
	s.Controller().AddHandler(name, sync)
}

func (s *alertClient) AddLifecycle(name string, lifecycle AlertLifecycle) {
	sync := NewAlertLifecycleAdapter(name, false, s, lifecycle)
	s.AddHandler(name, sync)
}

func (s *alertClient) AddClusterScopedHandler(name, clusterName string, sync AlertHandlerFunc) {
	s.Controller().AddClusterScopedHandler(name, clusterName, sync)
}

func (s *alertClient) AddClusterScopedLifecycle(name, clusterName string, lifecycle AlertLifecycle) {
	sync := NewAlertLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.AddClusterScopedHandler(name, clusterName, sync)
}
