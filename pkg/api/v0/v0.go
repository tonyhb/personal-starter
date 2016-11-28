package v0

import (
	"gitlab.com/tonyhb/keepupdated/pkg/api/apilib"
	"gitlab.com/tonyhb/keepupdated/pkg/api/v0/errors"
	"gitlab.com/tonyhb/keepupdated/pkg/manager"

	"github.com/Sirupsen/logrus"
	"github.com/emicklei/go-restful"
)

type V0 struct {
	mgr manager.Manager
	log *logrus.Logger
}

type Opts struct {
	Mgr    manager.Manager
	Logger *logrus.Logger
}

func New(opts Opts) *V0 {
	return &V0{
		mgr: opts.Mgr,
		log: opts.Logger,
	}
}

func (v *V0) AddRoutes(router *apilib.Router) {
	ws := new(restful.WebService)
	ws.Path("/api/v0").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	router.AddServiceRoutes(
		ws,
		v.LoginRoute(),
		v.RegisterRoute(),
	)
}

func (v *V0) WrapError(err error) errors.APIError {
	v.log.WithField("error", err).Error("error processing API call")
	return errors.APIError{
		Status:  500,
		Message: "Something bad happened",
	}
}
