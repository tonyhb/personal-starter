package errors

import (
	"github.com/emicklei/go-restful"
)

type APIError struct {
	Message string
	Status  int
	Detail  interface{}
}

func (a APIError) Write(w *restful.Response) {
	w.WriteHeaderAndEntity(a.Status, a)
}

var (
	ErrInvalidCredentials = APIError{
		Message: "Invalid authentication credentials",
		Status:  401,
	}
)
