package api

import (
	"net/http"

	"gitlab.com/tonyhb/keepupdated/pkg/api/v0"

	log "github.com/Sirupsen/logrus"
	"github.com/emicklei/go-restful"
)

func New(opts Opts) *api {
	api := &api{
		container: restful.NewContainer(),
		log:       opts.Log,
	}

	// Add V0 routes
	v0.New(v0.Opts{Log: opts.Log}).AddRoutes(api.container)

	return api
}

// Opts represents configuration and initialization options for a new API
// service.
type Opts struct {
	Log *log.Logger
}

// api is the parent container for creating a cascading KU API service
type api struct {
	container *restful.Container
	log       *log.Logger
}

func (a *api) Handler() http.Handler {
	return a.container
}
