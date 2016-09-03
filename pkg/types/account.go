package types

import (
	"time"
)

// Account represents
type Account struct {
	ID             int
	StripeID       string
	CompanyName    string
	CompanyAddress string
	CompanyCity    string
	CompanyState   string
	CompanyCountry string
	CompanyZip     string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	CancelledAt    time.Time

	Features
	// TODO: branding
}
