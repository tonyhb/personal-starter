package apilib

import (
	"context"

	"github.com/emicklei/go-restful"
)

// Handler represents a modified restful.RouteFunction which accepts
// a context.Context as the first argument.
//
// TODO: Meybe everything returned should be a Writer.
type Handler func(context.Context, *restful.Request, *restful.Response) interface{}

// Writer is an interface which, when returned from a Handler, writes its own
// headers and data to the *resftul.Response
//
// This is useful for error types to write a custom header before writing a
// response body.
type Writer interface {
	Write(*restful.Response)
}

// Route is used to define a single route within a rest API.
// TODO: GraphQL baby.
type Route struct {
	Path        string
	Method      string
	Handler     Handler
	Returns     Returns // Required
	Reads       interface{}
	Operation   string // Swagger "operationId", and opentracing span name
	Summary     string // Swagger "summary" (restful Doc). Optional.
	Description string // Swagger "description" (restful Notes). Optional.
}

type Returns struct {
	Status  int
	Message string
	Data    interface{}
}

// Middleware is an http handler which returns new context with added values
// (eg. from authentication) and an optional error.
//
// If the error returned is non-nil the middleware chain is halted and the
// error is written to the client.
type Middleware func(context.Context, *restful.Request, *restful.Response) (context.Context, error)

type chain struct {
	middlewares []Middleware
}

// Chain creates a middleware chain which will be called in sequence breaking
// on errors.
//
// Each middleware may return a context with new values which will be passed
// to the target handler.
func Chain(m ...Middleware) *chain {
	return &chain{
		middlewares: m,
	}
}

// Then returns a function which calls all chained middlewares generating new
// context values, then finally calls the target Handler with the middleware
// context values and returns the response.
//
// Unless specifically implemented by the target Handler this does not write
// its own response. This usually happens in the router's wrapping function.
func (c *chain) Then(h Handler) Handler {
	return func(ctx context.Context, r *restful.Request, w *restful.Response) interface{} {
		for _, item := range c.middlewares {
			var err error
			if ctx, err = item(ctx, r, w); err != nil {
				return err
			}
		}

		// This will delegate to the wrapper in router.AddRoutes to write
		// to the response
		return h(ctx, r, w)
	}
}
