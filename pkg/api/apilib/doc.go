package apilib

// apilib provides a wrapper around go-restful, adding more usable
// middleware which passes contexts to each Handler.
//
// The context can be modified by each middleware, allowing authors to create,
// for example, middleware which authenticate a user and adds the user to the
// given context.
//
// Each middleware can also return an error fulfilling our Writer interface
// which halts the middleware chain.
