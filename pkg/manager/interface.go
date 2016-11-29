package manager

import (
	"context"

	"gitlab.com/tonyhb/keepupdated/pkg/types"
)

type Manager interface {
	AccountManager
	UserManager
}

type AccountManager interface {
	CreateAccount(context.Context, *types.Account) error
	UpdateAccount(context.Context, *types.Account) error
}

type UserManager interface {
	CreateUser(context.Context, *types.User) error
	UpdateUser(context.Context, *types.User) error
	UserByID(context.Context, int) (*types.User, error)
	UserByEmail(context.Context, string) (*types.User, error)
}
