package v0

import (
	"context"
	"strconv"

	"gitlab.com/tonyhb/keepupdated/pkg/api/auth"
	"gitlab.com/tonyhb/keepupdated/pkg/api/v0/errors"

	"github.com/emicklei/go-restful"
)

// MustAuth is middleware which ensures that a user is authenticated before
// invoking the given handler.
//
// Typically this should be the first middleware in a chain to authenticated
// endpoints.
//
// This adds a *types.User as the "user" value of the context.
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
		return ctx, errors.ErrInvalidCredentials
	}

	return context.WithValue(ctx, "user", user), nil
}
