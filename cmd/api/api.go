package main

import (
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	"gitlab.com/tonyhb/keepupdated/pkg/api"
	"gitlab.com/tonyhb/keepupdated/pkg/manager/inmemory"
)

const (
	defaultPort = "8888"
)

var (
	port string
)

func init() {
	if port = os.Getenv("PORT"); port == "" {
		port = defaultPort
	}
}

func main() {
	log := logrus.StandardLogger()
	api := api.New(api.Opts{
		Log: log,
		Mgr: inmemory.New(),
	})
	log.WithField("port", port).Info("starting server")
	http.ListenAndServe(":"+port, api.Handler())
}
