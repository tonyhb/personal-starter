package inmemory

import (
	"fmt"
	"time"

	"gitlab.com/tonyhb/keepupdated/pkg/types"
)

func (m *MemMgr) CreateUser(user *types.User) error {
	user.ID = len(m.Users)
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	m.Users[user.ID] = user
	return nil
}

func (m *MemMgr) UpdateUser(user *types.User) error {
	if _, ok := m.Users[user.ID]; !ok {
		return fmt.Errorf("user with id %d does not exist", user.ID)
	}
	m.Users[user.ID] = user
	return nil
}
