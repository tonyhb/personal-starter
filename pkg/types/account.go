package types

import (
	"time"
)

const (
	StatusBillingError = "billingError"
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

	// Status is an enum which represents the standing of this account
	Status string

	Features
	// TODO: branding
}

func (a Account) IsActive() bool {
	return a.Status != StatusBillingError
}
