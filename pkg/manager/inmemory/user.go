package inmemory

import (
	"context"
	"time"

	"gitlab.com/tonyhb/keepupdated/pkg/manager"
	"gitlab.com/tonyhb/keepupdated/pkg/types"
)

type userManager struct {
	Users map[int]*types.User
}

func NewUserManager() *userManager {
	return &userManager{
		Users: make(map[int]*types.User),
	}
}

func (m *userManager) CreateUser(ctx context.Context, user *types.User) error {
	user.ID = len(m.Users)
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	m.Users[user.ID] = user
	return nil
}

func (m *userManager) UpdateUser(ctx context.Context, user *types.User) error {
	if _, ok := m.Users[user.ID]; !ok {
		return manager.ErrUserNotFound
	}
	m.Users[user.ID] = user
	return nil
}

func (m *userManager) UserByID(ctx context.Context, id int) (*types.User, error) {
	u, ok := m.Users[id]
	if !ok {
		return nil, manager.ErrUserNotFound
	}
	return u, nil
}

func (m *userManager) UserByEmail(ctx context.Context, email string) (*types.User, error) {
	for _, u := range m.Users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, manager.ErrUserNotFound
}
