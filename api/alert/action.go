package alert

import (
	"errors"
	"net/http"
	"strings"

	"github.com/rancher/alerting-api/types/config"
	"github.com/rancher/norman/types"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Formatter(apiContext *types.APIContext, resource *types.RawResource) {
	resource.Actions["activate"] = apiContext.URLBuilder.Action("activate", resource)
	resource.Actions["deactivate"] = apiContext.URLBuilder.Action("deactivate", resource)
	resource.Actions["mute"] = apiContext.URLBuilder.Action("mute", resource)
	resource.Actions["unmute"] = apiContext.URLBuilder.Action("unmute", resource)

}

type Handler struct {
	AlertContext *config.AlertingContext
}

func (h *Handler) ActionHandler(actionName string, action *types.Action, request *types.APIContext) error {
	logrus.Infof("do activity action:%s", actionName)

	//TODO: check if the right way to do this.
	//BUG: state keep active in the frontend
	store := request.Schema.Store
	if store == nil {
		return errors.New("no user store available")
	}

	parts := strings.Split(request.ID, ":")
	ns := parts[0]
	id := parts[1]

	client := h.AlertContext.Alert.Alerts(ns)
	alert, err := client.Get(id, metav1.GetOptions{})
	if err != nil {
		logrus.Errorf("Error while getting alert:%v", err)
		return err
	}

	switch actionName {
	case "activate":
		if alert.State == "inactive" {
			alert.State = "active"
		} else {
			return errors.New("the alert state is not inactive")
		}

	case "deactivate":
		if alert.State == "active" {
			alert.State = "inactive"
		} else {
			return errors.New("the alert state is not active")
		}

	case "mute":
		if alert.State == "alerting" {
			alert.State = "muted"
		} else {
			return errors.New("the alert state is not alerting")
		}

	case "unmute":
		if alert.State == "muted" {
			alert.State = "alerting"
		} else {
			return errors.New("the alert state is not muted")
		}

	}

	alert, err = client.Update(alert)
	if err != nil {
		logrus.Errorf("Error while updating alert:%v", err)
		return err
	}

	//TODO: how to write data back
	request.WriteResponse(http.StatusOK, alert)
	return nil
}
