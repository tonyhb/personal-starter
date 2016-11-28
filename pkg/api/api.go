package api

import (
	"net/http"

	"gitlab.com/tonyhb/keepupdated/pkg/api/apilib"
	"gitlab.com/tonyhb/keepupdated/pkg/api/v0"
	"gitlab.com/tonyhb/keepupdated/pkg/manager"

	"github.com/Sirupsen/logrus"
	"github.com/opentracing/opentracing-go"
)

func New(opts Opts) http.Handler {
	// Create a new apilib.Router, which handles the restful container,
	// middleware, and context management for the entire API suite.
	router := apilib.NewRouter(apilib.Opts{
		Tracer: opts.Tracer,
		Logger: opts.Logger,
	})

	// Add V0 routes
	v0.New(v0.Opts{
		Mgr:    opts.Mgr,
		Logger: opts.Logger,
	}).AddRoutes(router)

	return router.Handler()
}

// Opts represents configuration and initialization options for a new API
// service.
type Opts struct {
	Mgr    manager.Manager
	Tracer opentracing.Tracer
	Logger *logrus.Logger
}
