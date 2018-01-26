package controller

import (
	"github.com/rancher/alerting-api/types/apis/alerting.cattle.io/v1"
	"github.com/rancher/alerting-api/types/config"
	"k8s.io/apimachinery/pkg/labels"
	
)

type Controller struct {
	AlertLister    v1.AlertLister
	NotifierLister v1.NotifierLister
}

func Register(alertContext *config.AlertingContext) {
	alertClient := alertContext.Alert.Alerts("")
	notifierClient := alertContext.Notifier.Notifier("")
	
	
	c := &Controller{
		AlertLister: alertClient.Controller().Lister()
		NotifierLister: notifierClient.Controller().Lister()
	}
	alertContext.Alert.Alerts("").AddHandler("alert-sync", c.AlertSync)
	alertContext.Notifier.Notifiers("").AddHandler("notifier-sync", c.NotifierSync)
	
}

func (c *Controller) AlertSync(key string, alert *v1.Alert) error {
	return sync()
}

func (c *Controller) NotifierSync(key string, alert *v1.Notifier) error {
	return sync()
}


func (c *Controller) sync() error {
	alerts := c.AlertLister.List("", labels.NewSelector())
	notifiers := c.NotifierLister.List("", labels.NewSelector())
	

}