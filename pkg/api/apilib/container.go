package apilib

import (
	"context"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/emicklei/go-restful"
	"github.com/opentracing/opentracing-go"
)

func NewRouter(opts Opts) *Router {
	return &Router{
		Container: restful.NewContainer(),

		tracer: opts.Tracer,
		log:    opts.Logger,
	}
}

type Opts struct {
	Tracer opentracing.Tracer
	Logger *logrus.Logger
}

type Router struct {
	// Container is the restful container for all routes and restful
	// webservices
	Container *restful.Container

	tracer opentracing.Tracer
	log    *logrus.Logger
}

// AddServiceRoutes iterates through a list of Routes and registers them to
// the given restful.WebService.
//
// This should be called from each independent API package which generates
// a webservice in order to add the service to the router's container.
func (r Router) AddServiceRoutes(ws *restful.WebService, all ...Route) {
	for _, apiroute := range all {
		route := ws.Method(apiroute.Method).
			Path(apiroute.Path).
			Returns(apiroute.Returns.Status, apiroute.Returns.Message, apiroute.Returns.Data).
			Doc(apiroute.Summary).
			Notes(apiroute.Description).
			Operation(apiroute.Operation).
			To(r.getToFunc(apiroute))

		if apiroute.Reads != nil {
			route = route.Reads(apiroute.Reads)
		}

		ws.Route(route)
	}

	r.Container.Add(ws)
}

func (r *Router) getToFunc(apiroute Route) restful.RouteFunction {
	return func(req *restful.Request, w *restful.Response) {
		// Create a new opentracing span using the route's operationId as
		// the opentracing operation name, then add the span to context so
		// each middleware function and handler can access the span
		//
		// TODO: Attempt to get extract span from HTTP request in case span
		// started from client
		span := r.tracer.StartSpan(apiroute.Operation)
		defer span.Finish()
		ctx := opentracing.ContextWithSpan(context.Background(), span)

		r.log.Infof("calling %s", apiroute.Operation)

		// Invoke the handler, which may or may not have already written
		// a response. In the case of a handler with middleware the
		// middleware wrapper doesn't write a response; that's left to us
		// here.
		response := apiroute.Handler(ctx, req, w)
		if response == nil {
			return
		}

		// If the response fulfils a Writer interface use it to write our
		// response. Otherwise, just write the entity out using restful.
		if writer, ok := response.(Writer); ok {
			writer.Write(w)
			return
		}
		w.WriteHeaderAndEntity(http.StatusOK, response)
	}
}

func (r *Router) Handler() http.Handler {
	return r.Container
}
