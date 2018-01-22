package main

import (
	"context"
	"net/http"
	"os"

	"github.com/rancher/alerting-api/server"
	"github.com/rancher/alerting-api/types/config"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", os.Getenv("KUBECONFIG"))
	if err != nil {
		return err
	}

	alert, err := config.NewAlertContext(*kubeConfig)
	if err != nil {
		return err
	}

	handler, err := server.New(context.Background(), alert)
	if err != nil {
		return err
	}

	logrus.Info("Listening on 0.0.0.0:8888")
	return http.ListenAndServe("0.0.0.0:8888", handler)
}
