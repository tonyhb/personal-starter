package errors

import (
	"fmt"

	"github.com/emicklei/go-restful"
)

type APIError struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Detail  interface{} `json:"detail"`
}

func (a APIError) Error() string {
	return fmt.Sprintf("%s [%d]", a.Message, a.Status)
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
