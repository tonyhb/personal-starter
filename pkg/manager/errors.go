package manager

import "fmt"

var (
	ErrAccountNotFound = fmt.Errorf("account not found")
	ErrUserNotFound    = fmt.Errorf("user not found")
)
