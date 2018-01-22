package schema

import (
	"github.com/rancher/alerting-api/types/apis/alerting.cattle.io/v1"
	"github.com/rancher/norman/types"
	"github.com/rancher/types/factory"
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
		AddMapperForType(&Version, v1.Notifier{}).
		MustImport(&Version, v1.Notifier{})
}
