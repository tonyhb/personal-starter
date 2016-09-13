package manager

import (
	"gitlab.com/tonyhb/keepupdated/pkg/types"
)

type Manager interface {
	AccountManager
	UserManager
}

type AccountManager interface {
	CreateAccount(*types.Account) error
	UpdateAccount(*types.Account) error
}

type UserManager interface {
	CreateUser(*types.User) error
	UpdateUser(*types.User) error
	UserByID(int) (*types.User, error)
	UserByEmail(string) (*types.User, error)
}
