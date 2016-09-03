package inmemory

import (
	"gitlab.com/tonyhb/keepupdated/pkg/types"
)

type MemMgr struct {
	Accounts map[int]*types.Account
	Users    map[int]*types.User
}

func New() *MemMgr {
	return &MemMgr{
		Accounts: map[int]*types.Account{},
		Users:    map[int]*types.User{},
	}
}
