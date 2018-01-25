package setup

import (
	"context"

	"github.com/rancher/alerting-api/api/alert"
	"github.com/rancher/alerting-api/api/notifier"

	"github.com/rancher/alerting-api/store/scoped"
	alertingSchema "github.com/rancher/alerting-api/types/apis/alerting.cattle.io/v1/schema"
	"github.com/rancher/alerting-api/types/client/alerting/v1"
	"github.com/rancher/alerting-api/types/config"
	"github.com/rancher/norman/api/builtin"
	"github.com/rancher/norman/pkg/subscribe"
	"github.com/rancher/norman/store/crd"
	"github.com/rancher/norman/types"
)

var (
	crdVersions = []*types.APIVersion{
		&alertingSchema.Version,
	}
)

func Schemas(ctx context.Context, alerting *config.AlertingContext) error {
	schemas := alerting.Schemas
	subscribe.Register(&builtin.Version, schemas)
	Notifier(schemas)
	Alert(schemas, alerting)

	crdStore, err := crd.NewCRDStoreFromConfig(alerting.RESTConfig)
	if err != nil {
		return err
	}

	var crdSchemas []*types.Schema
	for _, version := range crdVersions {
		for _, schema := range schemas.SchemasForVersion(*version) {
			crdSchemas = append(crdSchemas, schema)
		}
	}

	if err := crdStore.AddSchemas(ctx, crdSchemas...); err != nil {
		return err
	}

	NamespacedTypes(schemas)

	return nil
}

func NamespacedTypes(schemas *types.Schemas) {
	for _, version := range crdVersions {
		for _, schema := range schemas.SchemasForVersion(*version) {
			if schema.Scope != types.NamespaceScope || schema.Store == nil {
				continue
			}

			for _, key := range []string{"projectId", "clusterId"} {
				ns, ok := schema.ResourceFields["namespaceId"]
				if !ok {
					continue
				}

				if _, ok := schema.ResourceFields[key]; !ok {
					continue
				}

				schema.Store = scoped.NewScopedStore(key, schema.Store)
				ns.Required = false
				schema.ResourceFields["namespaceId"] = ns
				break
			}
		}
	}
}

func Notifier(schemas *types.Schemas) {
	schema := schemas.Schema(&alertingSchema.Version, client.NotifierType)
	schema.Formatter = notifier.Formatter
	schema.CollectionFormatter = notifier.CollectionFormatter
	schema.ActionHandler = notifier.ActionHandler
}

func Alert(schemas *types.Schemas, alerting *config.AlertingContext) {
	schema := schemas.Schema(&alertingSchema.Version, client.AlertType)
	schema.Formatter = alert.Formatter

	handler := &alert.Handler{
		AlertContext: alerting,
	}
	schema.ActionHandler = handler.ActionHandler

}
