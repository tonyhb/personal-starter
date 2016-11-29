package inmemory

import (
	"context"
	"time"

	"gitlab.com/tonyhb/keepupdated/pkg/manager"
	"gitlab.com/tonyhb/keepupdated/pkg/types"
)

type accountManager struct {
	Accounts map[int]*types.Account
}

func NewAccountManager() *accountManager {
	return &accountManager{
		Accounts: make(map[int]*types.Account),
	}
}

func (m *accountManager) CreateAccount(ctx context.Context, acct *types.Account) error {
	acct.ID = len(m.Accounts)
	acct.CreatedAt = time.Now()
	acct.UpdatedAt = acct.CreatedAt
	m.Accounts[acct.ID] = acct
	return nil
}

func (m *accountManager) UpdateAccount(ctx context.Context, acct *types.Account) error {
	if _, ok := m.Accounts[acct.ID]; !ok {
		return manager.ErrAccountNotFound
	}
	m.Accounts[acct.ID] = acct
	return nil
}

func (m *accountManager) AccountByID(ctx context.Context, id int) (*types.Account, error) {
	if _, ok := m.Accounts[id]; !ok {
		return nil, manager.ErrAccountNotFound
	}
	return m.Accounts[id], nil
}
