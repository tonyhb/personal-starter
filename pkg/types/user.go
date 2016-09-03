package types

type User struct {
	ID               int
	AccountID        int
	Name             string
	Email            string
	IsBillingContact bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
