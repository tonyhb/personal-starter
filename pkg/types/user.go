package types

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID               int
	AccountID        int
	Name             string
	Email            string
	PasswordHash     []byte
	IsBillingContact bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (u *User) SetPassword(to string) error {
	var err error
	u.PasswordHash, err = bcrypt.GenerateFromPassword([]byte(to), bcrypt.DefaultCost)
	return err
}

func (u *User) CheckPassword(pw string) error {
	return bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(pw))
}
