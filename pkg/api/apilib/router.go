package apilib

import (
	"fmt"

	"github.com/emicklei/go-restful"
	"golang.org/x/net/context"
)

// Handler represents a modified restful.RouteFunction which accepts
// a context.Context as the first argument.
type Handler func(context.Context, *restful.Request, *restful.Response)

type Route struct {
	Path    string
	Method  string
	Handler interface{} // restful.RouteFunction or Handler
	Reads   interface{}
	Returns Returns
}

type Returns struct {
	Status  int
	Message string
	Data    interface{}
}

// add iterates through a list of Routes and registers them to the given
// restful.WebService
func AddRoutes(w *restful.WebService, all ...Route) {
	for _, r := range all {
		route := w.Method(r.Method).
			Path(r.Path).
			Returns(r.Returns.Status, r.Returns.Message, r.Returns.Data)

		switch r.Handler.(type) {
		case restful.RouteFunction:
			route = route.To(r.Handler.(restful.RouteFunction))
		case func(context.Context, *restful.Request, *restful.Response):
			route = route.To(wrap(r.Handler.(func(context.Context, *restful.Request, *restful.Response))))
		default:
			panic(fmt.Errorf("unknown handler type %T", r.Handler))
		}

		if r.Reads != nil {
			route = route.Reads(r.Reads)
		}

		w.Route(route)
	}
}

// wrap is used around your function if it has no middleware
func wrap(h Handler) restful.RouteFunction {
	return func(r *restful.Request, w *restful.Response) {
		ctx := context.Background()
		h(ctx, r, w)
	}
}

// Middleware
type Middleware func(context.Context, *restful.Request, *restful.Response) (context.Context, error)

type chain struct {
	middlewares []Middleware
}

func Chain(m ...Middleware) *chain {
	return &chain{
		middlewares: m,
	}
}

func (c *chain) Then(h Handler) restful.RouteFunction {
	ctx := context.Background()

	return func(r *restful.Request, w *restful.Response) {
		for i := range c.middlewares {
			var err error
			if ctx, err = c.middlewares[i](ctx, r, w); err != nil {
				// TODO: print error
				return
			}
		}

		h(ctx, r, w)
	}
}