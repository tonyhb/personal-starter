package v0

import (
	"context"
	"strconv"

	"gitlab.com/tonyhb/keepupdated/pkg/api/auth"
	"gitlab.com/tonyhb/keepupdated/pkg/api/v0/errors"
	"gitlab.com/tonyhb/keepupdated/pkg/types"

	"github.com/emicklei/go-restful"
)

// MustAuth is middleware which ensures that a user is authenticated before
// invoking the given handler. Being authenticated means:
//
// - The user has a valid JWT
// - We're able to load the user and the user's account
//
// Typically this should be the first middleware in a chain to authenticated
// endpoints.
//
// This adds a types.User as the "user" value of the context, and a
// types.Account as the "acct" value of the context.
func (v *V0) MustAuth(ctx context.Context, r *restful.Request, w *restful.Response) (context.Context, error) {
	idstr, err := auth.ParseJWT(r.Request)
	if err != nil {
		return ctx, errors.APIError{
			Status:  401,
			Message: "unable to authenticate",
			Detail:  err.Error(),
		}
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		return ctx, errors.APIError{
			Status:  401,
			Message: "unable to authenticate",
			Detail:  err.Error(),
		}
	}

	user, err := v.mgr.UserByID(id)
	if err != nil {
		return ctx, errors.ErrInvalidCredentials
	}

	acct, err := v.mgr.AccountByID(user.AccountID)
	if err != nil {
		return ctx, errors.ErrInvalidCredentials
	}

	// don't store pointers in contexts
	ctx = context.WithValue(ctx, "user", *user)
	ctx = context.WithValue(ctx, "acct", *acct)

	return ctx, nil
}

func (v *V0) AccountIsActive(ctx context.Context, r *restful.Request, w *restful.Response) (context.Context, error) {
	acct := ctx.Value("acct").(*types.Account)
	if !acct.IsActive() {
		return ctx, errors.ErrInactiveAccount
	}
	return ctx, nil
}
