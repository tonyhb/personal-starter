package v0

import (
	"gitlab.com/tonyhb/keepupdated/pkg/api/apilib"

	log "github.com/Sirupsen/logrus"
	"github.com/emicklei/go-restful"
)

type V0 struct {
	log *log.Logger
}

type Opts struct {
	Log *log.Logger
}

func New(opts Opts) *V0 {
	return &V0{
		log: opts.Log,
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
	)

	c.Add(ws)
}
