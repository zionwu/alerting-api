package schema

import (
	"github.com/rancher/alerting-api/types/apis/alerting.cattle.io/v1"
	"github.com/rancher/norman/types"
	"github.com/rancher/types/factory"
	"github.com/rancher/types/mapper"
)

var (
	Version = types.APIVersion{
		Version: "v1",
		Group:   "alerting.cattle.io",
		Path:    "/v1/alerting.cattle.io",
		SubContexts: map[string]bool{
			"clusters": true,
		},
	}

	Schemas = factory.Schemas(&Version).
		Init(alertingTypes)
)

func alertingTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v1.Notifier{}, &mapper.NamespaceIDMapper{}).
		MustImportAndCustomize(&Version, v1.Notifier{}, func(schema *types.Schema) {
			schema.CollectionActions = map[string]types.Action{
				"test": {Input: "notifier"},
			}
		}).
		AddMapperForType(&Version, v1.Alert{}, &mapper.NamespaceIDMapper{}).
		MustImportAndCustomize(&Version, v1.Alert{}, func(schema *types.Schema) {
			schema.MustCustomizeField("severity", func(field types.Field) types.Field {
				field.Type = "enum"
				field.Options = []string{"info", "critical", "warning"}
				field.Nullable = false
				return field
			})
			schema.MustCustomizeField("targetType", func(field types.Field) types.Field {
				field.Type = "enum"
				field.Options = []string{"node", "pod", "workload"}
				field.Nullable = false
				return field
			})
			schema.MustCustomizeField("notifierId", func(field types.Field) types.Field {
				field.Nullable = false
				return field
			})

			schema.MustCustomizeField("alertState", func(field types.Field) types.Field {
				field.Create = false
				field.Update = false
				field.Default = "active"
				field.Type = "enum"
				field.Options = []string{"active", "inactive", "alerting", "muted"}

				return field
			})
			schema.MustCustomizeField("startedAt", func(field types.Field) types.Field {
				field.Create = false
				field.Update = false
				return field
			})

			schema.ResourceActions = map[string]types.Action{
				"activate":   {},
				"deactivate": {},
				"mute":       {},
				"unmute":     {},
			}
		})
}
