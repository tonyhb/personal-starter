package v0

import (
	"golang.org/x/net/context"
	"net/http"

	"github.com/emicklei/go-restful"
	"gitlab.com/tonyhb/keepupdated/pkg/api/apilib"
)

func (v *V0) LoginRoute() apilib.Route {
	return apilib.Route{
		Path:    "/login",
		Method:  "POST",
		Handler: v.Login,
		Returns: apilib.Returns{
			Status: http.StatusOK,
			// Data:   responses.JWT{},
		},
	}
}

func (v *V0) Login(ctx context.Context, req *restful.Request, w *restful.Response) {
	w.Write([]byte("lol"))
}
