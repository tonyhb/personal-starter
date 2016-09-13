package v0

import (
	"context"
	"strconv"

	"gitlab.com/tonyhb/keepupdated/pkg/api/apilib"
	"gitlab.com/tonyhb/keepupdated/pkg/api/auth"
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

func (v *V0) MustAuth(ctx context.Context, r *restful.Request, w *restful.Response) (context.Context, error) {
	idstr, err := auth.ParseJWT(r.Request)
	if err != nil {
		return ctx, errors.APIError{
			Status:  401,
			Message: err.Error(),
		}
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		return ctx, errors.APIError{
			Status:  401,
			Message: err.Error(),
		}
	}

	user, err := v.mgr.UserByID(id)
	if err != nil {
		return ctx, errors.APIError{
			Status:  401,
			Message: "Invalid credentials",
		}
	}

	return context.WithValue(ctx, "user", user), nil
}
