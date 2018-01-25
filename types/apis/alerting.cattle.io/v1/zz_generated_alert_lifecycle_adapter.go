package v1

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type AlertLifecycle interface {
	Create(obj *Alert) (*Alert, error)
	Remove(obj *Alert) (*Alert, error)
	Updated(obj *Alert) (*Alert, error)
}

type alertLifecycleAdapter struct {
	lifecycle AlertLifecycle
}

func (w *alertLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*Alert))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *alertLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*Alert))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *alertLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*Alert))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewAlertLifecycleAdapter(name string, clusterScoped bool, client AlertInterface, l AlertLifecycle) AlertHandlerFunc {
	adapter := &alertLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *Alert) error {
		if obj == nil {
			return syncFn(key, nil)
		}
		return syncFn(key, obj)
	}
}
