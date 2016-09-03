package manager

import (
	"gitlab.com/tonyhb/keepupdated/pkg/types"
)

type Manager interface {
	CreateAccount(*types.Account) error
	UpdateAccount(*types.Account) error

	CreateUser(*types.User) error
	UpdateUser(*types.User) error
}
