package main

import (
	"net/http"
	"os"

	"gitlab.com/tonyhb/keepupdated/pkg/api"
	"gitlab.com/tonyhb/keepupdated/pkg/manager/inmemory"

	"github.com/Sirupsen/logrus"
	"github.com/opentracing/opentracing-go"
)

const (
	defaultPort     = "8888"
	defaultLogLevel = "info"
)

var (
	port, logFormat, logLevel string
	tracer                    opentracing.Tracer
)

func init() {
	tracer = opentracing.NoopTracer{}

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

	handler := api.New(api.Opts{
		Mgr:    inmemory.New(),
		Tracer: tracer,
		Logger: log,
	})
	log.WithField("port", port).Info("starting server")
	http.ListenAndServe(":"+port, handler)
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
