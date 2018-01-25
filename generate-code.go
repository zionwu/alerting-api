//go:generate go run generator/cleanup/main.go
//go:generate go run generate-code.go

package main

import (
	"github.com/rancher/alerting-api/generator"
	alertSchema "github.com/rancher/alerting-api/types/apis/alerting.cattle.io/v1/schema"
)

func generate() {
	generator.Generate(alertSchema.Schemas)
}
