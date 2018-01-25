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
		AddMapperForType(&Version, v1.Notifier{},
			&mapper.NamespaceIDMapper{},
			//		&m.Move{From: "displayName", To: "name"},
			//		&m.Move{From: "metadata/name", To: "id"}
		).
		MustImportAndCustomize(&Version, v1.Notifier{}, func(schema *types.Schema) {
			schema.CollectionActions = map[string]types.Action{
				//Add a message body as input
				"send": {},
			}
			schema.MustCustomizeField("displayName", func(field types.Field) types.Field {
				field.Nullable = false
				field.Required = true
				return field
			})
		}).
		MustImportAndCustomize(&Version, v1.SmtpConfig{}, func(schema *types.Schema) {
			schema.MustCustomizeField("host", func(field types.Field) types.Field {
				field.Type = "dnsLabel"
				field.Nullable = false
				field.Required = true
				return field
			})
			schema.MustCustomizeField("port", func(field types.Field) types.Field {
				field.Nullable = false
				field.Required = true
				//*field.Max = 65535
				//*field.Min = 1
				return field
			})
			schema.MustCustomizeField("username", func(field types.Field) types.Field {
				field.Nullable = false
				field.Required = true
				return field
			})
			schema.MustCustomizeField("password", func(field types.Field) types.Field {
				field.Nullable = false
				field.Required = true
				field.Type = "masked"
				return field
			})
			schema.MustCustomizeField("tls", func(field types.Field) types.Field {
				field.Nullable = false
				field.Required = true
				return field
			})
		}).
		MustImportAndCustomize(&Version, v1.SlackConfig{}, func(schema *types.Schema) {
			schema.MustCustomizeField("url", func(field types.Field) types.Field {
				field.Nullable = false
				field.Required = true
				return field
			})
		}).
		MustImportAndCustomize(&Version, v1.WebhookConfig{}, func(schema *types.Schema) {
			schema.MustCustomizeField("url", func(field types.Field) types.Field {
				field.Nullable = false
				field.Required = true
				return field
			})
		}).
		MustImportAndCustomize(&Version, v1.PagerdutyConfig{}, func(schema *types.Schema) {
			schema.MustCustomizeField("serviceKey", func(field types.Field) types.Field {
				field.Nullable = false
				field.Required = true
				field.Type = "masked"
				return field
			})
		}).
		AddMapperForType(&Version, v1.Alert{},
			&mapper.NamespaceIDMapper{},
			//		&m.Move{From: "displayName", To: "name"},
			//		&m.Move{From: "metadata/name", To: "id"}
		).
		MustImportAndCustomize(&Version, v1.Alert{}, func(schema *types.Schema) {
			schema.MustCustomizeField("displayName", func(field types.Field) types.Field {
				field.Nullable = false
				field.Required = true
				return field
			})

			schema.MustCustomizeField("severity", func(field types.Field) types.Field {
				field.Type = "enum"
				field.Options = []string{"info", "critical", "warning"}
				field.Nullable = false
				return field
			})

			schema.MustCustomizeField("notifierId", func(field types.Field) types.Field {
				field.Nullable = false
				field.Required = true
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
		}).
		MustImportAndCustomize(&Version, v1.TargetSystemService{}, func(schema *types.Schema) {
			schema.MustCustomizeField("type", func(field types.Field) types.Field {
				field.Type = "enum"
				field.Options = []string{"dns", "etcd", "controller manager", "scheduler", "network"}
				field.Nullable = false
				field.Required = true
				return field
			})
		}).
		MustImportAndCustomize(&Version, v1.TargetWorkload{}, func(schema *types.Schema) {
			schema.MustCustomizeField("unavailablePercentage", func(field types.Field) types.Field {
				field.Nullable = false
				field.Required = true
				//*field.Min = 1
				//*field.Max = 100
				return field
			})
		}).
		MustImportAndCustomize(&Version, v1.TargetPod{}, func(schema *types.Schema) {
			schema.MustCustomizeField("id", func(field types.Field) types.Field {
				field.Nullable = false
				field.Required = true
				return field
			})
		})

}
