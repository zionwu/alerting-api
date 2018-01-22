package server

import (
	"context"
	"net/http"

	"github.com/rancher/alerting-api/types/config"
	normanapi "github.com/rancher/norman/api"
	//"github.com/rancher/norman/types"
	//"github.com/rancher/norman/store/crd"
	//"github.com/rancher/pipeline-api/types/apis/pipeline.cattle.io/v3/schema"
	//"github.com/sirupsen/logrus"

	"github.com/rancher/alerting-api/api/setup"
)

func New(ctx context.Context, alert *config.AlertingContext) (http.Handler, error) {

	setup.Schemas(ctx, alert)
	/*
		store, err := crd.NewCRDStoreFromConfig(pipeline.RESTConfig)
		if err != nil {
			return nil,err
		}

		var crdSchemas []*types.Schema

		for _, schema := range pipeline.Schemas.SchemasForVersion(schema.Version) {
			crdSchemas = append(crdSchemas, schema)
		}

		if err := store.AddSchemas(ctx, crdSchemas...); err != nil {
			logrus.Error(err)
		}
	*/

	server := normanapi.NewAPIServer()
	if err := server.AddSchemas(alert.Schemas); err != nil {
		return nil, err
	}

	//controller.Register(alert)

	return server, nil
}
