package main

import (
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	"gitlab.com/tonyhb/keepupdated/pkg/api"
	"gitlab.com/tonyhb/keepupdated/pkg/manager/inmemory"
)

const (
	defaultPort     = "8888"
	defaultLogLevel = "info"
)

var (
	port, logFormat, logLevel string
)

func init() {
	if port = os.Getenv("PORT"); port == "" {
		port = defaultPort
	}
	if logLevel = os.Getenv("LOG_LEVEL"); logLevel != "" {
		logLevel = defaultLogLevel
	}
}

func main() {
	log, err := getLogger()
	if err != nil {
		panic(err)
	}

	api := api.New(api.Opts{
		Log: log,
		Mgr: inmemory.New(),
	})
	log.WithField("port", port).Info("starting server")
	http.ListenAndServe(":"+port, api.Handler())
}

func getLogger() (*logrus.Logger, error) {
	lvl, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return nil, err
	}
	return &logrus.Logger{
		Out:       os.Stderr,
		Formatter: new(logrus.JSONFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     lvl,
	}, nil
}
