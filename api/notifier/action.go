package notifier

import (
	"github.com/rancher/norman/types"
	"github.com/sirupsen/logrus"
)

func Formatter(apiContext *types.APIContext, resource *types.RawResource) {
	resource.Actions["update"] = apiContext.URLBuilder.Action("update", resource)
	resource.Actions["remove"] = apiContext.URLBuilder.Action("remove", resource)
	resource.Actions["approve"] = apiContext.URLBuilder.Action("approve", resource)
	resource.Actions["deny"] = apiContext.URLBuilder.Action("deny", resource)
	resource.Actions["rerun"] = apiContext.URLBuilder.Action("rerun", resource)
	resource.Actions["stop"] = apiContext.URLBuilder.Action("stop", resource)
}

func ActionHandler(actionName string, action *types.Action, apiContext *types.APIContext) error {
	logrus.Infof("do activity action:%s", actionName)
	//TODO Implement Actions
	return nil
}
