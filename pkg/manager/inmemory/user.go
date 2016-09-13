package inmemory

import (
	"fmt"
	"time"

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

func (m *userManager) CreateUser(user *types.User) error {
	user.ID = len(m.Users)
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	m.Users[user.ID] = user
	return nil
}

func (m *userManager) UpdateUser(user *types.User) error {
	if _, ok := m.Users[user.ID]; !ok {
		return fmt.Errorf("user with id %d does not exist", user.ID)
	}
	m.Users[user.ID] = user
	return nil
}

func (m *userManager) UserByID(id int) (*types.User, error) {
	u, ok := m.Users[id]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}

func (m *userManager) UserByEmail(email string) (*types.User, error) {
	for _, u := range m.Users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}
