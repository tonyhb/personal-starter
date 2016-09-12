package v0

import (
	"gitlab.com/tonyhb/keepupdated/pkg/api/apilib"
	"gitlab.com/tonyhb/keepupdated/pkg/api/v0/errors"
	"gitlab.com/tonyhb/keepupdated/pkg/manager"

	log "github.com/Sirupsen/logrus"
	"github.com/emicklei/go-restful"
)

type V0 struct {
	log *log.Logger
	mgr manager.Manager
}

type Opts struct {
	Log *log.Logger
	Mgr manager.Manager
}

func New(opts Opts) *V0 {
	return &V0{
		log: opts.Log,
		mgr: opts.Mgr,
	}
}

func (v *V0) AddRoutes(c *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/api/v0").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	apilib.AddRoutes(
		ws,
		v.LoginRoute(),
		v.RegisterRoute(),
	)

	c.Add(ws)
}

func (v *V0) WrapError(err error) errors.APIError {
	v.log.WithField("error", err).Error("error processing API call")
	return errors.APIError{
		Status:  500,
		Message: "Something bad happened",
	}
}
