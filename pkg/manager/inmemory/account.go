package inmemory

import (
	"fmt"
	"time"

	"gitlab.com/tonyhb/keepupdated/pkg/types"
)

func (m *MemMgr) CreateAccount(acct *types.Account) error {
	acct.ID = len(m.Accounts)
	acct.CreatedAt = time.Now()
	acct.UpdatedAt = acct.CreatedAt
	m.Accounts[acct.ID] = acct
	return nil
}

func (m *MemMgr) UpdateAccount(acct *types.Account) error {
	if _, ok := m.Accounts[acct.ID]; !ok {
		return fmt.Errorf("account with id %d does not exist", acct.ID)
	}
	m.Accounts[acct.ID] = acct
	return nil
}
